package game

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// npcScript 每个文件是一个 script
type npcScript struct {
	types  []int    // 出售的商品类型
	quests []int    // 任务相关
	trade  []string // 出售的商品
	pages  map[string]*page
}

// page 是每个 [...] 及以下部分
type page struct {
	name        string
	checkList   []*function
	actList     []*function
	elseActList []*function
	say         []string
	elseSay     []string
}

func (p *page) String() string {
	return fmt.Sprintf("page: %s, say: %s, elseSay: %s\n", p.name, p.say, p.elseSay)
}

var regexSharp = regexp.MustCompile(`#(\w+)`)

func newPage(ps *pageSource) *page {
	p := new(page)
	p.name = ps.name
	p.say = make([]string, 0)
	p.elseSay = make([]string, 0)
	p.checkList = make([]*function, 0)
	p.actList = make([]*function, 0)
	p.elseActList = make([]*function, 0)
	checkList := &list.List{}
	actList := &list.List{}
	elseActList := &list.List{}
	say := &list.List{}
	elseSay := &list.List{}
	var cur = say
	for i := 0; i < len(ps.lines); i++ {
		line := ps.lines[i]
		if line == "" {
			continue
		}
		if line[0] == '#' {
			match := regexSharp.FindStringSubmatch(line)
			switch strings.ToUpper(match[1]) {
			case "IF":
				cur = checkList
			case "SAY":
				cur = say
			case "ACT":
				cur = actList
			case "ELSEACT":
				cur = elseActList
			case "ELSESAY":
				cur = elseSay
			default:
				panic("error:" + p.name + "---" + match[1])
			}
			continue
		}
		cur.PushBack(trimEnd(line))
	}
	p.parseCheck(checkList)
	p.parseAct(actList)
	p.parseElseAct(elseActList)
	p.say = listToStringArray(say)
	p.elseSay = listToStringArray(elseSay)
	return p
}

func (p *page) parseCheck(checkList *list.List) {
	for it := checkList.Front(); it != nil; it = it.Next() {
		v := it.Value.(string)
		res := strings.Split(v, " ")
		p.checkList = append(p.checkList, newFunction(res[0], res[1:]))
	}
}

func (p *page) parseAct(actList *list.List) {
	// fmt.Println(">>>parseAct start")
	for it := actList.Front(); it != nil; it = it.Next() {
		v := it.Value.(string)
		res := strings.Split(v, " ")
		p.actList = append(p.actList, newFunction(res[0], res[1:]))
	}
	// fmt.Println("<<<parseAct end")
}

func (p *page) parseElseAct(elseActList *list.List) {
	// fmt.Println("===parseElseAct start")
	for it := elseActList.Front(); it != nil; it = it.Next() {
		v := it.Value.(string)
		res := strings.Split(v, " ")
		p.elseActList = append(p.elseActList, newFunction(res[0], res[1:]))
	}
	// fmt.Println("===parseElseAct end")
}

func trimEnd(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func listToStringArray(l *list.List) []string {
	ret := []string{}
	for it := l.Front(); it != nil; it = it.Next() {
		ret = append(ret, it.Value.(string))
	}
	return ret
}

type pageSource struct {
	name  string
	lines []string
}

func newNPCScript(path string) *npcScript {
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(reader)
	var cur *pageSource
	ps := make([]*pageSource, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") {
			if cur != nil {
				ps = append(ps, cur)
			}
			cur = new(pageSource)
			cur.name = line
			cur.lines = make([]string, 0)
		}
		cur.lines = append(cur.lines, line)
	}
	ps = append(ps, cur)

	// fmt.Println("============")
	ns := new(npcScript)
	ns.pages = make(map[string]*page)
	for _, s := range ps {
		p := newPage(s)
		ns.pages[strings.ToUpper(p.name)] = p
		if s.name == "[Types]" {
			ns.parseTypes(s.lines)
		}
		if s.name == "[Quests]" {
			ns.parseQuests(s.lines)
		}
		if s.name == "[Trade]" {
			ns.parseTrade(s.lines)
		}
	}
	return ns
}

func (ns *npcScript) parseTypes(lines []string) {
	// fmt.Println("types ->", lines)
	for index, line := range lines {
		if index == 0 || line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ns.types = append(ns.types, i)
	}
}

func (ns *npcScript) parseQuests(lines []string) {
	// fmt.Println("quests ->", lines)
	for index, line := range lines {
		if index == 0 || line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ns.quests = append(ns.quests, i)
	}
}

func (ns *npcScript) parseTrade(lines []string) {
	// fmt.Println("trade ->", lines)
	for index, line := range lines {
		if index == 0 || line == "" {
			continue
		}
		ns.trade = append(ns.trade, line)
	}
}

type cmdBreak struct{}

type cmdGoto struct {
	gotoPage string
}

func (ns *npcScript) call(pageKey string, n *npc, p *player) (say []string, err error) {
	pg, ok := ns.pages[pageKey]
	if !ok {
		return nil, fmt.Errorf("page key not found: %s", pageKey)
	}
	ok = true
	for _, c := range pg.checkList {
		if !c.check(n, p) {
			ok = false
			break
		}
	}
	var acts []*function
	if ok {
		acts = pg.actList
		say = pg.say
	} else {
		acts = pg.elseActList
		say = pg.elseSay
	}
ACTION:
	for _, act := range acts {
		res := act.execute(n, p)
		switch cmd := res.(type) {
		case cmdBreak: // 多用于任务，例如拿走任务物品后，跳出 acts，返回 say 给客户端。 例子 Envir/NPCs/MongchonProvince/14Qas-D604.txt
			break ACTION
		case cmdGoto:
			return ns.call(cmd.gotoPage, n, p)
		case nil:
			continue
		}
	}
	return
}
