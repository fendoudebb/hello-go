package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	Write()
	ReadBigFile()
	ReadSmallFile()

}

func Write() {
	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// UTF BOM 防止中文乱码
	_, _ = f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	_ = w.Write([]string{"第一列", "第二列", "第三列"})
	_ = w.Write([]string{"111", "222", "333"})
	_ = w.Write([]string{"111", "\"", "\""})
	_ = w.Write([]string{"111", `"""`, `"""`})
	_ = w.Write([]string{"111", `""`, `""`})
	_ = w.Write([]string{"111", `''`, `''`})
	_ = w.Write([]string{"111", `'`, `'`})
	_ = w.Write([]string{"111", `'''"`, `'''"`})

	_ = w.WriteAll([][]string{
		{"111", "\r\n", "333"},
		{"111", "\n", "333"},
		{"111", "\r", "333"},
	})

	w.Flush()
}

func ReadBigFile() {
	//准备读取文件
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
	}

}

func ReadSmallFile() {
	f, _ := os.Open("test.csv")
	defer f.Close()

	r1 := csv.NewReader(f)

	content, err := r1.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range content {
		fmt.Println(row)
	}
}
