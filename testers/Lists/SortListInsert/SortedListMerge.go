package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for n2 != nil {
		SortListInsert(n1, n2.Data)
		n2 = n2.Next
	}
	return n1
}