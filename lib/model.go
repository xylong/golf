package lib

import "time"

type New struct {
	ID      int       `gorm:"column:id" json:"id"`
	Title   string    `gorm:"column:title" json:"title"`
	Content string    `gorm:"column:content" json:"content"`
	Views   int       `gorm:"column:views" json:"views"`
	AddTime time.Time `gorm:"column:add_time" json:"add_time"`
}

func NewNew() *New {
	return &New{}
}
