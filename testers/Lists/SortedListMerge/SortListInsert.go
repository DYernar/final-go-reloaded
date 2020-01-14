package student

type NodeI struct {
	Data int
	Next *NodeI
}


func SortListInsert(l *NodeI, data_ref int) *NodeI{
	addr := l
	for l != nil {
		if l.Data > data_ref{
			var newNode NodeI
			newNode.Data = data_ref
			newNode.Next = l
			return &newNode
		}
		if l.Next.Data > data_ref {
			var newNode NodeI
			newNode.Data = data_ref
			newNode.Next = l.Next
			l.Next = &newNode
			return addr
		}
		l = l.Next
	}
	return addr
}