/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	for i := range inorder {
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1], inorder[0:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:])}

		}
	}
	return nil
}