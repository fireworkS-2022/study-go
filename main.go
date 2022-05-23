package main

import "go-study/basics"
import "fmt"

func main() {
	// helloパッケージの関数を呼び出し
	basics.HelloGo()
	basics.MathCal()
	basics.Add(4, 7)
	fmt.Println(basics.Add2(6, 9))
}