package game

import (
	"bufio"
	"container/list"
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
	checkList   interface{} // *function
	actList     interface{} // *function
	elseActList interface{} // *function
	say         []string
	elseSay     []string
}

var regexSharp = regexp.MustCompile(`#(\w+)`)

// TODO
func newPage(ps *pageSource) *page {
	p := new(page)
	p.name = ps.name
	checkList := &list.List{}
	actList := &list.List{}
	elseActList := &list.List{}
	say := &list.List{}
	elseSay := &list.List{}
	var cur = say
	for i := 0; i < len(ps.lines); i++ {
		line := ps.lines[i]
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
	p.say = listToStringArray(say)
	p.elseSay = listToStringArray(elseSay)
	return p
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

	ns := new(npcScript)
	ns.pages = make(map[string]*page)
	for _, s := range ps {
		p := newPage(s)
		ns.pages[p.name] = p
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
