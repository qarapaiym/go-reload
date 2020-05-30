package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	newList := &List{}
	for i := l.Head; i != nil; i = i.Next {
		if i.Data != data_ref {
			ListPushBack(newList, i.Data)
		}
	}
	*l = *newList
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n

}
