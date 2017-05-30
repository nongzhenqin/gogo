package main

import "fmt"
import "time"

//程序入口
func main()  {
	//fmt.Println("hello word!")
	test2()
}

/*
测试1
 */
func test1()  {
	a := 10
	var b int = 20
	c := a + b
	fmt.Println(c)
}

func test2()  {
	var a int = 0
	var b int64 = time.Now().UnixNano() / 1000000
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	var c int64 = time.Now().UnixNano() / 1000000
	fmt.Println(c - b)
	//fmt.Println(c)
}