package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("HelloWorld.txt")
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content))
}
