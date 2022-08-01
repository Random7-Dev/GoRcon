package config

import (
	"html/template"
	"log"

	"github.com/Random-7/GoRcon/pkg/database"
	"github.com/Random-7/GoRcon/pkg/rcon"
	"github.com/alexedwards/scs/v2"
)

//AppConfig hold application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	Rcon          rcon.Connection
	DbSession     database.Session
	Version       string
}
