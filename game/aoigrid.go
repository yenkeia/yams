package game

import "fmt"

type aoiGrid struct {
	gID       int          //格子ID
	minX      int          //格子左边界坐标
	maxX      int          //格子右边界坐标
	minY      int          //格子上边界坐标
	maxY      int          //格子下边界坐标
	objectIDs map[int]bool //当前格子内的玩家或者物体成员ID
}

//初始化一个格子
func newAOIGrid(gID, minX, maxX, minY, maxY int) *aoiGrid {
	return &aoiGrid{
		gID:       gID,
		minX:      minX,
		maxX:      maxX,
		minY:      minY,
		maxY:      maxY,
		objectIDs: make(map[int]bool),
	}
}

//向当前格子中添加一个玩家
func (g *aoiGrid) add(objID int) {
	g.objectIDs[objID] = true
}

//从格子中删除一个玩家
func (g *aoiGrid) remove(objID int) {
	delete(g.objectIDs, objID)
}

//得到当前格子中所有的玩家
func (g *aoiGrid) getObjectIDs() (objectIDs []int) {
	// for k, _ := range g.objectIDs {
	for k := range g.objectIDs {
		objectIDs = append(objectIDs, k)
	}
	return
}

//打印信息方法
func (g *aoiGrid) String() string {
	return fmt.Sprintf("Grid id: %d, minX:%d, maxX:%d, minY:%d, maxY:%d, objectIDs:%v",
		g.gID, g.minX, g.maxX, g.minY, g.maxY, g.objectIDs)
}
