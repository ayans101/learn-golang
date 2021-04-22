package main

import (
	"github.com/ayans101/learn-golang/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}