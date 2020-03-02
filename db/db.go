package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDBConnect() (*gorm.DB, error) {
	dbconn, err := gorm.Open("postgres", "host=172.16.0.4 port=5432 user=itop dbname=itopmiddleware password=goodluck@123 sslmode=disable")
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}
