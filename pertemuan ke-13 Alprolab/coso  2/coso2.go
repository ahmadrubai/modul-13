package main

import "fmt"

func main() {
	var word string
	var repetitions int
	fmt.Scan(&word, &repetitions)
	counter := 0
	for done := false; !done; {
		fmt.Println(word)
		counter++
		done = (counter >= repetitions)
		// for i:= 0; i <repetitions; i++
		// fmt.scan (word)
	}
}