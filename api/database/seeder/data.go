package seeder

import (
	"github.com/chandanghosh/blog/models"
)

var (
	users = []models.User{
		{Nickname: "batman", Email: "batman@email.com", Password: "123456789"},
		{Nickname: "superman", Email: "superman@email.com", Password: "123456789"},
		{Nickname: "spiderman", Email: "spiderman@email.com", Password: "123456789"},
		{Nickname: "ironman", Email: "ironman@email.com", Password: "123456789"},
	}
)
