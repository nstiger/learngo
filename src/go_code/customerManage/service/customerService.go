package service
import (
	"go_code/customerManage/model"
)

//定义一个结构体，完成对customer的增删改查操作
type CustomerService struct{
	//定义一个切片customers，它的类型是从model引入的Customer
	customers []model.Customer
	//定义一个字段，记录customer的记录数,还可以用来作为Id
	customerNum int
}

//定义一个方法，用来返回一个CustomerService结构体类型的指针
func NewCustomerService() *CustomerService {
	//初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "118", "zs@sohu.com")
	//用append方法将初始化的customer加入到customers切片中去
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
} 

//添加客户到customer切片
//返回客户切片
func (this *CustomerService) Add(customer model.Customer) bool {

	//处理Id，确定id的添加规则，就是添加顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i:=0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//找到了
			index = i
		}
	}
	return index
}

//
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	//如果index = -1,说明没有这个客户
	if index == -1 {
		return false
	}

	//如何从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//完成customerService结构体的Update方法
func (this *CustomerService) Update(id int, name string, gender string,
	 age int, phone string, email string) bool {

	//1.找到id号对应的索引（切片下标）
	index := this.FindById(id)
	if index == -1 {   //-1退出本函数     
		return false   
	}
	//2.根据切片下标修改对应的字段
	this.customers[index].Name = name
	this.customers[index].Gender = gender
	this.customers[index].Age = age
	this.customers[index].Phone = phone
	this.customers[index].Email = email 

	return true

}