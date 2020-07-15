package game

import (
	"github.com/yenkeia/yams/game/orm"
)

// item 是数据库里加载出来的 info 的包装
type item struct {
	info *orm.ItemInfo
}

func newItem(ii *orm.ItemInfo) *item {
	return &item{info: ii}
}

type userItem struct {
	*orm.UserItem
}
