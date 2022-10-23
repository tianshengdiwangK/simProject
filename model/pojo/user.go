package pojo

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Passwd   string
	Gender   int64 // 1 -> girl, 2 -> boy
	Age      int64
	Interest string
}
