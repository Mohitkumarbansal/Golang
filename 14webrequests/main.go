package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.amazon.com/"

func main() {
	fmt.Println("*** Learning webrequests ****")

	response, err := http.Get(url) // http ka get method
	if err != nil {
		panic(err)
	}
	fmt.Printf("*** Type of response => %T\n", response)
	defer response.Body.Close() // callers responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body) // As response body will be in type other than string
	if err != nil {
		panic(err)
	}
	content := string(databyte) // convert it into string
	fmt.Println("*** content => ", content)

}
