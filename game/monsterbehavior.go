package game

import "time"

type behavior struct {
	monster *monster
}

func newBehavior(m *monster) *behavior {
	return &behavior{
		monster: m,
	}
}

func (b *behavior) process(now time.Time) {

}
