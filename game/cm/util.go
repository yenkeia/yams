package cm

import (
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"unicode"
)

// GetFiles 返回目录下所有文件路径
func GetFiles(dir string, allow []string) []string {
	allowMap := map[string]bool{}
	if allow != nil {
		for _, v := range allow {
			allowMap[v] = true
		}
	}
	ret := []string{}
	filepath.Walk(dir, func(fpath string, f os.FileInfo, err error) error {
		if f == nil || f.IsDir() {
			return nil
		}

		ext := path.Ext(fpath)
		if allowMap[ext] {
			ret = append(ret, filepath.ToSlash(fpath))
		}

		return nil
	})
	return ret
}

// RandomInt 随机 [low, high]
func RandomInt(low int, high int) int {
	if low == high {
		return low
	}
	if low > high || high == 0 {
		return 0
	}
	return rand.Intn(high-low+1) + low
}

// RandomDirection ...
func RandomDirection() MirDirection {
	return MirDirection(RandomInt(0, MirDirectionCount))
}

// DirectionFromPoint 返回原点 source 到 dest 目标点的方向
func DirectionFromPoint(source, dest Point) MirDirection {
	if source.X < dest.X {
		if source.Y < dest.Y {
			return MirDirectionDownRight
		}
		if source.Y > dest.Y {
			return MirDirectionUpRight
		}
		return MirDirectionRight
	}
	if source.X > dest.X {
		if source.Y < dest.Y {
			return MirDirectionDownLeft
		}
		if source.Y > dest.Y {

			return MirDirectionUpLeft
		}
		return MirDirectionLeft
	}
	if source.Y < dest.Y {
		return MirDirectionDown
	}
	return MirDirectionUp
}

// AbsInt 绝对值
func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// MaxInt ..
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// InRange 判断 a, b 两个坐标是否在距离 i 范围内
func InRange(a, b Point, i int) bool {
	return AbsInt(int(a.X)-int(b.X)) <= i && AbsInt(int(a.Y)-int(b.Y)) <= i
}

// MaxDistance ..
func MaxDistance(p1, p2 Point) int {
	return MaxInt(AbsInt(int(p1.X)-int(p2.X)), AbsInt(int(p1.Y)-int(p2.Y)))
}

// PointMove ..
func PointMove(p Point, dir MirDirection, step int) Point {
	return p.NextPoint(dir, uint32(step))
}

// SplitString 按空格拆分字符串。如果加了引号，那么认为是一个字符串
func SplitString(s string) []string {
	ret := []string{}
	start := 0
	var stat rune
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if unicode.IsSpace(r[i]) {
			if stat == 1 {
				ret = append(ret, string(r[start:i]))
				stat = 0
			}
		} else if r[i] == '\'' || r[i] == '"' {
			if stat == r[i] {
				ret = append(ret, string(r[start:i]))
				stat = 0
			} else {
				if stat == 0 {
					stat = r[i]
					start = i + 1
				}
			}
		} else {
			if stat == 0 {
				start = i
				stat = 1
			}
		}
	}
	if stat != 0 {
		ret = append(ret, string(r[start:]))
	}
	return ret
}
