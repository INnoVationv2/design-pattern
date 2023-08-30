package iterator_pattern

import (
	"fmt"
)

type Iterator interface {
	HasNext() bool
	Next() (any, bool)
}

type Collection interface {
	CreateIterator() Iterator
}

func visit(iterator Iterator) {
	for iterator.HasNext() {
		val, _ := iterator.Next()
		fmt.Printf("  %v  ", val)
	}
	fmt.Println()
}

func Main() {
	list := NewList(10)
	linkList := NewLinkList()
	for i := 0; i < 10; i++ {
		list.Add(i)
		linkList.Add(i)
	}
	fmt.Println("List:")
	visit(list.CreateIterator())
	fmt.Println("LinkList:")
	visit(linkList.CreateIterator())
}
