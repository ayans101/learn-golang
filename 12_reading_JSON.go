package main

import (
	"encoding/json"
	"log"
)

type Person struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
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
}
