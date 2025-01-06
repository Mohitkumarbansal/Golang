package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please give rating to out shop : ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {

		newRating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
		fmt.Println(newRating + 1)

	}

}
