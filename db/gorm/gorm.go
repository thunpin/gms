package gorm

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var m sync.Mutex
var db *gorm.DB

func Init(driverName string, dataSourceName string) error {
	m.Lock()
	defer m.Unlock()
	if db != nil {
		return nil
	}

	currentDB, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	db = currentDB
	return nil
}

func Close() {
	m.Lock()
	defer m.Unlock()
	if db != nil {
		db.Close()
		db = nil
	}
}
