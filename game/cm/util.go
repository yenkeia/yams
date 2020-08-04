package cm

import (
	"math/rand"
	"os"
	"path"
	"path/filepath"
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
	} else {
		return MirDirectionUp
	}
}

// AbsInt 绝对值
func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// InRange 判断 a, b 两个坐标是否在距离 i 范围内
func InRange(a, b Point, i int) bool {
	return AbsInt(int(a.X)-int(b.X)) <= i && AbsInt(int(a.Y)-int(b.Y)) <= i
}
