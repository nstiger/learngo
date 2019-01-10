package main
import (
	"fmt"
)

//定义一个猫猫的结构体
type CatNode struct {
	no int 
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加的第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head	//构成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//定义一个临时变量，帮忙，找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}


//输出这个环形链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表......")
		return
	}
	for {
		fmt.Printf("猫的信息为=[id=%d name=%s] ->\n", temp.no, temp.name )
		if temp.next == head { //说明到了环状结尾
			break
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表， 不能删除")
		return head
	}
	//如果只有一个结点
	if temp.next == head {	//只有一个结点
		temp.next = nil
		return head
	}

	//将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果有两个及以上结点
	flag := true 	//这里假设进入环形链表能够找到要删除的结点
	for {
		if temp.next == head {	//如果来到这里，说明比较到了最后一个[最后一个还没比较]
			break
		}
		if temp.no == id {
			if temp == head {	//说明删除的是头结点
				head = head.next
			}
			//说明找到了，可以直接删除
			helper.next = temp.next
			fmt.Printf("猫猫=%d被删除了\n", id)
			flag = false
			break
		}
		temp = temp.next 	//移动 [比较]
		helper = helper.next	//移动[一旦找到要删除的结点 helper]
	}
	//这里还要比较一次
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Println("对不起，没有no=\n", id)
		}
	}
	return head
}

func main() {
	//这里我们初始化一个环形链表的头结点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no : 1,
		name : "tom",
	}
	cat2 := &CatNode{
		no : 2,
		name : "tom2",
	}
	cat3 := &CatNode{
		no : 3,
		name : "tom3",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DelCatNode(head, 30)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	ListCircleLink(head)

}