package main

import "log"

func main() {
	log.Println("MAIN APP")
	calculate()
}

func calculate(x, i int) {
	log.Println(x + i)
}
