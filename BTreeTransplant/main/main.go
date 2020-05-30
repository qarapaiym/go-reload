package main

import (
	"fmt"

	student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}
