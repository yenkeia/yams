package game

type mirMap struct {
	width   int
	height  int
	version int
	cells   []*cell
}

func newMirMap(width, height, version int) *mirMap {
	return nil
}

func (m *mirMap) setCellAttribute(x, y int, attr cellAttribute) {

}
