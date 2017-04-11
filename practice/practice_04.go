package main

import "fmt"

func main() {
	xi := []int{4, 3, 7, 5, 4}
	fmt.Println(xi)

	for i, v := range xi {
		fmt.Println(i, v)
	}

	m := map[string]int{"vince": 32, "roxi": 31}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
