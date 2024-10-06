package service

import (
	"fmt"

	"github.com/awatst/go-example/internal/internalService"
)

func callInternalService() {
	internalService.TestInternalService()
}

func ExampleService() {
	fmt.Println("run ExampleService")
	callInternalService()
}
