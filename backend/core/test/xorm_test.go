package test

import (
	"backend/core/models"
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	dsn := "root:123456@(39.101.1.158:3003)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(data) // -> byte array

	if err != nil {
		t.Fatal(err)
	}

	dst := new(bytes.Buffer)
	if err = json.Indent(dst, b, "", " "); err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
