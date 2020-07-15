package game

import "github.com/yenkeia/yams/game/orm"

type item struct {
	info *orm.ItemInfo
}

func newItem(ii *orm.ItemInfo) *item {
	return &item{info: ii}
}
