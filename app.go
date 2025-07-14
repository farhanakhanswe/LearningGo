package main

import "fmt"

func main(){
	//	Print is a function provided by built-in "fmt" package

	fmt.Print("Hello World\n")
	fmt.Print(`Hello World with Back Ticks!`)

	/* fmt.Print('Hello World') -> This will cause error as strings need to be either wrapped with double quotes
	or back ticks
	*/
}

