package main

import "log"

func main() {
	// var myString string
	// var myInt int

	// myString = "Hi"
	// myInt = 11

	// mySecondString := "another string"
	// log.Println(myString, mySecondString, myInt)

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])

	myMap["dog"] = "Fido"

	log.Println(myMap["dog"])

}
