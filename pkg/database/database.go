package database

import (
	"fmt"

	"github.com/Random-7/GoRcon/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Session struct {
	Db       *gorm.DB
	IP       string
	User     string
	Password string
	DbName   string
}

// SetupInitial forms the initial dsn from the populated values in session struct and connects to the db
func (d *Session) Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.User, d.Password, d.IP, d.DbName)
	var err error

	d.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error with database:", err)
	}

	d.migrate()
	d.createAdmin()
}

// TODO specify a admin account in config.json
// createAdmin Creates the primary admin account if not already made.
func (d *Session) createAdmin() {
	_, err := d.GetUser("admin")
	if err != nil {
		_, err := d.CreateUser(models.User{
			Username: "admin",
			Password: "password",
			Email:    "admin@email.com"})
		if err != nil {
			fmt.Println("Create Admin error:", err)
		}
		fmt.Println("Admin account created")
		return
	}
}

// Migrate uses gorms automigrate to check/create the need tables based on models
func (d *Session) migrate() {
	//automigrate databases
	d.Db.AutoMigrate(&models.User{})
	d.Db.AutoMigrate(&models.CommandLog{})
	d.Db.AutoMigrate(&models.OnlinePlayers{})
}
