package iface

import (
	"github.com/jinzhu/gorm"
)

type ItopMidContext interface {
	OpenDB()
	CloseDB()
	GetDB() *gorm.DB
}
