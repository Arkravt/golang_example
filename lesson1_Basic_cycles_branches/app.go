package main

import (
	"fmt"
)

func main() {
	/*
		var a1 int8   // -128 -> 127, 1 byte
		var a2 int16  // -32768 -> 32767, 2 byte
		var a3 int32  // -2bil -> 2bill, 4 byte
		var a4 int64  // -9pent -> 9pent, 8 byte
		var a5 uint8  // 0 -> 255, 1 byte
		var a6 uint16 // 0 -> 65535, 2 byte
		var a7 uint32 // 0 -> 4 bill, 4 byte
		var a8 uint64 // 0 -> 18pent, 8 byte

		var a9 byte        // synonym int8
		var a10 rune       // synonym int32
		var a11 int        // int32 or int64
		var a12 uint       // uint32 or uint64
		var a13 float32    // 4 byte
		var a14 float64    // 8 byte
		var a15 complex64  //
		var a16 complex128 //

		var a17 bool    // true or false
		var name string //
	*/

	deferExample()         // defer
	assigningVariables()   // Assigning variables
	callingFunctions()     // Calling functions
	cycles()               // cycles
	res, b := branches(-2) // branches
	fmt.Println(res, b)

}

func assigningVariables() {
	var name = "Artem"
	var (
		age  = 32
		city = "Samara"
	)
	var country, sex = "Russia", "male"
	fmt.Println(name, age, city, country, sex)

	res := fmt.Sprintf("My name is %s , age is %d", name, age)
	fmt.Println(res)
}

func callingFunctions() {
	var personName, lastName, age = getPersonInfo()
	res1 := fmt.Sprintf("Person info is: %s, %s, %d", lastName, personName, age)

	personName, lastName, _ = getPersonInfo()
	res2 := fmt.Sprintf("Person info is: %s, %s", lastName, personName)

	personName, _, age = getPersonInfo()
	res3 := fmt.Sprintf("Person info is: %s, %d", lastName, age)

	someName := getName(personName, age)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(someName)
	fmt.Println(getName(personName, 30))
}

func cycles() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)

	// while
	for sum > 0 {
		sum--
		fmt.Println(sum)
	}

}

func branches(value int) (bool, int) {
	//if value >= 18 {
	//	return true, value
	//}
	//return false, 0

	switch {
	case value >= 18:
		return true, value
	case value < 18 && value >= 0:
		return false, 0
	default:
		return false, -1
	}
}

func deferExample() {
	/*
		defer - Всегда выполняется в самом конце программы, когда метод main заканчивает работу.
		Работает по принципу LIFO
	*/
	defer fmt.Println("last1")
	defer fmt.Println("last2")
	defer fmt.Println("last3")
	defer fmt.Println("last4")
	fmt.Println("first")
}

func getPersonInfo() (string, string, int) {
	name := "Artem"
	lastName := "Kravtsev"
	age := 32

	return name, lastName, age
}

func getName(name string, age int) string {
	return fmt.Sprintf("name: %s, age:  %d", name, age)
}
