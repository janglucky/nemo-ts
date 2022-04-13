package file

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/fs"
	"path/filepath"
)

// ListFiles ...列出fold下的所有文件（不递归遍历）
func ListFiles(fold string)  []string{
	var files []string
	filepath.Walk(fold, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println(path)
			files = append(files, path)
		}
		return nil
	})
	return files
}

// DecodeTomlFile ...读取toml配置文件
func DecodeTomlFile(filename string, content interface{})  error{
	if _, err := toml.DecodeFile(filename, content); err != nil {
		return err
	}
	return nil
}