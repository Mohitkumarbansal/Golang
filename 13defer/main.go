package main

import "fmt"

func main() {
	/*
		syntax => defer <some_expression>

		the expressions called by defer will be execute in LIFO mode at the end of sorounding function returns
	*/
	defer fmt.Println(" World")
	fmt.Println("Hello")
	defer myDefer()

	/*
		as I mentioned if multiple defer statements then LIFO execution

		result => Hello

		LIFO => World myDefer()

		So myDefer() will get called
			in myDefer() also we have defer so LIFO order will be 0,1,2,3,4

		result => Hello , 4,3,2,1,0, World
	*/
}

func myDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
