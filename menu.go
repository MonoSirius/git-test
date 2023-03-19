package main

import (
	"fmt"
)

func main() {
	var cmd string
	for {
		fmt.Println("请输入命令：")
		fmt.Scanln(&cmd)
		switch cmd {
		case "help":
			fmt.Println("--Help")
		case "quit":
			fmt.Println("再见")
			return
		default:
			fmt.Println("错误命令")
		}
	}
}
