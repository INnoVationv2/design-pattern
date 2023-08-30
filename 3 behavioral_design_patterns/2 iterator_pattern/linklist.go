package iterator_pattern

type LinkList struct {
	head *Node
}

type Node struct {
	val  interface{}
	next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{
		val: val,
	}
}
func NewLinkList() *LinkList {
	return &LinkList{
		head: NewNode(nil),
	}
}

func (l *LinkList) Add(val interface{}) {
	node := NewNode(val)
	node.next = l.head.next
	l.head.next = node
}

func (l *LinkList) CreateIterator() Iterator {
	return &LinkListIterator{
		head: l.head,
		cur:  l.head}
}

type LinkListIterator struct {
	head *Node
	cur  *Node
}

func (i *LinkListIterator) HasNext() bool {
	if i.cur.next == nil {
		return false
	}
	return true
}

func (i *LinkListIterator) Next() (any, bool) {
	if !i.HasNext() {
		return nil, false
	}
	i.cur = i.cur.next
	return i.cur, true
}
