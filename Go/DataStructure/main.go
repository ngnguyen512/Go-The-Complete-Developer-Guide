package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	//ARRAY
	//Arrays are used to store multiple values of the same type in a single variable,
	//instead of declaring separate variables for each value.
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
	//access elements of an array
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
	//change elements of array
	prices[2] = 50
	fmt.Println(prices)
	//array initialization
	arrr1 := [5]int{}              //not initialized
	arrr2 := [5]int{1, 2}          //partially initialized
	arrr3 := [5]int{1, 2, 3, 4, 5} //fully initialized
	arrr4 := [5]int{1: 10, 2: 40}  // Initialized only specific elements

	fmt.Println(arrr1)
	fmt.Println(arrr2)
	fmt.Println(arrr3)
	fmt.Println(arrr4)

	//SLICE
	// unlike arrays, the length of a slice can grow and shrink.
	myslice := []int{}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice)) // return the capacity of the slice (the number of elements the slice can grow or shrink to)
	fmt.Println(myslice)
	// create slice from array
	arrr5 := [6]int{10, 11, 12, 13, 14, 15}
	myslice1 := arrr5[2:4]

	fmt.Printf("myslice = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1)) // cap in this case is 4
	// create slice with make function
	myslice2 := make([]int, 5, 10) // make([]type, length, capacity)
	fmt.Printf("myslice2 = %v\n", myslice2)
	//append elements to a slice
	myslice2 = append(myslice2, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice2)

	//STRUCT
	//A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
	var pers1 Person
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	printPerson(pers1)

	//MAP
	//Maps are used to store data values in key:value pairs.
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Printf("b\t%v\n", b)
	a := make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	fmt.Printf("a\t%v\n", a)
	//Iterate Over Map
	//Maps are unordered data structure.
	c := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	var d []string // defining the order
	d = append(d, "one", "two", "three", "four")

	for _, element := range d { // loop with the defined order
		fmt.Printf("%v : %v, ", element, c[element])
	}
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
