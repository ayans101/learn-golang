package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)
	defer close(fooVal)

	go foo(fooVal, 5)
	v1 := <-fooVal
	go foo(fooVal, 3)
	v2 := <-fooVal
	go foo(fooVal, 2)
	v3 := <-fooVal

	// v1 := <-fooVal
	// v2 := <-fooVal
	// v3 := <-fooVal

	// v1, v2, v3 := <-fooVal, <-fooVal, <-fooVal

	fmt.Println(v1, v2, v3)	//	25 15 10
}
