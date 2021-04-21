package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice) //	[1 2 3]



	numbers := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	log.Println(numbers)

	log.Println(numbers[0:2])	//	[11, 12]
	log.Println(numbers[9:10])	//	[20]



	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)


}
