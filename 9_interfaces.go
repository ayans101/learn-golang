package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Horse struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	horse := Horse{
		Name:  "Samson",
		Breed: "Morgan",
	}

	PrintInfo(horse)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(gorilla)

}

func (h Horse) Says() string {
	return "neigh"
}

func (h Horse) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "grunt"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
