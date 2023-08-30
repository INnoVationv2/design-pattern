package iterator_pattern

type List struct {
	data []int
	sz   int
}

func NewList(sz int) *List {
	return &List{
		data: make([]int, sz),
		sz:   0,
	}
}

func (l *List) Add(val int) {
	l.data[l.sz] = val
	l.sz++
}

func (l *List) CreateIterator() Iterator {
	return &ListIterator{
		list: l,
		idx:  0}
}

type ListIterator struct {
	list *List
	idx  int
}

func (i *ListIterator) HasNext() bool {
	if i.idx == i.list.sz {
		return false
	}
	return true
}

func (i *ListIterator) Next() (any, bool) {
	if !i.HasNext() {
		return nil, false
	}
	val := i.list.data[i.idx]
	i.idx++
	return val, true
}
