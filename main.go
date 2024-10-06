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
	exampleMap()
}

func exampleSlice() {
	mySlice := []int{10, 20, 30, 40, 50}

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySubSlice := mySlice[1:3] //select start at index 1 and end at index 3
	fmt.Println("======================")
	fmt.Println("subslice: ", mySubSlice)
	fmt.Println("len subslice: ", len(mySubSlice))
	fmt.Println("cap subslice: ", cap(mySubSlice)) //check memory, slice keep index 1 to end
}

func exampleMap() {
	myMap := make(map[string]int)

	myMap["Ford"] = 10
	myMap["Honda"] = 20
	myMap["Toyota"] = 30
	myMap["Subaru"] = 40

	delete(myMap, "Toyota")

	fmt.Println("======================")
	fmt.Println("Ford: ", myMap["Ford"])
	fmt.Println("Toyota: ", myMap["Toyota"])

	for key, value := range myMap {
		// fmt.Printf("%s -> %d\n", key, value)
		if key == "Subaru" {
			fmt.Printf("%s -> %d\n", key, value)
		}
	}

	//check if key is exists
	val, exists := myMap["Toyota"]
	if exists {
		fmt.Println("Toyota: ", val)
	} else {
		fmt.Println("Toyota key not exists")
	}

}
