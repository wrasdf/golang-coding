package main

import (
	"fmt"
	// "math"
)

// "errors"

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

func main() {
	// fmt.Printf("Hello, world.\n")

	//=========  s := `
	// 	Hello
	// 	World
	// `
	// fmt.Printf("%s\n", s)

	//=========  Demo of Errors
	// err := errors.New("================================")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	//=========  dynamic array
	// c := [...]int{4, 5, 6}
	// fmt.Println(c)

	//=========  Map
	// map[keyType]valueType

	// numbers := make(map[string]int)
	// numbers["one"] = 1
	// numbers["ten"] = 10
	// numbers["three"] = 3

	// fmt.Println(numbers["one"], numbers["ten"])

	//========= example of any type of Map
	// k := make(map[string]interface{})
	// k["asf"] = "asf"
	// k["asf2"] = 12
	//
	// fmt.Println(k)

	// make只能创建slice、map和channel

	process()
	method()
	structs()
	interfacesRunner()

}

func process() {

	// if x > 10 {
	// 	fmt.Println("x is greater than 10")
	// } else {
	// 	fmt.Println("x is less than 10")
	// }

	// sum := 0
	// for index := 0; index < 10; index++ {
	// 	sum += index
	// }
	// fmt.Println("sum is equal to ", sum)

	// for index := 10; index>0; index-- {
	// 	if index == 5{
	// 		break // or continue
	// 	}
	// 	fmt.Println(index)
	// }
	// // break prints 10、9、8、7、6
	// // continue prints 10、9、8、7、6、4、3、2、1

	// integer := 6
	// switch integer {
	// case 4:
	// 	fmt.Println("The integer was <= 4")
	// 	fallthrough
	// case 5:
	// 	fmt.Println("The integer was <= 5")
	// 	fallthrough
	// case 6:
	// 	fmt.Println("The integer was <= 6")
	// 	fallthrough
	// case 7:
	// 	fmt.Println("The integer was <= 7")
	// 	fallthrough
	// case 8:
	// 	fmt.Println("The integer was <= 8")
	// 	fallthrough
	// default:
	// 	fmt.Println("default case")
	// }

	// fallthrough强制执行后面的case代码。
	// outputs:
	// The integer was <= 6
	// The integer was <= 7
	// The integer was <= 8
	// default case

}

func functions() {

	// func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	// 	//这里是处理逻辑代码
	// 	//返回多个值
	// 	return value1, value2
	// }

	// 	//返回 A+B 和 A*B
	// func SumAndProduct(A, B int) (int, int) {
	// 	return A+B, A*B
	// }

	// Go函数支持变参
	// func myfunc(arg ...int) {}

	// points
	// func add1(a *int) int { // 请注意，
	// 	*a = *a+1 // 修改了a的值
	// 	return *a // 返回新值
	// }

	// func main() {
	// 	x := 3

	// 	fmt.Println("x = ", x)  // 应该输出 "x = 3"

	// 	x1 := add1(&x)  // 调用 add1(&x) 传x的地址

	// 	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	// 	fmt.Println("x = ", x)    // 应该输出 "x = 4"
	// }

	// 在defer后指定的函数会在函数退出前调用
	// func ReadWrite() bool {
	// 	file.Open("file")
	// 	defer file.Close()
	// 	if failureX {
	// 		return false
	// 	}
	// 	if failureY {
	// 		return false
	// 	}
	// 	return true
	// }

}

func importFunction() {
	// import(
	// 	. "fmt"
	// )
	// 这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名.
	// 也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")

	// import(
	// 	f "fmt"
	// )
	// 别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")

	// import (
	// 	"database/sql"
	// 	_ "github.com/ziutek/mymysql/godrv"
	// )
	// _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。

}

func method() {
	//====== type of array
	// m := []int{3,4}
	// m = append(m, 12)
	// fmt.Printf(m)

	// ====== example of type with map
	// type months map[string]int
	//
	// m := months{
	// 	"January":  31,
	// 	"February": 28,
	// 	"December": 31,
	// }
	//
	// fmt.Println("---------------", m["January"])

	// ====== example of type with map
	// type months map[string]string
	//
	// m := months{
	// 	"January":  "first month",
	// 	"February": "second month",
	// 	"December": "twelve month",
	// }
	//
	// // delete(m, "January")
	//
	// fmt.Println("---------------", m["February"])

	// ====== example of loop array
	// for i := 0; i < 5; i ++ {
	// 	fmt.Println(i)
	// }

	// ====== example2 of loop array
	// arr := []string{"a", "b", "c"}
	// for index, value := range arr {
	// 	fmt.Println("index: ", index, "value: ", value)
	// }

	// ====== example3 of loop map
	// map1 := map[string]string {
	// 	"a": "b",
	// 	"c": "d",
	// }
	//
	// for index, value := range map1 {
	// 	fmt.Println("index: ", index, "value: ", value)
	// }

}

func structs() {

	//====== example 1
	// type Person struct {
	// 	name string
	// 	age int
	// }
	//
	// mike := Person {
	// 	name: "Mike John",
	// 	age: 12,
	// }
	//
	// fmt.Println(mike)


	//====== example 2
	// type Rectangle struct {
	// 	width, height float64
	// }
	//
	// type Circle struct {
	// 	radius float64
	// }
	//
	// func (r Rectangle) area() float64 {
	// 	return r.width * r.height
	// }
	//
	// func (c Circle) area() float64 {
	// 	return c.radius * c.radius * math.Pi
	// }
	//
	// r1 := Rectangle{12, 2}
	// r2 := Rectangle{9, 4}
	// c1 := Circle{10}
	// c2 := Circle{25}
	//
	// fmt.Println("Area of r1 is: ", r1.area())
	// fmt.Println("Area of r2 is: ", r2.area())
	// fmt.Println("Area of c1 is: ", c1.area())
	// fmt.Println("Area of c2 is: ", c2.area())

}
//====== interface example1
// type Writer interface {
// 	write() (int, error)
// }
//
// type ConsoleWriter struct {
// 	content string
// }
//
// func (cw ConsoleWriter) write () (int, error) {
// 	s, err := fmt.Println(cw.content)
// 	return s, err
// }
//
// func interfacesRunner() {
// 	var w Writer = ConsoleWriter{
// 		content: "Update your string here",
// 	}
// 	w.write()
// }


//====== interface example2
type Writer interface {
	publicBook(string, string)
	deleteBook(string)
	listBooks() (map[string]string)
}

type Details struct {
	name string
	age int
	books map[string]string
}

func NewWriter(details Details) *Details {
	return &details
}

func (w *Details) publicBook(bookName string, description string) {
	w.books[bookName] = description
}

func (w *Details) deleteBook(bookName string) {
	delete(w.books, bookName)
}

func (w *Details) listBooks() (map[string]string){
	return w.books
}

func interfacesRunner() {

	var mike Writer = NewWriter(Details{
		name: "Mike Strus",
		age: 43,
		books: make(map[string]string),
	})

	mike.publicBook("The Sky1", "Sky colors")
	mike.publicBook("The Sky2", "Sky colors")
	mike.publicBook("The Sky3", "Sky colors")
	mike.deleteBook("The Sky2")

	fmt.Println(mike.listBooks())
}
