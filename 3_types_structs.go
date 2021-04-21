package main

import (
	"log"
	"time"
)

// var s = "seven"
// var S = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// log.Println(s)

	user := User{
		FirstName: "Ayan",
		LastName:  "Sarkar",
	}

	log.Println(user)
}
