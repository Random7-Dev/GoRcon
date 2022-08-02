package database

import (
	"github.com/Random-7/GoRcon/pkg/models"
)

//GetUser uses gorm to query the db for a user by user name and returns a models.user and error
func (d *Session) GetUser(username string) (models.User, error) {
	var user models.User
	result := d.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

//CreateUser uses gorm to insert a new user into the database and returns a rows effected count and error
func (d *Session) CreateUser(user models.User) (int, error) {
	result := d.Db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
