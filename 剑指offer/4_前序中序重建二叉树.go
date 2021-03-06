package 剑指offer

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/26 上午1:41'

输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
*/

type TreeNode = leetcode_golang_version.TreeNode

func reConstructBinaryTree(preorder, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	if len(inorder) == 1 {
		return root
	}

	idx := getIndex(root.Val, inorder)

	root.Left = reConstructBinaryTree(preorder[1:idx+1], inorder[:idx])
	root.Right = reConstructBinaryTree(preorder[idx+1:], inorder[idx+1:])
}

func getIndex(val int, nums []int) int {
	for i, num := range nums {
		if num == val {
			return i
		}
	}

	return 0
}
