package main

import (
	"errors"
	"fmt"
)

const (
	first  = iota + 6
	second = 2 << iota
	third
)

const (
	fourth = iota
)

type address struct {
	LineOne string
	LineTwo string
}

type user struct {
	ID        int
	FirstName string
	LastName  string
	Address   address
}

func main() {
	// variables()
	// pointers()
	// constants()
	// arrays()
	// slices()
	// maps()
	// structs()
	// functions()
	// loops()
	conditionals()
}

func variables() {
	var test int
	test = 42
	fmt.Println(test)

	var floating float32 = 3.14
	fmt.Println(floating)

	firstName := "Person"
	fmt.Println(firstName)

	boo := true
	fmt.Println(boo)

	comp := complex(3, 4)
	fmt.Println(comp)

	re, im := real(comp), imag(comp)
	fmt.Println(re, im)
}

func pointers() {
	var firstName *string = new(string)
	*firstName = "Person"
	fmt.Println(*firstName)

	testName := "Testing"
	pointing := &testName
	fmt.Println(pointing, *pointing)

	testName = "Second"
	fmt.Println(pointing, *pointing)
}

func constants() {
	const pi = 3.1415
	fmt.Println(pi)

	const test = 3
	fmt.Println(test + 3)

	fmt.Println(test + 1.5)

	const integer int = 3
	fmt.Println(float32(integer) + 1.5)

	fmt.Println(first, second, third, fourth)
}

func arrays() {
	var firstArray [3]int
	firstArray[0] = 1
	firstArray[1] = 2
	firstArray[2] = 3
	fmt.Println(firstArray)

	secondArray := [3]int{1, 2, 3}
	fmt.Println(secondArray)
}

func slices() {
	array := [3]int{1, 2, 3}
	slice := array[:]
	fmt.Println(array, slice)

	array[0] = 42
	slice[1] = 19
	fmt.Println(array, slice)

	more := []int{1, 2, 3}
	fmt.Println(more)

	more = append(more, 4, 5)
	fmt.Println(more)

	skipStart := more[1:]
	skipEnd := more[:4]
	middle := more[2:3]
	fmt.Println(skipStart, skipEnd, middle)
}

func maps() {
	myMap := map[string]int{"one": 1, "two": 2}
	fmt.Println(myMap, myMap["one"])

	myMap["one"] = 42
	fmt.Println(myMap["one"])

	delete(myMap, "one")
	fmt.Println(myMap)
}

func structs() {
	var home address
	home.LineOne = "LineOne"
	home.LineTwo = "LineTwo"

	var person user
	person.ID = 1
	person.FirstName = "FirstName"
	person.LastName = "LastName"
	person.Address = home
	fmt.Println(person)

	second := user{ID: 2, FirstName: "Another", LastName: "Somebody", Address: address{LineOne: "More", LineTwo: "Place"}}
	fmt.Println(second)

	tidy := user{
		ID:        3,
		FirstName: "ABC",
		LastName:  "DEF",
		Address: address{
			LineOne: "Test",
			LineTwo: "Test",
		},
	}
	fmt.Println(tidy)
}

func functions() {
	fmt.Println(add(2, 3))
	fmt.Println(explode())
	one, two := test()
	fmt.Println(one, two)
	_, more := test()
	fmt.Println(more)
}

func add(one, two int) int {
	return one + two
}

func explode() error {
	return errors.New("Something blew up")
}

func test() (string, string) {
	return "One", "Two"
}

func loops() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	var i int
	for {
		if i == 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	slice := []int{1, 2, 3}
	for i, v := range slice {
		fmt.Println(i, v)
	}

	aMap := map[string]int{"One": 1, "Two": 2}
	for k, v := range aMap {
		fmt.Println(k, v)
	}
	for k := range aMap {
		fmt.Println(k)
	}
	for _, v := range aMap {
		fmt.Println(v)
	}
}

func conditionals() {
	// panic("Something happened")

	u1 := user{ID: 1}
	u2 := user{ID: 2}
	if u1.ID == u2.ID {
		fmt.Println("Same user")
	}

	if u1 == u2 {
		fmt.Println("Same user")
	} else if u1.FirstName == u2.FirstName {
		fmt.Println("Similar")
	} else {
		fmt.Println("Else different")
	}
}
