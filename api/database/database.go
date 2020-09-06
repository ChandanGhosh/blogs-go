package database

import (
	"fmt"

	"github.com/chandanghosh/blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var BlogDB *gorm.DB

// func init() {
// 	var err error
// 	BlogDB, err = gorm.Open(mysql.Open(config.DBURI), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error openning DB connection: %v", err)
// 	}

// }

// Connect ..
func Connect() error {
	var err error
	BlogDB, err = gorm.Open(mysql.Open(config.DBURI), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error opening database; %v", err)
		return err
	}
	return nil
}
