package lib

import (
	"encoding/json"
	"fmt"
	"golf/gedis"
)

func NewsDbGetter(id string) gedis.DBGetFunc {
	return func() string {
		fmt.Println("from db")

		model := NewNew()
		Gorm.Where("id=?", id).Find(model)
		b, _ := json.Marshal(model)
		return string(b)
	}

}
