package main

import (
	"fmt"
)

type person struct {
	fname string
	lname   string
}

func main() {
	s := person{
		"Vincent",
		"Bailey",
	}
	fmt.Println(s)
}