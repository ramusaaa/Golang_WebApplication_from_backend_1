package models

import "gorm.io/gorm"

type Product struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (p Product) Count(db *gorm.DB) int64 {
	//TODO implement me
	panic("implement me")
}

func (p Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	//TODO implement me
	panic("implement me")
}
