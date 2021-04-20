package main

import "log"

func main() {
	var whatToSay string
	var saySomethingElse string
	var i int

	whatToSay = saySomething("Hello World")

	log.Println(whatToSay)

	log.Println(i)

	saySomethingElse, _ = saySomething2("Hello")

	log.Println(saySomethingElse)

	log.Println(saySomething2("Hello"))
}

func saySomething(s string) string {
	return s
}

func saySomething2(s string) (string, string) {	//	return two strings
	return s, "World"
}