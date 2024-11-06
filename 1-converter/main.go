package main

import "fmt"

func main() {
	const USDtoEUR = 0.93
	const USDtoRUB = 98.06
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)

}
