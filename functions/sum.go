package main

import "fmt"

func main() {

	var sum_ = sum(5, 6)
	fmt.Println("SUM :", sum_)

	var doubled_sum = double(sum_)
	fmt.Println("DOUBLED SUM :", doubled_sum)

}

func sum(num1 int, num2 int) int {
	var sum_ = num1 + num2
	return sum_
}

func double(x int) int {

	return x * 2

}
