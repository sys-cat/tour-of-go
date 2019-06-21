package main

import (
	"fmt"
)

func fibonacci() func() int {
	f, g := 0, 1
	return func() int {
		f, g = g, f+g // 前の値と新しい値を扱う定形
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
		// ↑初期条件(a0 = 0, a1 = 1) さえ有れば引数を取ずに数列が生成出来る
	}
}
