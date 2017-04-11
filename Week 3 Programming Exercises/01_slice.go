package main

import (
	"fmt"
)

func main() {
	x := []int{7, 9, 42}
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}
}
