package main

import (
	"fmt"
	"golf/gedis"
	"time"
)

func main() {
	/*result := gedis.
		NewStringOperation().
		MGet("name", "age", "abc").
		Iterate()

	for result.HasNext() {
		fmt.Println(result.Next())
	}*/
	fmt.Println(gedis.
		NewStringOperation().
		Set("name", "jj", gedis.WithExpire(15*time.Second), gedis.WithNx()))
}
