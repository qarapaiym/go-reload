package student

type NodeI struct {
	Data int
	Next *NodeI
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

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

func ListSort(l *NodeI) *NodeI {
	i := l
	for i != nil {
		if i.Next != nil && i.Data > i.Next.Data {
			i.Data, i.Next.Data = i.Next.Data, i.Data
			i = l
		} else {
			i = i.Next
		}
	}
	return l
}

func ListMerge(l1 *NodeI, l2 *NodeI) *NodeI {
	for i := l2; i != nil; i = i.Next {
		l1 = listPushBack(l1, i.Data)
	}
	return l1
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListMerge(n1, n2)
	n1 = ListSort(n1)
	return n1
}
