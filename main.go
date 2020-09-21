package main

import (
	"fmt"
	"golf/gedis"
)

func main() {

	fmt.Println(gedis.NewStringOperation().Get("abc").UnwrapOr("xx"))
}
