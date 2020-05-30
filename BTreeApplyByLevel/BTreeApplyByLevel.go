package main

import "fmt"

func main() {
	root := &TreeNode{Data: "01"}
	BTreeInsertData(root, "07")
	BTreeInsertData(root, "05")
	BTreeInsertData(root, "12")
	BTreeInsertData(root, "02")
	BTreeInsertData(root, "10")
	BTreeInsertData(root, "03")

	BTreeApplyByLevel(root, fmt.Println)
}

//func main() {
// 	root := &TreeNode{Data: "01"}
// 	BTreeInsertData(root, "07")
// 	BTreeInsertData(root, "05")
// 	BTreeInsertData(root, "12")
// 	BTreeInsertData(root, "02")
// 	BTreeInsertData(root, "10")
// 	BTreeInsertData(root, "03")

// 	BTreeApplyByLevel(root, fmt.Print)
// }

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	height := BTreeLevelCount(root)
	for level := 1; level <= height; level++ {
		BTreeApplyToGivenLevel(root, level, f)
	}
}

func BTreeApplyToGivenLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		BTreeApplyToGivenLevel(root.Left, level-1, f)
		BTreeApplyToGivenLevel(root.Right, level-1, f)
	}
}

func BTreeLevelCount(root *TreeNode) int {
	countl := 0
	countr := 0

	if root != nil {
		countl = BTreeLevelCount(root.Left)
		countl++
		countr = BTreeLevelCount(root.Right)
		countr++
	}
	if countl > countr {
		return countl
	}
	return countr
}

func BTreeInsertData(root *TreeNode, data string) {
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
}
