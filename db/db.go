package db

import (
	"github.com/go-pg/pg"
	"errors"
)

var dbConfig map[string]*pg.Options

func GetDbConn(name string) (*pg.DB, error) {
	if c, ok := dbConfig[name]; ok && c != nil {
		return pg.Connect(dbConfig[name]), nil
	}
	return nil, errors.New("no such db "+name)
}

func LoadConfig() {
	//this is a demo
	//In the production env, load config from a file is better
	dbConfig = map[string]*pg.Options{
		"demo": &pg.Options{
			Database: "demo",
			User: "dushengchen",
		},
	}
}