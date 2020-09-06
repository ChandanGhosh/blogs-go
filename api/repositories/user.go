package repositories

import (
	"log"

	"github.com/chandanghosh/blog/api/database"
	"github.com/chandanghosh/blog/models"
	"github.com/chandanghosh/blog/util/console"
)

// UserRepository ..
type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindByID(uid uint32) (models.User, error)
	Update(uint32, models.User) (models.User, error)
	Delete(uid uint32) error
}

// UserRepo ..
type UserRepo struct {
}

// NewUserRepo ..
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// Save ..
func (u *UserRepo) Save(user models.User) (models.User, error) {

	err := database.BlogDB.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}
	console.Pretty(user)
	return user, nil
}

// FindAll ..
func (u *UserRepo) FindAll() ([]models.User, error) {
	users := []models.User{}
	if err := database.BlogDB.Debug().Model(&models.User{}).Limit(100).Find(&users).Error; err != nil {
		return users, err
	}

	console.Pretty(users)
	return users, nil
}

// FindByID ..
func (u *UserRepo) FindByID(uid uint32) (models.User, error) {
	user := models.User{}
	if err := database.BlogDB.Debug().Model(&models.User{}).First(&user, uid).Error; err != nil {
		return user, err
	}

	console.Pretty(user)
	return user, nil
}

// Update ..
func (u *UserRepo) Update(uid uint32, user models.User) (models.User, error) {

	if err := database.BlogDB.Debug().Model(&models.User{}).Where("id=?", uid).Updates(&user).Error; err != nil {
		return models.User{}, err
	}
	console.Pretty(user)
	return user, nil
}

// Delete ..
func (u *UserRepo) Delete(uid uint32) error {
	usr := models.User{}
	if err := database.BlogDB.Debug().Model(&models.User{}).Delete(&usr, uid).Error; err != nil {
		return err
	}
	console.Pretty(usr)
	return nil
}
