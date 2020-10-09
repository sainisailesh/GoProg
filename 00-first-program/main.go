package main

import "fmt"

func main() {
	fmt.Println("Hello World!!")

	hi()
	even()
	fmt.Println("Bye!!")
}

func hi() {
	fmt.Println("This is Hi function")
}

func even() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
