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

	node := BTreeSearchItem(root, "02")
	fmt.Println("Before delete:")
	BTreeApplyInorder(root, fmt.Println)
	root = BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	BTreeApplyInorder(root, fmt.Println)
}

// func main() {
// 	root := &TreeNode{Data: "33"}
// 	BTreeInsertData(root, "20")
// 	BTreeInsertData(root, "5")
// 	BTreeInsertData(root, "13")
// 	BTreeInsertData(root, "31")
// 	BTreeInsertData(root, "52")
// 	BTreeInsertData(root, "11")

// 	node := BTreeSearchItem(root, "20")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

// func main() {
// 	root := &TreeNode{Data: "03"}
// 	BTreeInsertData(root, "39")
// 	BTreeInsertData(root, "11")
// 	BTreeInsertData(root, "99")
// 	BTreeInsertData(root, "14")
// 	BTreeInsertData(root, "44")
// 	BTreeInsertData(root, "11")

// 	node := BTreeSearchItem(root, "03")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

// func main() {
// 	root := &TreeNode{Data: "03"}
// 	BTreeInsertData(root, "01")
// 	BTreeInsertData(root, "03")
// 	BTreeInsertData(root, "94")
// 	BTreeInsertData(root, "19")
// 	BTreeInsertData(root, "111")
// 	BTreeInsertData(root, "24")

// 	node := BTreeSearchItem(root, "03")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)

	} else {
		if root.Left == nil {
			//
			a := root.Right
			root = nil
			return a
		} else if root.Right == nil {
			a := root.Left
			root = nil
			return a
		}
		a := BTreeMin(root.Right)

		root.Data = a.Data
		root.Right = BTreeDeleteNode(root.Right, a)
	}
	return root
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
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

func BTreeSearchItem(root *TreeNode, data string) *TreeNode {

	if root != nil {
		if data < root.Data {
			return BTreeSearchItem(root.Left, data)
		}

		if data > root.Data {
			return BTreeSearchItem(root.Right, data)
		}
	}

	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}
