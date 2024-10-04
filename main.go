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
}
