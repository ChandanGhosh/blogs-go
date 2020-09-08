package database

import (
	"log"

	"github.com/chandanghosh/blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectMySQL ..
func ConnectMySQL() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DBURI), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return db
}
