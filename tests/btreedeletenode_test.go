package tests

import (
	"fmt"
	"testing"

	student "../BTreeDeleteNode"
)

func TestBTreeDeleteNode(t *testing.T) {
	tables := []struct {
		tree   []string
		delete string
		f      func(...interface{}) (int, error)
		fstr   string
		answer string
	}{
		{
			[]string{"01", "07", "12", "10", "05", "02", "03"},
			"02",
			fmt.Println,
			"fmt.Println",
			"01\n03\n05\n07\n10\n12\n",
		},
		{
			[]string{"33", "5", "52", "20", "31", "13", "11"},
			"20",
			fmt.Println,
			"fmt.Println",
			"11\n13\n31\n33\n5\n52\n",
		},
		{
			[]string{"03", "39", "99", "44", "11", "14", "11"},
			"03",
			fmt.Println,
			"fmt.Println",
			"11\n11\n14\n39\n44\n99\n",
		},
		{
			[]string{"03", "03", "94", "19", "24", "111", "01"},
			"03",
			fmt.Println,
			"fmt.Println",
			"01\n03\n111\n19\n24\n94\n",
		},
	}

	for _, table := range tables {
		root := &student.TreeNode{Data: table.tree[0]}

		for _, data := range table.tree[1:] {
			student.BTreeInsertData(root, data)
		}

		node := student.BTreeSearchItem(root, table.delete)
		root = student.BTreeDeleteNode(root, node)

		res := CaptureOutput(func() {
			student.BTreeApplyInorder(root, table.f)
		})

		if res != table.answer {
			t.Errorf("BTreeDeleteNode (%s, %s, %s): [your answer - (%s)] != [(%s) - test's answer]", table.tree, table.delete, table.fstr, res, table.answer)
		}
	}
}
