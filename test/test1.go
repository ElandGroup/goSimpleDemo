package test

import (
	"fmt"
)

func TestBase() {
	v := ifOne(1)
	fmt.Println(v)
	TestSlice()
	TestMap()
}

func ifOne(x int) string {
	r := "greater than 10"
	if x > 10 {

	} else if x > 5 {
		r = "greater than 5"
	} else {
		r = "less than 5"
	}
	return r
}

func TestSlice() {
	J := []string{"Java", "C#", "Go"}
	for k, v := range J {
		fmt.Println(k, v)
	}
}

func TestMap() {
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	for k, v := range rating {
		fmt.Println(k, v)
	}
}
