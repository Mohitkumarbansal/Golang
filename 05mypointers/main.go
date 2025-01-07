package main

import "fmt"

func main() {

	var myNumber int = 2
	// almost basic pointers in go are same as C++
	// ptr is a pointer to myNumber address
	var ptr *int = &myNumber // could be written using volorous operator => ptr:=&myNumber
	fmt.Println("**** addres tored in ptr => ", ptr)
	fmt.Println("**** value stored un myNumber through pointer => ", *ptr)

	*ptr = *ptr + 1
	fmt.Println("**** value stored un myNumber through pointer => ", *ptr)

}
