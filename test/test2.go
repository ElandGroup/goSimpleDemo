package test

import (
	"fmt"
)

func TestStruct() {
	p1 := new(Student)
	p1.Name = "xiaomiao1"
	p1.Age = 11
	fmt.Println(p1)

	p := Student{Name: "xiaomiao", Age: 12}
	fmt.Println(p)
	p.SayHi()

}

type Student struct {
	Name string
	Age  int
}

func (s *Student) SayHi() {
	h := fmt.Sprintf("I am a student,my name is %s:Good morning", s.Name)
	fmt.Printf("\n%#v", h)
}

type Teacher struct {
	Name string
	Age  int
}

func (s *Teacher) SayHi() {
	h := fmt.Sprintf("I am a teacher,my name is %s:Good morning", s.Name)
	fmt.Printf("\n%#v", h)
}
