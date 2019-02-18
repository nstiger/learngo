/*
1、strings包，prefix，suffix，contains，判断字符串前缀，后缀和包含子串

strings包中还有一个常用于包含判断的函数ContainsAny()，能够匹配Unicode字符
2、strings索引
strings.index()返回指定字符/串的第一个字符的索引值，如果不存在相应的字符串则返回-1

3、如果是非ASCII字符，可以用IndexRune函数来对字符进行定位

4 、用Replace函数替换字符串
strings.Replace()

5、 统计
5.1出现频率：统计指定字符串出现的频率 strings.Count(str, manyO)
5.2字符数量：统计字符串的字符数量 len([]rune(str))

6、大小写转换
ToLower将字符串中的Unicode字符全部转换为相应的小写字符
ToUpper将字符串中的Unicode字符全部转换为相应的大写字符

7、修剪
Trim(string, string)实际上是把字符串转换成rune之后操作的
strings.Trim 要注意只修剪匹配开头和结尾的字符串
strings.TrimLeft
strings.TrimSpace

8、分割
strings.Split函数返回一个切片

9、插入字符
strings.Join用于将元素类型为string的slice使用分隔符号拼接组成一个字符串


*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("Does the string \"%s\" have suffix %s? \n", str, "string")
	fmt.Printf("%t\n", strings.HasSuffix(str, "string"))
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", str, "example")
	fmt.Printf("%t\n", strings.Contains(str, "example"))

	fmt.Println("====================  strings.index  ========================")
	str1 := "Hi, i'm Job, Hi."
	fmt.Println(str1)
	fmt.Printf("The position of \"Job\" is: ")
	fmt.Printf("%d\n", strings.Index(str1, "Job"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str1, "Hi"))

	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str1, "Hi"))

	fmt.Printf("The position of \"Tim\" is: ")
	fmt.Printf("%d\n", strings.Index(str1, "Tim"))

	fmt.Println()
	fmt.Println("====================  strings.IndexRune  ========================")
	str2 := "大王叫我来巡山"
	fmt.Println(str2)
	fmt.Printf("The position of \"来\" is: ")
	fmt.Printf("%d\n",strings.IndexRune(str2, '来'))

	fmt.Printf("\n\n")
	fmt.Println("=====================  strings.Replace =========================")
	str3 := "你好世界，这个世界真好。"
	fmt.Println(str3)
	new := "地球"
	old := "世界"
	n := 1
	fmt.Println(strings.Replace(str3, old, new, n))
	//如果n为-1, 则匹配所有

	fmt.Printf("\n\n")
	fmt.Println("=====================  strings.Count =========================")
	str4 := "Golang is cool, right?"
	fmt.Println(str4)
	var manyO = "o"
	fmt.Printf("%d\n", strings.Count(str4, manyO))
	fmt.Printf("%d\n", strings.Count(str4, "oo"))

	fmt.Println("字符数量统计")
	fmt.Printf("str3共有 %d 个字符\n", len([]rune(str3)))
	fmt.Printf("str4 %d\n", len([]rune(str4)))

	//fmt.Println(utf8.RuneCountInString(str))
	//通过引入unicode/utf8包的统计方式

	fmt.Printf("\n\n")
	fmt.Println("=====================  strings.Trim =========================")
	fmt.Println("## Trim函数的用法")
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", "! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", " ! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Golang !!! ", "!"))

	fmt.Println("## TrimLeft 函数的用法")
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Golang !!! ", "! "))
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Golang !!! ", " ! "))
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Golang !!! ", "!"))
	fmt.Println()

	fmt.Println("## TrimSpace 函数的用法")
	fmt.Println(strings.TrimSpace(" \t\n 这是\t一句话 \n\t\r\n"))

	fmt.Println(" #utf8")
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "今"))
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "天"))
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "今天"))

	fmt.Printf("\n\n")
	fmt.Println("=====================  strings.Split =========================")
	fmt.Println("分割函数Split的用法")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a boy a girl a dog a cat", "a "))	//a空格，否则把cat分割了
	fmt.Printf("%q\n", strings.Split("xyz ", ""))

	fmt.Printf("\n\n")
	fmt.Println("=====================  strings.Join =========================")
	fmt.Println("插入字符函数join的用法")

	str5 := "The quick brown fox jumps over the lazy dog 中文"
	strSli := strings.Fields(str5)
	//strings.Fields()用于把字符串转换成字符串切片，然后通过range获得每个切片的值。
	fmt.Printf("%s\n", strSli)
	for _, val := range strSli {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
	str6 := strings.Join(strSli, ";")
	fmt.Printf("%s\n", str6)

	

}