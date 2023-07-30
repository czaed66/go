package main

import (
	"fmt"
	"strings"
)

const Logo = "Jessica老哥学习Go"

func main() {
	//1、打印数组
	arr1 := [...]int{11, 26, 3}
	fmt.Print(arr1[1])
	//换行
	fmt.Println()

	//2、循环遍历
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//3、字符串判断、截取
	var str = "hello world"
	fmt.Print(strings.Contains(str, "h"))
	fmt.Println()
	fmt.Print(str[0:4])
	fmt.Println()

	//3A、对象tostring
	user1 := User{"Jessica,", 26}
	fmt.Print(user1)
	fmt.Println()

	//5、构造函数打印对象
	user2 := NewUser("666", 18)
	fmt.Print(user2.name)
	fmt.Println()

	//6、调用一个方法
	fmt.Print(test())
}

func test() string {
	return "你好呀"
}

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{name: name, age: age}
}
