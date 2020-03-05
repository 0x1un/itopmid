package support

import "github.com/jinzhu/gorm"

type ItopMidContext struct {
	db *gorm.DB
}

func (self *ItopMidContext) GetDB() *gorm.DB {
	return self.db
}
