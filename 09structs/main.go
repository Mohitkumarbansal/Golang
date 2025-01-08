package main

import "fmt"

func main() {
	fmt.Println("*** Welcome to structs ***")

	mohit_struct := User{"mohit", "bansal"}
	fmt.Printf("*** mohit_struct =>%+v\n", mohit_struct)
	fmt.Printf("*** FirstName => %v\n", mohit_struct.FirstName)
}

type User struct {
	FirstName string
	LastName  string
}
