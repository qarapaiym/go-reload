package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	i := l
	j := 0
	for i != nil {
		if j == 0 && n.Data < i.Data {
			n.Next = l
			l = n
			return l
		} else if i.Next == nil {
			i.Next = n
			return l
		} else if i.Data <= data_ref && i.Next.Data > data_ref {
			n.Next = i.Next
			i.Next = n
			return l
		} else {
			i = i.Next
		}
		j++
	}
	return l
}
