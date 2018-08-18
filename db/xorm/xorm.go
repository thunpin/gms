package xorm

import (
	"sync"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine
var sessionMap = make(map[string]*xorm.Session)
var m sync.Mutex

func Initialize(driverName, dataSourceName string) error {
	m.Lock()
	defer m.Unlock()
	if engine != nil {
		return nil
	}

	currentEngine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return err
	}

	engine = currentEngine
	return nil
}

func Close() error {
	return engine.Close()
}

func BeginSession(uuid string) {
	_, contains := sessionMap[uuid]
	if !contains {
		sessionMap[uuid] = engine.NewSession()
	}
}

func GetSession(uuid string) *xorm.Session {
	BeginSession(uuid)
	return sessionMap[uuid]
}

func Commit(uuid string) error {
	session := GetSession(uuid)
	return session.Commit()
}

func Rollback(uuid string) error {
	session := GetSession(uuid)
	return session.Rollback()
}

func CloseSession(uuid string) {
	session := GetSession(uuid)
	Rollback(uuid)
	session.Close()
}
