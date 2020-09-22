package main

import (
	"encoding/json"
	"fmt"
	"golf/gedis"
	"golf/lib"
	"log"
	"time"
)

func main() {
	cache := gedis.NewSimpleCache(gedis.NewStringOperation(), time.Second*15)

	id := 1
	cache.Getter = func() string {
		log.Println("from db")
		model := lib.NewNew()
		lib.Gorm.Where("id=?", id).Find(model)
		b, _ := json.Marshal(model)
		return string(b)
	}

	fmt.Println(cache.GetCache("new:1"))
}
