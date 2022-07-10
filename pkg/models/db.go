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
	User    string
	Command string
	Result  string
}
