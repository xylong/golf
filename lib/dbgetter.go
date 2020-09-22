package lib

import (
	"fmt"
	"golf/gedis"
)

func NewsDbGetter(id string) gedis.DBGetFunc {
	return func() interface{} {
		fmt.Println("from db")

		news := NewNew()
		Gorm.Where("id=?", id).Find(news)
		return news
	}

}
