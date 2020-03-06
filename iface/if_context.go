package iface

import (
	"github.com/jinzhu/gorm"
)

type ItopMidContext interface {
	OpenDB(sqlType string)
	CloseDB()
	GetDB() *gorm.DB
}
