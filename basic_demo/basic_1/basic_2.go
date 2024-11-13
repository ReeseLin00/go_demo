package main

import "fmt"

func main() {

	fmt.Println("请输入一个年龄：")
	var age int
	fmt.Scan(&age)
	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age <= 35 {
		fmt.Println("青年")
		return
	}
	fmt.Println("壮年")
}
