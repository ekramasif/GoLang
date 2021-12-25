package main

import "fmt"

func main() {
	fmt.Println("Odd Number[1 to 10]")
	for i := 1; i <= 10; i = i + 2 {
		fmt.Println(i)
	}
	fmt.Println("Even Number[1 to 10]")
	for i := 2; i <= 10; i = i + 2 {
		fmt.Println(i)
	}
}
