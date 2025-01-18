package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username  string `"gorm:'type:varchar(20);not null'"`
	Telephone string `"gorm:'type:varchar(11);not null;unique'"`
	Password  string `"gorm:'type:size:255;not null'"`
}
