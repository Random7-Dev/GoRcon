package database

import (
	"fmt"

	"github.com/Random-7/GoRcon/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Session struct {
	Db *gorm.DB
}

//TODO move the connection params into the session struct.
func (d *Session) SetupInitial(ip, user, password, database string) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, database)
	fmt.Println("DSN:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d.Db = db
	if err != nil {
		fmt.Println("Error with database:", err)
	}

	d.Db.AutoMigrate(&models.Users{})

	var testuser models.Users
	d.Db.First(&testuser, 1)
	fmt.Println(testuser)
}
