package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) returnFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println(myVar.returnFirstName());
	log.Println(myVar2.returnFirstName());
}
