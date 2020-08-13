package game

import (
	"container/list"
	"time"

	"github.com/yenkeia/yams/game/cm"
)

type poison struct {
	ownerID      int
	poisonType   cm.PoisonType
	value        int           // 效果总数
	tickDuration time.Duration // 两次生效间隔
	tickNextTime time.Time     // 下次生效时间
	tickTime     int           // 当前第几跳
	tickNum      int           // 总共跳几次
}

// newPoison
// tickDuration 两次生效间隔
// tickNum 总共跳几次
// value 总数值
func newPoison(tickDuration time.Duration, tickNum int, ownerID int, poisonType cm.PoisonType, value int) *poison {
	return &poison{
		ownerID:      ownerID,
		poisonType:   poisonType,
		value:        value,
		tickDuration: tickDuration,
		tickNextTime: time.Now().Add(tickDuration),
		tickTime:     0,
		tickNum:      tickNum,
	}
}

type poisonList struct {
	l *list.List
}

func newPoisonList() *poisonList {
	ls := new(poisonList)
	ls.l = new(list.List)
	return ls
}

func (pl *poisonList) add(ps *poison) {
	pl.l.PushBack(ps)
}
