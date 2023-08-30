package strategy_pattern

import (
	"fmt"
)

type Strategy interface {
	Sort(data []any)
}

type QuickSort struct{}

func (q *QuickSort) Sort(data []any) {
	fmt.Printf("QuickSort %v\n", data)
}

type MergeSort struct{}

func (q *MergeSort) Sort(data []any) {
	fmt.Printf("MergeSort %v\n", data)
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (c *Context) performSort(data []any) {
	c.strategy.Sort(data)
}

func Main() {
	data := []any{1, 2, 3, 4}
	quickSortStrategy := QuickSort{}
	mergeSortStrategy := MergeSort{}

	context := NewContext(&quickSortStrategy)
	context.performSort(data)

	context.strategy = &mergeSortStrategy
	context.performSort(data)
}
