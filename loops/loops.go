package main

import "fmt"

func main() {

	for_loop()
	var sum = incremental_sum()
	fmt.Println("The incremental sum is :", sum)
}

func for_loop() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func incremental_sum() int {

	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	return sum
}
