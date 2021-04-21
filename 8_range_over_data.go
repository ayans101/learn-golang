package main

import "log"

type User struct {
	FirstName string
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	log.Println(i)
	// }

	mySlice := []string{"apple", "banana", "fish", "horse"}

	for _, x := range mySlice {
		log.Println(x)
	}

	//	//	//	//	//

	myMap := make(map[string]string)
	myMap["apple"] = "apple"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	for i, x := range myMap { //	for maps the order won't be same every time
		log.Println(i, x)
	}

	//	//	//	//	//

	var mySlice2 []User
	u1 := User{
		FirstName: "Ayan",
	}
	u2 := User{
		FirstName: "Arnav",
	}
	u3 := User{
		FirstName: "Aryan",
	}

	mySlice2 = append(mySlice2, u1)
	mySlice2 = append(mySlice2, u2)
	mySlice2 = append(mySlice2, u3)

	for i, x := range mySlice2 {
		log.Println(i, x.FirstName)
	}
}
