package game

import "time"

type behaviorTree struct {
	monster *monster
	root    behavior
}

func newBehaviorTree(m *monster) *behaviorTree {
	return &behaviorTree{
		monster: m,
		root:    newRootNode(m),
	}
}

func (bt *behaviorTree) update(now time.Time) {

}

type behavior interface {
	tick(time.Time) status
}

func newRootNode(m *monster) behavior {
	switch m.info.AI {
	default:
		return nil
	}
}

type node struct {
	status status
}

func (n *node) tick(time.Time) status { return SUCCESS }

// 控制节点 - 选择
// 选择其子节点中的一个执行
type selectNode struct {
	node
}

// 控制节点 - 序列
// 将其所有子节点依次执行，即当前执行的一个子节点返回成功后，再执行下一个子节点
type sequenceNode struct {
	node
}

// 控制节点 - 并行
// 将其所有子节点都运行一遍，不管运行结果
type parallelNode struct {
	node
}

// 条件节点 执行返回 status
type conditionNode struct {
	node
}

// 行为节点 执行返回 status
type actionNode struct {
	node
}
