package main

import "log"

func main() {
	// var isTrue bool

	// isTrue = true

	// if isTrue {
	// 	log.Println(isTrue)
	// } else {
	// 	log.Println(isTrue)
	// }

	//	//	//	//	//

	// myNum := 100
	// isTrue := false

	// if myNum > 99 && !isTrue {
	// 	log.Println(1)
	// } else if myNum < 100 && isTrue {
	// 	log.Println(2)
	// } else if myNum ==101 || isTrue {
	// 	log.Println(3)
	// } else {
	// 	log.Println(4)
	// }

	//	//	//	//	//

	myVar := "a"

	switch myVar {
	case "a":
		log.Println(myVar)	// no need to write break statement explicitly
	case "b":
		log.Println(myVar)
	default:
		log.Println(myVar)
	}

}
