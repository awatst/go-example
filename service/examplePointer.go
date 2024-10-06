package service

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func ExamplePointer() {
	x := 10
	var p *int = &x

	fmt.Println("before re assign x = ", x)
	fmt.Println("before re assign p = ", p)

	*p = 20

	fmt.Println("after re assign x = ", x)
	fmt.Println("after re assign p = ", p)
}

func GiveRise(e *Employee, raise int) {
	e.Salary += raise
	fmt.Printf("Now %s salaly is %d", e.Name, e.Salary)
}
