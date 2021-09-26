package tree

import (
	"fmt"
)

/**
golang 没有 class 没有继承，多态的概念
golang 只有结构体
 */
type TreeNode struct{
	Value int
	Left,Right *TreeNode
}

//为结构体 定义 方法
//显示 定义 和 命名方法接收者
// (node treeNode)为接收者
//这里也是传值
func (node TreeNode) Print(){
	fmt.Print(node.Value," ")
}

//只有使用指针才可以改变结构内容
//nil指针也可以调用方法
//这里是传递地址
func (node *TreeNode) SetValue(Value int){
	if node == nil {
		fmt.Println("setting Value to  nil node ! Ignored.")
		return
	}
	node.Value = Value
}

//结构分配在 栈或者 堆  是由 编译器自己来决定的
//工厂函数，自己控制创建结构体
func CreateNode(Value int) *TreeNode{
	//注意在 golang中是可以允许返回 局部变量的地址
	//而在 C++中不可以（报错 <局部变量的生命周期是在函数内>）
	return &TreeNode{Value:Value}
}


func treeFunc(){
	var root TreeNode
	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5,nil,nil}
	//创建一个空的treeNode 返回地址
	root.Right.Left = new(TreeNode)
	root.Left.Right = CreateNode(2)

	//值接收者，拷贝一份值
	root.Print()
	fmt.Println()

	root.Left.Right.Print()
	fmt.Println()
	//golang 在调用时在取地址或是值的时候 做一些人性化的设计不用区分
	//指针接收者，传递指针地址
	root.Left.Right.SetValue(4)
	root.Left.Right.Print()

	fmt.Println()
	root.Print()
	root.SetValue(100)

	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()
	//fmt.Println()

	//nil
	var pRoot1 *TreeNode
	pRoot1.SetValue(3)

	root.Left.Right.SetValue(2)
	root.SetValue(3)

	fmt.Println()

	root.Traverse()


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




