package game

import "time"

type behaviorTree struct {
	monster *monster
	// root    node
}

func newBehaviorTree(m *monster) *behaviorTree {
	return &behaviorTree{
		monster: m,
	}
}

func (bt *behaviorTree) process(now time.Time) {

}
