package main

import "fmt"

func main() {
	fmt.Println("*** Welcome to go maps ***")

	go_maps := make(map[string]string)
	go_maps["JS"] = "Javascript" // for assignment
	go_maps["RB"] = "Ruby"
	go_maps["PY"] = "Python"

	fmt.Printf("*** Type of go_maps => %T\n", go_maps)
	fmt.Println("*** print go_maps => ", go_maps)
	fmt.Println("*** print go_maps[JS] => ", go_maps["JS"])

	// to delete particular key from map
	delete(go_maps, "JS")
	fmt.Println("*** print go_maps => ", go_maps)

	for key, value := range go_maps {
		fmt.Printf("key => %v, value => %v\n", key, value)
	}

}
