package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Passwd string
	Auth uint32
}
