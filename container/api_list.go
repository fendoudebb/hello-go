package main

import (
	"container/list"
)

func main() {
	l := list.New()
	l.PushFront("aaa")
	elementC := l.PushBack("ccc")
	l.InsertBefore("bbb", elementC)
	l.InsertAfter("ddd", elementC)

}
