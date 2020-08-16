package game

import (
	"fmt"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

// AnyMap ...
type AnyMap map[string]interface{}

type bag struct {
	characterID int
	typ         cm.UserItemType
	items       []*userItem
}

func newBag(cid int, typ cm.UserItemType, n int) *bag {
	return &bag{characterID: cid, typ: typ, items: make([]*userItem, n)}
}

func bagLoadFromDB(characterID int, typ cm.UserItemType, n int) *bag {
	b := newBag(characterID, typ, n)
	cui := []*orm.CharacterUserItem{}
	pdb.db.Table("character_user_item").Where("character_id = ? AND type = ?", characterID, typ).Find(&cui)
	ids := make([]int, n)
	userItemIDIndexMap := make(map[int]int)
	for i, item := range cui {
		ids[i] = item.UserItemID
		userItemIDIndexMap[item.UserItemID] = item.Index
	}
	items := []*orm.UserItem{}
	pdb.db.Table("user_item").Where("id in (?)", ids).Find(&items)
	for _, ui := range items {
		b.items[userItemIDIndexMap[ui.ID]] = newUserItemFromORM(ui)
	}
	return b
}

func (b *bag) serverUserItems() []*server.UserItem {
	res := make([]*server.UserItem, len(b.items))
	for i := range b.items {
		item := b.items[i]
		if item == nil {
			res[i] = nil
			continue
		}
		res[i] = item.serverUserItem()
	}
	return res
}

func (b *bag) move(from, to int) (err error) {
	return b.moveTo(from, to, b)
}

func (b *bag) moveTo(from, to int, tobag *bag) error {
	if from < 0 || to < 0 || from > len(b.items) || to > len(tobag.items) {
		return fmt.Errorf("Move: 位置不存在 from=%d to=%d", from, to)
	}
	item := b.items[from]
	if item == nil {
		return fmt.Errorf("格子 %d 没有物品", from)
	}
	pdb.db.Table("character_user_item").Where("user_item_id = ?", item.id).Update(AnyMap{
		"type":  tobag.typ,
		"index": to,
	})
	toItem := tobag.items[to]
	if toItem != nil {
		pdb.db.Table("character_user_item").Where("user_item_id = ?", toItem.id).Update(AnyMap{
			"type":  b.typ,
			"index": from,
		})
	}
	b.items[from], tobag.items[to] = tobag.items[to], b.items[from]
	return nil
}

func (b *bag) get(i int) *userItem {
	return b.items[i]
}

func (b *bag) set(i int, ui *userItem) (err error) {
	if ui != nil {
		if b.items[i] != nil {
			return fmt.Errorf("该位置有物品了")
		}
		oui := ui.ormUserItem()
		pdb.db.Table("user_item").Create(oui)
		ui.id = oui.ID
		pdb.db.Table("character_user_item").Create(&orm.CharacterUserItem{
			CharacterID: int(b.characterID),
			UserItemID:  int(ui.id),
			Type:        int(b.typ),
			Index:       i,
		})
		b.items[i] = ui
	} else {
		ui = b.items[i]
		if ui == nil {
			return fmt.Errorf("尝试删除空位置的物品")
		}
		pdb.db.Table("user_item").Where("id = ?", ui.id).Delete(&orm.UserItem{})
		pdb.db.Table("character_user_item").Where("user_item_id = ?", ui.id).Delete(&orm.CharacterUserItem{})
		b.items[i] = nil
	}
	return
}

func (b *bag) setCount(index int, count int) {
	if count == 0 {
		log.Infof("Delete UserItem %d \n", b.items[index].id)
		b.set(index, nil)
	} else {
		pdb.db.Table("user_item").Where("id = ?", b.items[index].id).Update("count", count)
		b.items[index].count = count
	}
	return
}

func (b *bag) useCount(i int, c int) {
	b.setCount(i, b.items[i].count-c)
}
