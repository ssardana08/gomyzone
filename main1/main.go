package main

import (
	"fmt"
	"gomyzone/main1/calc"
	"gomyzone/main1/caller"
)

// This main was written to check how init works from other packages
func main() {
	fmt.Println("Hello World")
	fmt.Println(calc.X)
	caller.Run()
}
