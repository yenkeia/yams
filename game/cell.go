package game

import "github.com/yenkeia/yams/game/cm"

type cell struct {
	attribute cm.CellAttribute
	objects   []interface{}
}
