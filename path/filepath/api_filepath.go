package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(string(filepath.Separator))
	fmt.Println(string(filepath.ListSeparator ))

	wd, _ := os.Getwd()

	abs, _ := filepath.Abs(wd)

	fmt.Println(abs)

	base := filepath.Base(wd)
	fmt.Println(base)

	dir := filepath.Dir(wd)
	fmt.Println(dir)

	// 获取磁盘名
	volumeName := filepath.VolumeName(wd)
	fmt.Println(volumeName)

	// 获取文件后缀名
	// 可以是全路径
	ext := filepath.Ext("test.jpg")
	fmt.Println(ext)

	// 拼接路径
	// windows: aaa\bbb\ccc
	join := filepath.Join("aaa", "bbb", "ccc")
	fmt.Println(join)

}
