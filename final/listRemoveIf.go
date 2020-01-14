package student


type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		l.Tail = l.Head
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	first := l.Head
	var second *NodeL
	if l.Head.Next != nil {
		second = l.Head.Next
	} else {
		if l.Head.Data == data_ref {
			l.Head = l.Head.Next
		}
		return
	}
	if first.Data == data_ref {
		l.Head = l.Head.Next
		ListRemoveIf(l, data_ref)
	}
	for second != nil {
		if second.Data == data_ref {
			first.Next = second.Next
		}
		first = second
		second = second.Next
	}
}
