package models

import "gorm.io/gorm"

//OnlinePlayers is storage for current players to reduce the rcon calls
type OnlinePlayers struct {
	gorm.Model
	CurrentCount int
	Players      []string
}

//CommandLog is for logging every command sent from the webapp
type CommandLog struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	User    string
	Command string
	Result  string
}

//Users is for website users
type Users struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
}
