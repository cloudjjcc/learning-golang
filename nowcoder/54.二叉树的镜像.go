package nowcoder

//题目描述
//操作给定的二叉树，将其变换为源二叉树的镜像。

type treeNode struct {
	left  *treeNode
	right *treeNode
	Value interface{}
}

func mirror(tree *treeNode) {
	if tree == nil {
		return
	}
	ntree := new(treeNode)
	ntree.left, ntree.right = ntree.right, ntree.left
	mirror(ntree.left)
	mirror(ntree.right)
}
