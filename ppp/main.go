package main

import "fmt"

func say(word string) {
	for i := 0; i < 3; i++ {
		fmt.Println(word)
	}
}

func main() {
	go say("world")
	say("hello")
}
