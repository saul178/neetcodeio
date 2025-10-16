package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return root
	}
}

func dfs(root *TreeNode, curSum int) int {
	if root == nil {
		return 0
	}

	curSum = curSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return curSum
	}
	return dfs(root.Left, curSum) + dfs(root.Right, curSum)
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}
