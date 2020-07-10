package cm

import (
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
