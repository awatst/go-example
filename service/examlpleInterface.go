package service

import "fmt"

type Event interface {
	Speak() string
	Walk() string
}

type Dog struct {
	Name string
}

type Person struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " Says Woof!"
}

func (p Person) Speak() string {
	return p.Name + " Says Hi!"
}

func (d Dog) Walk() string {
	return d.Name + " Walk 4 legs"
}

func (p Person) Walk() string {
	return p.Name + " Walk 2 legs"
}

func MakeSound(e Event) {
	fmt.Println(e.Speak())
}

func TakeAWalk(e Event) {
	fmt.Println(e.Walk())
}
