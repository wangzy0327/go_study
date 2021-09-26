package main

import (
	"fmt"
	"go-demo/tree"
)

/**
golang 大写是public 小写是private
golang 文件夹名并不一定是 包名

 */

//这里采用 组合 来扩展 TreeNode 的功能
type MyTreeNode struct{
	node *tree.TreeNode
}

//扩展的 后序遍历的新功能
func (myNode *MyTreeNode) postOrder(){
	//myNode.node 防止 MyTreeNode包装一个空的node
	if myNode == nil || myNode.node == nil{
		return
	}
	left := MyTreeNode{myNode.node.Left}
	right := MyTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()

	myNode.node.Print()
}

func main(){
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5,nil,nil}
	//创建一个空的treeNode 返回地址
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	root.Traverse()
	fmt.Println()
	fmt.Println("=====================postOrder=======================")
	myNode := MyTreeNode{&root}
	myNode.postOrder()

	/**
						3
	                  /    \
	                  0     5
	                  \    /
	                   2  0

	*/

	//nodes := []treeNode{
	//	{Value:3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)
	/**
	[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000004480}]
	*/
}
