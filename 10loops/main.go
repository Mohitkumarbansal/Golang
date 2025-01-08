package main

import "fmt"

func main() {
	fmt.Println("*** Welcome to loops in go ***")

	slice_for_names := []string{"mohit", "rohit", "keshav"}

	// 1st case
	// for i := 0; i < len(slice_for_names); i++ { // pre increament is not present as we have it in C++
	// 	fmt.Printf("Printing through 1st case => Value at %v => %v\n", i, slice_for_names[i])
	// }
	// fmt.Println()

	// 2nd case
	// for i := range slice_for_names { // here like other languages return directly the element here in golang its a index
	// 	fmt.Printf("Printing through 2nd case => Value at %v => %v\n", i, slice_for_names[i])
	// }
	// fmt.Println()

	// 3rd case
	// for index, data := range slice_for_names { // here index refers to index of array and data refers to data present at index
	// 	fmt.Printf("Printing through 3rd case => Value at %v => %v\n", index, data)
	// }

	index := 0

	for index < len(slice_for_names) { // like a while loop
		fmt.Println(index)
		index++
	}

	for ; index < len(slice_for_names); index++ {
		fmt.Println(index)
	}

	// fmt.Println(index)

}
