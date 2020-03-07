package support

import (
	"github.com/0x1un/itopmid/iface"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ItopMidContext struct {
	db *gorm.DB
}

func (self *ItopMidContext) GetDB() *gorm.DB {
	return self.db
}

func (self *ItopMidContext) OpenDB(sqlType string) {
	var err error
	self.db, err = gorm.Open(sqlType, iface.CONFIG.GetDatabaseURL()+" sslmode=disable")
	if err != nil {
		iface.LOGGER.Panic(err.Error())
	}
	self.db.LogMode(false)
}

func (self *ItopMidContext) CloseDB() {
	if self.db != nil {
		err := self.db.Close()
		if err != nil {
			iface.LOGGER.Error(err.Error())
		}
	}
}
