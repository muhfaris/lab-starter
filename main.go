package main

import "log"

func main() {
	log.Println("MAIN APP")
	calculate(9, 1)
	sub(10, 9)
}

func calculate(x, i int) {
	log.Println(x + i)
}

func sub(x, i int) {
	log.Println(x - i)
}
