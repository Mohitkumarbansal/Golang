package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// bufio is a package which takes stdin as a input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the user input")

	// comma ok | error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("input => " + input)

}
