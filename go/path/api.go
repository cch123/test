package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	str := "/user/local/go/a.txt"
	println("base", path.Base(str))
	println("clean", path.Clean(str))
	println("dir", path.Dir(str))
	println(path.Ext(str))
	println(path.IsAbs(str))
	println("join", path.Join("/abc/", str))
	m, err := path.Match(str, "/user/local")
	println(m, err)
	fmt.Println(path.Split(str))

	// 针对目录下的所有文件做遍历，不用自己写递归了
	filepath.Walk("/usr/local/go/src", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.IsDir(), info.Name())
		return nil
	})
}

//type WalkFunc func(path string, info os.FileInfo, err error) error
