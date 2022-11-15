package model

import (
	"github.com/Random7-JF/go-rcon-svelte/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(App *config.App) {
	var err error
	dsn := "host=" + App.DbSettings.Host + " user=" + App.DbSettings.User + " password=" + App.DbSettings.Password + " dbname=" + App.DbSettings.DbName + " port=" + App.DbSettings.Port + " sslmode=disable TimeZone=America/Winnipeg"
	App.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	App.Db.AutoMigrate(&Users{})
	App.Db.AutoMigrate(&Players{})
	App.Db.AutoMigrate(&CommandLog{})
}
