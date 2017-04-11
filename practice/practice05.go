package main

import "fmt"

type person struct {
	first string
}

func (p person) speak(){
	fmt.Println(p)
}

func main() {
	p1 := person{"vince is"}
	p1.speak()

	xi := []int{30,31,32}
	for i, v := range xi{
		fmt.Println(i, v)
	}

	m := map[string]int{"going on": 33}
	for k, v := range m{
		fmt.Println(k, v)
	}
}
