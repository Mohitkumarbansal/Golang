package main

import "fmt"

func main() {
	u := User{"mohit", "bansal"}
	u.GetFirstName()
	u.SetFirstName()
	u.GetFirstName()
}

type User struct {
	FirstName string
	LastName  string
}

func (u User) GetFirstName() {
	fmt.Printf("*** FirstName => %v\n", u.FirstName)
}

func (u User) SetFirstName() {
	u.FirstName = "rohit" // here you can see we have changes the value of FirstName but it will not be reflect in main u because it is pass by value
	fmt.Printf("*** New FirstName =>%v\n", u.FirstName)
}
