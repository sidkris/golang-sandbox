package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDie(r *rand.Rand) int {
	return r.Intn(6) + 1
}

func main() {
	// Create a new random source with the current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Roll two dice
	die1 := rollDie(r)
	die2 := rollDie(r)

	total := die1 + die2

	fmt.Printf("You rolled a %d and a %d, for a total of %d\n", die1, die2, total)

	switch total {
	case 2:
		fmt.Println("Snake eyes!")
	case 7:
		fmt.Println("Lucky seven!")
	case 6:
		fmt.Println("Hit for a six!")
	case 5:
		fmt.Println("High five!")
	case 12:
		fmt.Println("That's a maximum!")
	case 10:
		fmt.Println("Perfect ten!")
	default:
		if total%2 == 0 {
			fmt.Println("Even!")
		} else {
			fmt.Println("Odd!")
		}
	}
}
