package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Bruce",
			"last_name": "Wayne"
		},
		{
			"first_name": "Clark",
			"last_name": "Kent"
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	//	//	//	//	//

	//	write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Tony"
	m1.LastName = "Stark"

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Steve"
	m2.LastName = "Rogers"

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "	")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}
