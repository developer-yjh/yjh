package main

import (
	"fmt"
)

func main() {
	msg := "我们都是中国人呀, 中国龙!"
	PrintMessage(msg)
	fmt.Println("this is test main func")

	fmt.Println("this is main.go")
}

func PrintMessage(msg string) {
	fmt.Println("this is message : ", msg)
}