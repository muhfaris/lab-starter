package main

import "log"

func main() {

	num := calculate(9, 1)
	log.Println(num)

	num2 := sub(10, 9)
	log.Println(num2)
}

func calculate(x, i int) int {
	return x + i
}

func sub(x, i int) int {
	return x - i
}
