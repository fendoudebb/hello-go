package main

import (
	"fmt"
	"os"
)

func main() {
	environ := os.Environ()
	fmt.Println(environ)

	// 获取hostname
	hostname, _ := os.Hostname()
	fmt.Println(hostname)

	// 获取环境变量
	envGOPATH := os.Getenv("GOPATH")
	fmt.Println(envGOPATH)

	// 设置环境变量
	_ = os.Setenv("test-key", "test-value")
	fmt.Println("获取", os.Getenv("test-key"))

	// 清空环境变量
	//os.Clearenv()

	// 删除环境变量
	_ = os.Unsetenv("test-key")
	fmt.Println(os.Getenv("test-key"))

	// 查找环境变量，成功返回字符串和true，失败返回空串和false
	envGOOS, b := os.LookupEnv("GOOS")
	if b {
		fmt.Println(envGOOS)
	}

	// 新建文件
	if file, err := os.Create("test-file"); err == nil {
		fmt.Println(file.Name())
		file.Close()
	}

	if fileInfo, err := os.Stat("test-file"); err == nil {
		fmt.Println(fileInfo)
	}

	// 删除文件
	err := os.Remove("test-file")
	fmt.Println(err)

	dir, _ := os.Getwd()
	fmt.Println(dir)


	if fileInfo, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println(fileInfo)
	}

	// 路径分割符
	pathSeparator := string(os.PathSeparator)
	fmt.Println(pathSeparator)

	// 分号
	pathListSeparator := string(os.PathListSeparator)
	fmt.Println(pathListSeparator)

	tempDir := os.TempDir()
	fmt.Println(tempDir)

	cacheDir, _ := os.UserCacheDir()
	fmt.Println(cacheDir)

	configDir, _ := os.UserConfigDir()
	fmt.Println(configDir)

	homeDir, _ := os.UserHomeDir()
	fmt.Println(homeDir)

	lstat, err := os.Lstat("README.md")
	fmt.Println(lstat)

	err = os.Rename("abc", "ddd")
	fmt.Println(err)

}
