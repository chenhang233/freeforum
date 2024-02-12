package main

import "fmt"

type A struct {
	id int
}

func (a *A) SetId(id int) {
	a.id = id
}

func (a *A) GetId() int {
	return a.id
}

func main() {
	a := A{}
	a.SetId(10)
	fmt.Println(a.GetId())
}
