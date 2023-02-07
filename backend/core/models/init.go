package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine failed: %v", err)
	}
	return engine
}
