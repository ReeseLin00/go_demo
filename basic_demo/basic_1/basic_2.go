package main

import "fmt"

func main() {

	fmt.Println("请输入一个年龄：")
	var age int
	fmt.Scan(&age)
	//if age <= 0 {
	//	fmt.Println("未出生")
	//	return
	//}
	//if age <= 18 {
	//	fmt.Println("未成年")
	//	return
	//}
	//if age <= 35 {
	//	fmt.Println("青年")
	//	return
	//}
	//fmt.Println("壮年")

	//if age <= 18 {
	//	if age <= 0 {
	//		fmt.Println("未出生")
	//	} else {
	//		fmt.Println("未成年")
	//	}
	//} else {
	//	if age <= 35 {
	//		fmt.Println("青年")
	//	} else {
	//		fmt.Println("中年")
	//	}
	//}

	//if age <= 0 {
	//	fmt.Println("未出生")
	//}
	//if age > 0 && age <= 18 {
	//	fmt.Println("未成年")
	//}
	//if age > 18 && age <= 35 {
	//	fmt.Println("青年")
	//}
	//if age > 35 {
	//	fmt.Println("中年")
	//}

	//	Switch
	//如果希望它输出满足的所有条件 添加fallthrough
	switch {
	case age <= 0:
		fmt.Println("未出生")
		fallthrough
	case age <= 18:
		fmt.Println("未成年")
		fallthrough
	case age <= 35:
		fmt.Println("青年")
		fallthrough
	default:
		fmt.Println("中年")
	}
}
