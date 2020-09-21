package main

import (
	"context"
	"golf/gedis"
	"log"
)

func main() {
	ctx:=context.Background()
	r:=gedis.Redis().Get(ctx,"name")
	if v,err:=r.Result();err!=nil {
		log.Fatal(err)
	}else {

		log.Println(v)
	}
}
