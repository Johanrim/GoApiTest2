package main

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

type MariaDB struct {
	DB *gorm.DB
}

func (db *MariaDB) connect(user, password, host, dbName string) (err error) {
	if user == "" {
		return errors.New("username is required")
	}
	if password == "" {
		return errors.New("password is required")
	}
	if dbName == "" {
		return errors.New("database name is required")
	}
	if host == "" || host == "localhost" || host == "127.0.0.1" {
		host = "tcp(localhost)"
	} else {
		host = "tcp(" + host + ")"
	}
	databaseStr := user + ":" + password + "@" + host + "/" + dbName + "?charset=utf8&parseTime=True"
	conn, err := gorm.Open("mysql", databaseStr)
	db.DB = conn
	return
}

func (db *MariaDB) close() {
	err := db.DB.Close()
	if err != nil {
		log.Println("Maria DB close connection error:", err)
	}
}

func (db *MariaDB) Migrate(migrations []*gormigrate.Migration) (err error) {
	m := gormigrate.New(db.DB, gormigrate.DefaultOptions, migrations)
	m.Migrate()
	return
}
