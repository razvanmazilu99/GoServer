package db

import "goserver/entity"

func InitData() {
	person := entity.Person{
		ID:   "1",
		Name: "John Doe",
		Age:  10,
	}
	GetDB().Save(&person)
}
