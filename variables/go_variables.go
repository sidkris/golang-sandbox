package main

import "fmt"

func main() {

	var myName = "Sid"
	fmt.Println("Hi, my name is", myName)

	var userName string = "_sid_"
	fmt.Println("username :", userName)

	var sum int = 5 + 6
	fmt.Println("The sum is :", sum)

	var (
		firstInitial = 'S'
		lastInitial  = 'K'
	)

	initials := fmt.Sprintf("First Initial : %c | Last Initial : %c", firstInitial, lastInitial)
	fmt.Println(initials)
}
