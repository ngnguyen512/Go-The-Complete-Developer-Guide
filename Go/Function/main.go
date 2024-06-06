package main

import (
	"errors"
	"fmt"
)

// function:A function is a block of statements that can be used repeatedly in a program.
// struct with function
var ErrDivisionByZero = errors.New("division by zero is not allowed")

type myStruct struct {
	FirstName string
	Age       int
}

// Function with Parameter Example
func firstname(fname string, age int) {
	fmt.Println("Hello", fname, age, "years old")
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
} // type myStruct has an associate function with it.

// Multiple returns

func main() {
	var myVar myStruct
	myVar.FirstName = "John"
	myVar.Age = 20
	myVar2 := myStruct{
		FirstName: "Mary",
		Age:       15,
	}
	fmt.Println("myVar is set to", myVar.printFirstName())
	fmt.Println("myVar is set to", myVar2.printFirstName())
	firstname(myVar.printFirstName(), myVar.Age)
	firstname(myVar2.printFirstName(), myVar2.Age)

	//Error handling
	result, err := divide(10, 0)

	if err == ErrDivisionByZero {
		fmt.Println("Error: Division by zero")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}
