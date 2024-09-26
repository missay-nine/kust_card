package main

import (
	"fmt"
	"os"
)

func main() {
	// 设置环境变量进行测试
	//os.Setenv("username1", "15197903439")

	a := os.Getenv("card_username")
	fmt.Println(a) // 应该打印 15197903439
}
