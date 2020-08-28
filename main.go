package main

import "log"

func main() {
	log.Println("MAIN APP")
	calculate(9, 1)
}

func calculate(x, i int) {
	log.Println(x + i)
}
