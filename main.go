package main

import "fmt"

//struct defines the data types we need
type helloWorldIterator struct {
	helloWorld string
	iterations int
}

//referencing our struct as the argument implicitly applies iterate to our type.
func (hwi helloWorldIterator) iterate() {
	for i := 0; i < hwi.iterations; i++ {
		fmt.Println(hwi.helloWorld)
	}
}

//Main loop that always executes on run or execution.
func main() {
	//create our struct and store it
	helloWorldIteration := helloWorldIterator{ helloWorld: "whats good?", iterations: 2}

	//call our implicitly referenced function on our type.
	helloWorldIteration.iterate()
}