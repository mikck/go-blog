package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName  string
	Email     string `gorm:"type:varchar(100);unique_index"`
	Phone     string
	Password  string `gorm:"-"`
	LoginTime time.Time
}
