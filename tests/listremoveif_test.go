package tests

import (
	"fmt"
	"testing"

	student "../ListRemoveIf"
)

func TestListRemoveIf(t *testing.T) {
	tables := []struct {
		list    []interface{}
		dataRef interface{}
		answer  string
	}{
		{
			[]interface{}{},
			1,
			"<nil>",
		},
		{
			[]interface{}{},
			96,
			"<nil>",
		},
		{
			[]interface{}{98, 98, 33, 34, 33, 34, 33, 89, 33},
			34,
			"98 -> 98 -> 33 -> 33 -> 33 -> 89 -> 33 -> <nil>",
		},
		{
			[]interface{}{79, 74, 99, 79, 7},
			99,
			"79 -> 74 -> 79 -> 7 -> <nil>",
		},
		{
			[]interface{}{56, 93, 68, 56, 87, 68, 56, 68},
			68,
			"56 -> 93 -> 56 -> 87 -> 56 -> <nil>",
		},
		{
			[]interface{}{"mvkUxbqhQve4l", "4Zc4t hnf SQ", "q2If E8BPuX"},
			"4Zc4t hnf SQ",
			"mvkUxbqhQve4l -> q2If E8BPuX -> <nil>",
		},
	}

	for _, table := range tables {
		link := &student.List{}

		for _, data := range table.list {
			student.ListPushBack(link, data)
		}

		student.ListRemoveIf(link, table.dataRef)

		res := getList(link)

		if res != table.answer {
			t.Errorf("ListRemoveIf (%d, %d): [your answer - (%s)] != [(%s) - test's answer]", table.list, table.dataRef, res, table.answer)
		}
	}
}

func getList(l *student.List) string {
	list := ""
	it := l.Head
	for it != nil {
		list += fmt.Sprintf("%v", it.Data) + " -> "
		it = it.Next
	}
	list += "<nil>"
	return list
}
