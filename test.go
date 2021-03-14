package main
import (
	"fmt"
	"myMath"
	"runtime"
)

//结构类
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

//输出类
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func main() {

	fmt.Println(runtime.Version())
	//字符串加减
	fmt.Println("Hello World!")
	fmt.Println(mathClass.Add(1, 1))
	fmt.Println(mathClass.Sub(1, 1))

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	var bb = "aaa"
	fmt.Println("a++" + bb)

	//测试
	var ra string = "Runoob"
	fmt.Println(ra)

	var b, c int = 1, 2
	fmt.Println(b, c)

	//测试数据类型
	var ii int
	var f float64
	var dd bool
	var s string
	fmt.Printf("%v %v %v %q\n", ii, f, dd, s)

	//再测试
	intVal1, intVal2 := 4, 100 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	fmt.Println(intVal1, intVal2)

	//侧额侧额

	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs)
	ccc := "abc"
	fmt.Println("hello, world", ccc)

	fmt.Println("---------测试数组------------")
	//测试数组
	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	/* 定义局部变量 */
	var _a int = 100
	var _b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(_a, _b)
	fmt.Printf("最大值是 : %d\n", ret)

	//
	aaaa, bbbb := swap("Google", "Runoob")
	fmt.Println(aaaa, bbbb)

	//

	fmt.Printf("a变量的地址: %x\n", &ra)

	//
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)

	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)
}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

/* 交换  */

func swap(x, y string) (string, string) {
	return y, x
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
