package main

import "fmt"

func main() {

	fmt.Println("** Welcome to session of slices which internally use arrays")

	var fruitlist = []string{} // giving elements is optional

	fmt.Printf("*** Type of fruitlist=> %T\n", fruitlist)
	fmt.Println("*** length of fruitlist => ", len(fruitlist))

	fruitlist = append(fruitlist, "mohit", "rohit", "keshav", "vishnu", "babita")
	fmt.Println("*** length of fruitlist => ", len(fruitlist))

	// in following first range is inclusice and second range is not inclusive and if we do not give range then default 0 and length will be considered
	fruitlist = append(fruitlist[1:3])
	fmt.Println("*** fruitlist => ", fruitlist)

	// different syntax
	var slice1 []string
	fmt.Println("*** slice1 => ", slice1)

	//declare slices using make()
	highscores := make([]int, 4) // here we can see only 4 elements can be stores
	highscores[0] = 0
	highscores[1] = 1
	highscores[2] = 2
	highscores[3] = 3
	// highscores[4] = 4 // It will give out of bound error as make has allocated a highscores with 4 size
	// but for above we can use append method which will reallocate the memory
	highscores = append(highscores, 5, 6)
	fmt.Println("*** highscores => ", highscores)
	fmt.Printf("*** Type of highscores => %T \n", highscores)

	// How to remove specific index value from the slice
	// for this we have to use append method
	slice_with_index := []int{1, 2, 3, 4, 5}
	index_to_be_removed := 2
	slice_with_index = append(slice_with_index[:index_to_be_removed], slice_with_index[index_to_be_removed+1:]...)
	fmt.Println("*** slicd with index => ", slice_with_index)

}
