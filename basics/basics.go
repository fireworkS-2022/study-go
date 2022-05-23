package basics

import "fmt"
import "math"

func HelloGo() {
	fmt.Println("Hello! Golang")
}

// 引数なし、返り値なし
func MathCal() {
	fmt.Println(math.Pi)
}

// 引数あり、戻り値なし
func Add(x int, y int) {
	fmt.Println(x + y)
}

// 引数あり、戻り値あり
func Add2(x, y int) int {
	return(x + y)
}

