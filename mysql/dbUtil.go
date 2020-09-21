package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func DbObj() *gorm.DB {
	dataUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := gorm.Open("mysql", dataUrl)
	if err != nil {
		fmt.Println(err)
		defer db.Close()
		return db
	} else {
		fmt.Println("connection succedssed")
		db.SingularTable(true)
	}
	// defer mysql.Close()
	return db
}
