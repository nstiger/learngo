/*
有一个公司，当有新的员工来报到时，要求将该员工的信息加入（id，姓名，性别，年龄，住址...），
当输入该员工的id时，要求查找到该员工的所有信息
要求：不使用数据库，尽量节省内存，速度越快越好
添加时，保证按照雇员的id从低到高插入
*/

package main
import (
	"fmt"
)

type Emp struct{
	Id int
	Name string
	Next * Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d\n", this.Id % 7, this.Id)
}

type EmpLink struct{
	Head *Emp
}
//添加员工的方法,保证添加时编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head //辅助指针
	var pre *Emp = nil //辅助指针变量， pre在cur之前
	if cur == nil {
		//如果当前链表就是一个空链表
		this.Head = emp
		return
	}
	//如果不是空链表，给emp找到合适的位置插入
	//思路：让cur与emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出for循环时，把雇员信息插入
	pre.Next = emp
	emp.Next = cur
}

//显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d为空\n", no)
		return
	}
	//如果当前链表不为空，遍历当前链表，显示所有结点
	cur := this.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据id查找对应的雇员，如果没有就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

type HashTable struct{
	LinkArr [7]EmpLink
}

//给HashTable编写添加雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	//使用散列算法，决定将雇员添加到哪一个链表
	linkno := this.HashFun(emp.Id)
	this.LinkArr[linkno].Insert(emp)
}

//编写方法，显示hashtable的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//编写散列算法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到对应链表的下标
}

//编写一个方法，完成查找
func (this *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("============雇员系统菜单=================")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员Id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id : id,
				Name : name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//编写一个方法，显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			return
		default:
			fmt.Println("输入错误")
			
		}


	}
}