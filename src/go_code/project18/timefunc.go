package main
import (
	"fmt"
	"time"
)
func main() {
	now := time.Now()
	fmt.Printf("now %v, now的数据类型 %T", now , now)

	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n",now.Month())
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
	//fmt.Printf("年=%v"，now.Year())

	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(),now.Month(),now.Day(),
	now.Hour(),now.Minute(),now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(),now.Month(),now.Day(),
	now.Hour(),now.Minute(),now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)

	fmt.Printf(now.Format("2006/01/02 15:04:05")) //格式化日期时间，参考时间必须是2006/01/02 15:04:05
	fmt.Println()
	fmt.Printf(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Println(now.Format("15:04:05"))
	fmt.Println()

	fmt.Printf(now.Format("01"))
	fmt.Println()
	
}