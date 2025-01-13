package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&productid=jnf"

func main() {
	fmt.Println("*** Welcome to Handling URLs in Golang")
	fmt.Println("*** myurl => ", myurl)

	urlParser, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("*** Schema => ", urlParser.Scheme)     // output is "https"
	fmt.Println("*** Host => ", urlParser.Host)         // output is "lco.dev:3000"
	fmt.Println("*** Path => ", urlParser.Path)         // output is "/learn"
	fmt.Println("*** RawQuery => ", urlParser.RawQuery) // output is"coursename=reactjs&productid=jnf"
	fmt.Println("*** Port => ", urlParser.Port())       // output is 3000
	query_in_map := urlParser.Query()
	fmt.Printf("*** Type of query_in_map => %T\n", query_in_map)
	for key, value := range query_in_map {
		fmt.Println("*** key => %v , value => %v", key, value)
	}
}
