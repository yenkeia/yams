package game

const (
	// LOGIN 客户端连接到服务器，正在输入账号密码的状态
	LOGIN = iota
	// SELECT 选角色状态
	SELECT
	// GAME 进入游戏状态
	GAME
	// DISCONNECTED 应该是小退后的状态
	DISCONNECTED
)

type cellAttribute int

const (
	cellAttributeWalk     cellAttribute = 0
	cellAttributeHighWall cellAttribute = 1
	cellAttributeLowWall  cellAttribute = 2
)
