//go:build !windows
// +build !windows

package cfg

import (
	"os"
	"path"
)

var (
	FilePrefix = path.Join(os.Getenv("HOME"), ".etcdm")
)

func init() {
	// 检查 FilePrefix 文件夹是否存在
	if _, err := os.Stat(FilePrefix); os.IsNotExist(err) {
		// 如果不存在则创建文件夹
		if err := os.MkdirAll(FilePrefix, 0755); err != nil {
			panic(err)
		}
	}
}
