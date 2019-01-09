package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitCurrDb(t *testing.T) {
	curDb := InitCurrDb(&Configuration{
		Env:"dev",

	})
	fmt.Printf("%#v", curDb)
}


// 创建db
func TestCreateDb(t *testing.T) {
	err := CreateDb(&Configuration{
		Env:"dev",
		Host:"127.0.0.1",
		Port:3306,
		Password:"123456",
		UserName:"root",
		MaxIdleConn:10,
		MaxOpenConn:100,
		Prefix:"ex_",
		Type:"mysql",
		Name:"user",
	})
	assert.NoError(t, err)
}

// 删除db
func TestDropDb(t *testing.T) {
	err := DropDb(&Configuration{
		Env:"dev",
		Host:"127.0.0.1",
		Port:3306,
		Password:"123456",
		UserName:"root",
		MaxIdleConn:10,
		MaxOpenConn:100,
		Prefix:"ex_",
		Type:"mysql",
		Name:"user",
	})
	assert.NoError(t, err)
}

// 自动迁移db
func TestAutoMigrateDb(t *testing.T) {
	err := AutoMigrateDb(&Configuration{
		Env:"dev",
		Host:"127.0.0.1",
		Port:3306,
		Password:"123456",
		UserName:"root",
		MaxIdleConn:10,
		MaxOpenConn:100,
		Prefix:"ex_",
		Type:"mysql",
		Name:"user",
	})
	assert.NoError(t, err)
}