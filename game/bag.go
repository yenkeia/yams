package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

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
