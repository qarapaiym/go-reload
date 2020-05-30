package tests

import (
	"fmt"
	"testing"

	student "../BTreeApplyByLevel"
)

func TestBTreeApplyByLevel(t *testing.T) {
	tables := []struct {
		tree   []string
		f      func(...interface{}) (int, error)
		fstr   string
		answer string
	}{
		{
			[]string{"01", "07", "12", "10", "05", "02", "03"},
			fmt.Println,
			"fmt.Println",
			"01\n07\n05\n12\n02\n10\n03\n",
		},
		{
			[]string{"01", "07", "12", "10", "05", "02", "03"},
			fmt.Print,
			"fmt.Print",
			"01070512021003",
		},
	}

	for _, table := range tables {
		root := &student.TreeNode{Data: table.tree[0]}

		for _, data := range table.tree[1:] {
			student.BTreeInsertData(root, data)
		}

		res := CaptureOutput(func() {
			student.BTreeApplyByLevel(root, table.f)
		})

		if res != table.answer {
			t.Errorf("BTreeApplyByLevel (%s, %s,): [your answer - (%s)] != [(%s) - test's answer]", table.tree, table.fstr, res, table.answer)
		}
	}
}
