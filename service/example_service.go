package service

import (
	"fmt"

	"github.com/awatst/go-example/service/internal/internal_service"
)

func callInternalService() {
	internalservice.TestInternalService()
}

func ExampleService() {
	fmt.Println("run ExampleService")
	callInternalService()
}
