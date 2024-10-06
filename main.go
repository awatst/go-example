package main

import (
	"fmt"

	"github.com/awatst/go-example/service"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Print("main_test \n")
	fmt.Printf("main_uuid: %s \n", id)
	service.ExampleService()
	a, b := service.SwapString("1", "2")
	fmt.Printf("main_swapString: a: %s, b: %s", a, b)
	exampleSlice()
}

func exampleSlice() {
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println("======================")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySubSlice := mySlice[1:3] //select start at index 1 and end at index 3
	fmt.Println("subslice: ", mySubSlice)
	fmt.Println("len subslice: ", len(mySubSlice))
	fmt.Println("cap subslice: ", cap(mySubSlice))
}
