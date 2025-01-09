package main

import "fmt"

// execution of go starts from main function
func main() {
	addition, return_string := proAdder(1, 2, 3, 4, 5)
	fmt.Println("*** addition => ", addition)
	fmt.Println("*** return_string => ", return_string)
}

// syntax for function
// func function_name(variable <variable_type>) (can be multiple return types){
// }

func proAdder(values ...int) (int, string) { // for arguments we are expecting multiple values and return type we are returning two values

	result := 0

	for index := range values {
		result += values[index]
	}

	return result, "moit"
}
