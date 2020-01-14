package student

import (
	"fmt"

	student ".."
)

func PrintList(l *student.NodeI) {
	it := lpackage student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for n2 != nil {
		SortListInsert(n1, n2.Data)
		n2 = n2.Next
	}
	return n1
}
	for it != nil {package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for n2 != nil {
		SortListInsert(n1, n2.Data)
		n2 = n2.Next
	}
	return n1
}
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *student.NodeI
	var link2 *student.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(student.SortedListMerge(link2, link))
}