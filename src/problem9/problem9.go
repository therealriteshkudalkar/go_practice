package problem9

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeStack []*TreeNode

func (t TreeNodeStack) Push(node *TreeNode) TreeNodeStack {
	return append(t, node)
}

func (t TreeNodeStack) Pop() (TreeNodeStack, *TreeNode) {
	l := len(t)
	if l == 0 {
		return t, nil
	}
	return t[:l-1], t[l-1]
}

func KthSmallest(root *TreeNode, k int) int {
	// Perform inorder traversal until k is zero
	stack := make(TreeNodeStack, 0)

	currNode := root

	for currNode != nil || len(stack) != 0 {
		for currNode != nil {
			stack = stack.Push(currNode)
			currNode = currNode.Left
		}

		stack, currNode = stack.Pop()
		k -= 1
		if k == 0 {
			return currNode.Val
		}

		currNode = currNode.Right
	}
	return root.Val
}
