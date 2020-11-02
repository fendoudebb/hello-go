package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`不转义\r\n等特殊符号
原样输出`)

	s := "Hello" + " " +
		"World"
	fmt.Println("+号拼接字符串，若换行则+号必须置于末尾：", s)

	// 字符串转字节数组
	str := "Hello Golang，你好"
	fmt.Println("字符串转字节数组：", []byte(str))

	// 字节数组转字符串
	bytes := []byte{72, 101, 108, 108, 111, 32, 71, 111, 108, 97, 110, 103}
	fmt.Println("字节数组转字符串：", string(bytes))

	fmt.Println("字符串长度：", len(str))

	fmt.Println(str[0:strings.Index(str, "你好")])

	fmt.Println("是否以Hello开头：", strings.HasPrefix(str, "Hello"))
	fmt.Println("是否以Golang结尾：", strings.HasSuffix(str, "Golang"))

	fmt.Println("是否包含Go：", strings.Contains(str, "Go"))

	// 中文占3个字节
	fmt.Println("你好的索引起始位置：", strings.Index(str, "你好"))    // 15
	fmt.Println("你好的索引末尾位置：", strings.LastIndex(str, "你好"))    // 15
	fmt.Println("你字符的索引起始位置：", strings.IndexRune(str, '你')) // 15

	// n为替换的次数，n=-1时，全部替换
	fmt.Println("将`你好`全部替换为`还可以`：", strings.Replace(str, "你好", "还可以", -1))

	fmt.Println("数量统计：", strings.Count(str, "o"))

	fmt.Println("重复2次字符串：", strings.Repeat(str, 2))

	fmt.Println("转为小写：", strings.ToLower(str))
	fmt.Println("转为大写：", strings.ToUpper(str))

	fmt.Println("去除首尾空格：", strings.TrimSpace("\t  a b c    \t")) // a b c, \t也被修剪掉

	fmt.Println("去除末尾包含bc中任意一个字符（以bc结尾）：", strings.TrimRight("aaabbbcccbc", "bc")) // aaa
	fmt.Println("去除末尾包含bc中任意一个字符（以cb结尾）：", strings.TrimRight("aaabbbccccb", "bc")) // aaa

	fmt.Println("去除开头包含ac中任意一个字符（以ac开头）：", strings.TrimLeft("acaaabbbccc", "ac")) // bbbccc

	fmt.Println("去除末尾包含bc（仅去除一个）：", strings.TrimSuffix("aaabbbcccbcbc", "bc")) // aaabbbcccbcbc
	fmt.Println("去除开头包含ac（仅去除一个）：", strings.TrimPrefix("acaaabbbccc", "ac")) // aaabbbccc

	fmt.Println("去除开头结尾包含ac（去除多个）：", strings.Trim("bbbbcccaaabbbcccbc", "bc")) // aaa

	fmt.Println("以空格分隔字符串：", strings.Fields("Hello Golang! Let's study!"))

	fmt.Println("以-分割字符串：", strings.Split("2020-10-25", "-"))

	fmt.Println("以~拼接字符串数组：", strings.Join([]string{"a", "b", "c","d"},"~"))

	validUTF8Str2 := strings.ToValidUTF8("posic\xefo", "")
	fmt.Println("转换成合法的UTF8字符串（去除不合法的字符）：", validUTF8Str2)
}
