package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	res := new(TreeNode)
	root := preorder[0]
	bound := -1
	for k, v := range inorder {
		if v == root {
			bound = k
			break
		}
	}

	res.Left = buildTree(preorder[1:bound+1], inorder[0:bound])
	res.Right = buildTree(preorder[bound+1:], inorder[bound+1:])
	res.Val = root
	return res
}
