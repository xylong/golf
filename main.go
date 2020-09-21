package main

import (
	"fmt"
	"golf/gedis"
)

func main() {
	result := gedis.
		NewStringOperation().
		MGet("name", "age", "abc").
		Iterate()

	for result.HasNext() {
		fmt.Println(result.Next())
	}
}
