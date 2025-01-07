package main

import "fmt"

func main() {
	// arrays are not something which go developers most usely, they have slices which they do use a lot
	// 1st scenerio
	var array [4]int // declaring a array of size 4, all elements will be initialized to 0 as int a type, for string it will be empty string
	array[0] = 1     // assigning the 1 to 0th element and all other elements will
	fmt.Println("array printing => ", array)
	fmt.Println("*** length of array => ", len(array))
	fmt.Printf("*** Type of array => %T\n", array)
	fmt.Println()

	// 2nd scenerio

	var array2 = [4]string{"mohit", "rohit"}
	fmt.Println("array2 printing => ", array2)
	fmt.Println("*** length of array2 => ", len(array2))
	fmt.Printf("*** Type of array2 => %T\n", array2)

}
