package game

type function struct {
	name   string
	method func(*npc, *player, []string) interface{}
	args   []string
}

func newFunction(name string, args []string) *function {
	return &function{
		name:   name,
		method: npcFunctionMap[name],
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
	"GOTO":         GOTO,
	"TAKEGOLD":     TAKEGOLD,
}

// CHECKPKPOINT 检查玩家善恶点数
func CHECKPKPOINT(n *npc, p *player, args []string) (res interface{}) {
	return false
}

// GOTO 跳转到下一个 page
func GOTO(n *npc, p *player, args []string) (res interface{}) {
	return
}

// TAKEGOLD 拿走玩家金币
func TAKEGOLD(n *npc, p *player, args []string) (res interface{}) {
	log.Debugf("TAKEGOLD npc: %s, player: %s, args: %s", n, p, args)
	return
}
