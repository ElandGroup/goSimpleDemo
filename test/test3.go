package test

type IStudent interface {
	SayHi()
}

func TestInterface() {
	s := Student{Name: "xiaomiao", Age: 12}
	s.SayHi()

	t := Teacher{Name: "damiao", Age: 21}
	t.SayHi()
}
