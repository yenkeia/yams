package game

import (
	"container/list"
	"time"

	"github.com/yenkeia/yams/game/cm"
)

type buff struct {
	buffType   cm.BuffType
	casterID   int
	visible    bool      // 是否可见
	expireTime time.Time // 过期时间️
	values     []int     // public int[] Values
	infinite   bool      // 是否永久
	paused     bool
}

func newBuff(buffType cm.BuffType, casterID int, expireTime time.Time, values []int) *buff {
	return &buff{
		buffType:   buffType,
		casterID:   casterID,
		visible:    false,
		expireTime: expireTime,
		values:     values,
		infinite:   false,
		paused:     false,
	}
}

type buffList struct {
	ls *list.List
}

func newBuffList() *buffList {
	bl := new(buffList)
	bl.ls = new(list.List)
	return bl
}

// 判断是否存在这种 buff
func (bl *buffList) has(buffType cm.BuffType) bool {
	for it := bl.ls.Front(); it != nil; it = it.Next() {
		b := it.Value.(*buff)
		if b.buffType == buffType {
			return true
		}
	}
	return false
}

func (bl *buffList) addBuff(b *buff) {
	for it := bl.ls.Front(); it != nil; it = it.Next() {
		buf := it.Value.(*buff)
		if buf.buffType != b.buffType {
			continue
		}
		// 新的替换旧的
		b.paused = false
		bl.ls.InsertBefore(b, it)
		bl.ls.Remove(it)
		return
	}
	bl.ls.PushBack(b)
}

func (bl *buffList) removeBuff(buffType cm.BuffType) {
	for it := bl.ls.Front(); it != nil; it = it.Next() {
		buf := it.Value.(*buff)
		if buf.buffType != buffType {
			continue
		}
		buf.infinite = false
		buf.expireTime = time.Now()
	}
}
