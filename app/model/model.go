package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	DisplayName string
	UserName    string
	Password    string
	Email       string
}

type CommandLog struct {
	gorm.Model
	CommandType string
	Command     string
	SentBy      string
	Response    string
}
