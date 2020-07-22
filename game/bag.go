package game

import (
	"fmt"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type AnyMap map[string]interface{}

type bag struct {
	typ   cm.UserItemType
	items []*userItem
}

func newBag(typ cm.UserItemType, n int) *bag {
	return &bag{typ: typ, items: make([]*userItem, n)}
}

func bagLoadFromDB(characterID int, typ cm.UserItemType, n int) *bag {
	b := newBag(typ, n)
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
