package main

import "log"

func main() {
	num := calculate(9, 1)
	log.Println(num)

	num2 := sub(10, 9)
	log.Println(num2)

	x := calculate(num, num2)
	log.Println(x)

	callAPI()

	// login with email and username

	// feature register
}

func calculate(x, i int) int {
	return x + i
}

func sub(x, i int) int {
	return x - i
}

func callAPI() {
	// read user

	// create user

	// update user

	// delete user

	// update status
}
