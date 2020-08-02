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
	status   status
	children []behavior
}

func (n *node) tick(time.Time) status { return SUCCESS }

// 控制节点 - 选择
// 顺序执行所有的子节点，当一个子节点执行结果为 SUCCESS/RUNNING 的时候终止执行并返回 SUCCESS/RUNNING
// 选择节点可以被理解为一个或门（OR gate）
type selectNode struct {
	node
}

func (n *selectNode) tick(time.Time) status {
	return SUCCESS
}

// 控制节点 - 序列
// 将其所有子节点依次执行，即当前执行的一个子节点返回成功后，再执行下一个子节点
// 顺序依次执行子节点，如果所有子节点都返回 SUCCESS/RUNNING，则向其父节点返回 SUCCESS/RUNNING
// 序列节点可以理解为与门（AND gate）
type sequenceNode struct {
	node
}

func (n *sequenceNode) tick(time.Time) status {
	return SUCCESS
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
