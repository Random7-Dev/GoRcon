package model

import (
	"gorm.io/gorm"
)

type Players struct {
	gorm.Model
	Name string
	UUID string
}

type Users struct {
	gorm.Model
	DisplayName string
	UserName    string
	Password    string
	Email       string
}

type CommandLog struct {
	Command  string
	SentBy   string
	Response string
}
