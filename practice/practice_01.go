package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println(p)
}

func main() {
	p1 := person{"vince"}
	p1.speak()

	xi := []int{2, 4, 6, 8}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	m := map[string]int{"roxi": 31}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
