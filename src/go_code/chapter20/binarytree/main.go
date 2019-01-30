/*
二叉数的三种遍历方法，前序、中序、后续
前序：依次遍历根结点、左子结点、右子结点
中序：依次遍历左子结点、根结点、右子结点
后序：依次遍历左子结点、右子结点、根结点
*/
package main
import (
	"fmt"
)

type Hero struct {
	No int
	Name string
	Left *Hero
	Right *Hero
}

func PreOrder(node *Hero) {
	if node != nil {

		fmt.Printf("结点的no=%d，name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func InfixOrder(node *Hero) {
	if node != nil {

		InfixOrder(node.Left)
		fmt.Printf("结点的no=%d，name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

func PostOrder(node *Hero) {
	if node != nil {

		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("结点的no=%d，name=%s\n", node.No, node.Name)
	}
}

func main() {

	root := &Hero{
		No : 1,
		Name : "宋江",
	}

	left1 :=&Hero{
		No : 2,
		Name : "吴用",
	}

	right1 :=&Hero{
		No : 3,
		Name : "卢俊义",
	}

	right2 :=&Hero{
		No : 4,
		Name : "林冲",
	}
	
	root.Left = left1
	root.Right = right1

	right1.Right = right2

	left10 :=&Hero{
		No : 10,
		Name : "tom",
	}

	right12 :=&Hero{
		No : 12,
		Name : "jerry",
	}

	left1.Left = left10
	left1.Right = right12

//	PreOrder(root)

//	InfixOrder(root)
	PostOrder(root)

}