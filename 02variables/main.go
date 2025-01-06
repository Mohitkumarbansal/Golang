package main

import "fmt"

const Firstname string = "mohit" // here constant is starting by capital letter means its a public

func main() {

	var username string = "mohit"
	fmt.Println(username)
	fmt.Printf("Variable type of username =>  %T \n", username)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("Variable type of isBool => %T \n", isBool)

	var unitvariable uint8 = 255 // unit8 simply means it can only store values till 2^8 - 1
	fmt.Println(unitvariable)
	fmt.Printf("Variable type of unitvariable => %T \n", unitvariable)

	var float32variable float32 = 255.5459345034 // here float32 simply related to precision not like 32 bits or something
	fmt.Println(float32variable)
	fmt.Printf("Variable type of float32variable => %T \n", float32variable)

	var float64variable float64 = 255.5459345034
	fmt.Println(float64variable)
	fmt.Printf("Variable type of float64variable => %T \n", float64variable)

	// default values
	var defaultint int
	fmt.Println(defaultint)
	fmt.Printf("Default value of defaultint int => %d\n", defaultint)

	var defaultstring string
	fmt.Println(defaultstring)
	fmt.Printf("Default value of defaultstring string => %s\n", defaultstring)

	var defaultfloat float32
	fmt.Println(defaultfloat)
	fmt.Printf("Default value of defaultfloat float32 => %f\n", defaultfloat)

	// impicit type identification by go
	var implicittype = 3
	fmt.Printf("Variable type of implicittype => %T\n", implicittype)

	// variable decleration without var, ":=" is called as volorous operator it can be use in function but not outside
	withoutVarInt := 3000
	fmt.Println(withoutVarInt)
	fmt.Printf("Variable type of withoutVarInt => %T \n", withoutVarInt)

	// constant
	fmt.Println(Firstname)
	fmt.Printf("Variable type of Firstname => %T \n", Firstname)

}
