package seeder

import (
	"log"

	"github.com/chandanghosh/blog/api/database"
	"github.com/chandanghosh/blog/models"
	"github.com/chandanghosh/blog/util/console"
)

// SeedDatabase ..
func SeedDatabase() {
	db := database.ConnectMySQL()
	err := db.Debug().AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(user)
	}
}
