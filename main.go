package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	fmt.Print("test \n")
	fmt.Printf("uuid: %s", id)
}
