package seeder

import (
	"log"

	"github.com/chandanghosh/blog/api/database"
	"github.com/chandanghosh/blog/models"
	"github.com/chandanghosh/blog/util/console"
)

// SeedDatabase ..
func SeedDatabase() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = database.BlogDB.Debug().AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {

		err = database.BlogDB.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(user)
	}
}
