package main

import (
	"fmt"

	"github.com/awatst/go-example/service"
	"github.com/google/uuid"
)

type Student struct {
	Name     string
	LastName string
	Weight   int
	Height   int
	Grade    string
	Address  Address
}

type Address struct {
	Street  string
	City    string
	ZipCode int
}

func main() {
	id := uuid.New()
	fmt.Print("main_test \n")
	fmt.Printf("main_uuid: %s \n", id)
	service.ExampleService()
	a, b := service.SwapString("1", "2")
	fmt.Printf("main_swapString: a: %s, b: %s", a, b)
	exampleSlice()
	exampleMap()
	exampleStruct()

	//method
	studentFullName := Student{Name: "John", LastName: "Snow"}
	fmt.Println("FullName: ", studentFullName.exampleMethod())

	//interface
	fmt.Println("Dog ", service.Dog.Speak(service.Dog{}))
	fmt.Println("Person ", service.Person.Speak(service.Person{}))

	dog := service.Dog{Name: "Daeng"}
	person := service.Person{Name: "Sompong"}
	service.MakeSound(dog)
	service.MakeSound(person)
	service.TakeAWalk(dog)
	service.TakeAWalk(person)

	//pointer
	//pointer is like Ref in oop
	service.ExamplePointer()
	emp := service.Employee{Name: "A", Salary: 100000}
	service.GiveRise(&emp, 5000)
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

func exampleStruct() {
	var student Student
	student.Name = "A"
	student.Weight = 10
	student.Height = 20
	student.Grade = "A"
	student.Address.Street = "aa1"
	student.Address.City = "cc1"
	student.Address.ZipCode = 1234

	fmt.Println("Student: ", student)

	var studentArray [3]Student
	studentArray[0].Name = "AA"
	studentArray[0].Weight = 100
	studentArray[0].Height = 200
	studentArray[0].Grade = "XX"
	studentArray[0].Address.Street = "a1"
	studentArray[0].Address.City = "c1"
	studentArray[0].Address.ZipCode = 123

	studentArray[1].Name = "BBB"
	studentArray[1].Weight = 300
	studentArray[1].Height = 400
	studentArray[1].Grade = "YY"
	studentArray[1].Address.Street = "a2"
	studentArray[1].Address.City = "c2"
	studentArray[1].Address.ZipCode = 456

	studentArray[2].Name = "CCC"
	studentArray[2].Weight = 500
	studentArray[2].Height = 600
	studentArray[2].Grade = "ZZ"
	studentArray[2].Address.Street = "a2"
	studentArray[2].Address.City = "c2"
	studentArray[2].Address.ZipCode = 789

	for i := 0; i < len(studentArray); i++ {
		fmt.Println("Student no.", i+1)
		fmt.Println("Name: ", studentArray[i].Name)
		fmt.Println("Weight: ", studentArray[i].Weight)
		fmt.Println("Height: ", studentArray[i].Height)
		fmt.Println("Grade: ", studentArray[i].Grade)
		fmt.Println("Street: ", studentArray[i].Address.Street)
		fmt.Println("City: ", studentArray[i].Address.City)
		fmt.Println("ZipCode: ", studentArray[i].Address.ZipCode)
	}

	studentMap := make(map[string]Student)
	studentMap["std01"] = Student{Name: "Somchai", Height: 1, Weight: 2, Grade: "A"}
	studentMap["std02"] = Student{Name: "Somsak", Height: 3, Weight: 4, Grade: "B"}
	studentMap["std03"] = Student{Name: "Somsri", Height: 5, Weight: 6, Grade: "C"}
	studentMap["std04"] = Student{Name: "Sompong", Height: 7, Weight: 8, Grade: "D"}

	for key, val := range studentMap {
		fmt.Printf("Id: %s, Name: %s, Height: %d, Weight: %d, Grade: %s \n", key, val.Name, val.Height, val.Weight, val.Grade)
	}
}

func (s Student) exampleMethod() string {
	return s.Name + " " + s.LastName
}
