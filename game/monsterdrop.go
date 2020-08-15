package game

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type dropInfo struct {
	low      int    // 1/3 中的分子
	high     int    // 1/3 中的分母
	count    int    // 掉落数量
	itemName string // orm.ItemInfo.Name
	isQuest  bool   // 是否是任务物品
}

func loadDropFile(path string, itemInfoNameMap map[string]*orm.ItemInfo) (res []*dropInfo, err error) {
	reader, err := os.Open(path)
	if err != nil {
		return
	}
	lines := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	chanceReg := regexp.MustCompile(`(\d+)/(\d+)`)
	lineError := func(line int, detail string) error {
		return fmt.Errorf("DropInfo 格式不正确，%s行%d:%s %s", path, line, lines[line], detail)
	}
	res = make([]*dropInfo, 0)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}
		str := cm.SplitString(line)
		if len(str) != 3 && len(str) != 2 {
			return nil, lineError(i, "参数个数")
		}
		match := chanceReg.FindStringSubmatch(str[0])
		low, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, lineError(i, "分子错误")
		}
		high, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, lineError(i, "分母错误")
		}
		itemName := str[1]
		// _, ok := itemInfoNameMap[itemName]
		// if !ok {
		// 	log.Warnln("物品不存在: " + itemName)
		// }
		info := &dropInfo{low: low, high: high, itemName: itemName, count: 1}
		if len(str) == 3 { // 1/10 Gold 500
			if strings.ToUpper(str[2]) == "Q" {
				info.isQuest = true
			} else {
				count, err := strconv.Atoi(str[2])
				info.count = count
				if err != nil {
					for i, v := range str {
						fmt.Println(i, v)
					}
					return nil, lineError(i, "参数错误")
				}
			}
		}
		res = append(res, info)
	}
	return
}
