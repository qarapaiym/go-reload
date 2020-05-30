package tests

import (
	"fmt"
	"testing"

	student "../BTreeTransplant"
)

func TestBTreeTransplant(t *testing.T) {
	tables := []struct {
		tree       []string
		search     string
		replacment []string
		f          func(...interface{}) (int, error)
		fstr       string
		answer     string
	}{
		{
			[]string{"01", "07", "12", "10", "05", "02", "03"},
			"12",
			[]string{"55", "60", "33", "12", "15"},
			fmt.Println,
			"fmt.Println",
			"01\n02\n03\n05\n07\n12\n15\n33\n55\n60\n",
		},
		{
			[]string{"33", "5", "52", "20", "31", "13", "11"},
			"20",
			[]string{"55", "60", "33", "12", "15"},
			fmt.Println,
			"fmt.Println",
			"12\n15\n33\n55\n60\n33\n5\n52\n",
		},
		{
			[]string{"03", "39", "99", "44", "11", "14", "11"},
			"11",
			[]string{"55", "60", "33", "12", "15"},
			fmt.Println,
			"fmt.Println",
			"03\n12\n15\n33\n55\n60\n39\n44\n99\n",
		},
	}

	for _, table := range tables {
		root := &student.TreeNode{Data: table.tree[0]}

		for _, data := range table.tree[1:] {
			student.BTreeInsertData(root, data)
		}

		node := student.BTreeSearchItem(root, table.search)

		replacement := &student.TreeNode{Data: table.replacment[0]}

		for _, replData := range table.replacment[1:] {
			student.BTreeInsertData(replacement, replData)
		}

		root = student.BTreeTransplant(root, node, replacement)

		res := CaptureOutput(func() {
			student.BTreeApplyInorder(root, table.f)
		})

		if res != table.answer {
			t.Errorf("BTreeTransplant (%s, %s, %s, %s): [your answer - (%s)] != [(%s) - test's answer]", table.tree, table.search, table.replacment, table.fstr, res, table.answer)
		}
	}
}
