package main

import (
	"fmt"
)

func main() {
	x := map[string]int{"Vince": 32, "Roxi": 31}
	for k, _ := range x {
		fmt.Println(k, "-", x)
	}
}
