package game

import (
	"fmt"
	"strings"
)

type function struct {
	name   string
	method func(*npc, *player, []string) interface{}
	args   []string
}

func (f *function) String() string {
	return fmt.Sprintf("function: %s, args: %v", f.name, f.args)
}

func newFunction(name string, args []string) *function {
	m, ok := npcFunctionMap[name]
	if !ok {
		panic(fmt.Errorf("找不到方法: %s", name))
	}
	return &function{
		name:   name,
		method: m,
		args:   args,
	}
}

func (f *function) check(n *npc, p *player) bool {
	return f.method(n, p, f.args).(bool)
}

func (f *function) execute(n *npc, p *player) interface{} {
	return f.method(n, p, f.args)
}

var npcFunctionMap = map[string]func(*npc, *player, []string) interface{}{
	"CHECKPKPOINT": CHECKPKPOINT,
	"CHECKGOLD":    CHECKGOLD,
	"CHECKITEM":    CHECKITEM,

	"MOVE":        MOVE,
	"GOTO":        GOTO,
	"TAKEGOLD":    TAKEGOLD,
	"GIVESKILL":   GIVESKILL,
	"GIVEGOLD":    GIVEGOLD,
	"LEVEL":       LEVEL,
	"CHANGELEVEL": CHANGELEVEL,
	"GIVEITEM":    GIVEITEM,
	"CLEARPETS":   CLEARPETS,
	"LINEMESSAGE": LINEMESSAGE,
	"SET":         SET,
}

// CHECKPKPOINT 检查玩家善恶点数
func CHECKPKPOINT(n *npc, p *player, args []string) (res interface{}) {
	return false
}

// CHECKGOLD 检查玩家金币
func CHECKGOLD(n *npc, p *player, args []string) (res interface{}) {
	return false
}

// CHECKITEM 检查是否有物品
func CHECKITEM(n *npc, p *player, args []string) (res interface{}) {
	return false
}

// MOVE 传送到新地图
func MOVE(n *npc, p *player, args []string) (res interface{}) {
	return
}

// GOTO 跳转到下一个 page
func GOTO(n *npc, p *player, args []string) (res interface{}) {
	key := "[" + strings.ToUpper(args[0]) + "]"
	return cmdGoto{gotoPage: key}
}

// TAKEGOLD 拿走玩家金币
func TAKEGOLD(n *npc, p *player, args []string) (res interface{}) {
	log.Debugf("TAKEGOLD npc: %v, player: %v, args: %v", n, p, args)
	return
}

// GIVESKILL 给玩家添加技能
func GIVESKILL(n *npc, p *player, args []string) (res interface{}) {
	return
}

// GIVEGOLD 给玩家增加金币
func GIVEGOLD(n *npc, p *player, args []string) (res interface{}) {
	return
}

// LEVEL 升级
func LEVEL(n *npc, p *player, args []string) (res interface{}) {
	return
}

// CHANGELEVEL ...
func CHANGELEVEL(n *npc, p *player, args []string) (res interface{}) {
	return
}

// GIVEITEM 给物品
func GIVEITEM(n *npc, p *player, args []string) (res interface{}) {
	return
}

// CLEARPETS ..
func CLEARPETS(n *npc, p *player, args []string) (res interface{}) {
	return
}

// LINEMESSAGE ..
func LINEMESSAGE(n *npc, p *player, args []string) (res interface{}) {
	return
}

// SET ..
func SET(n *npc, p *player, args []string) (res interface{}) {
	return
}
