package model

import (
	"errors"
	_ "config/logger"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var conns = make(map[string]*gorm.DB)
var mutex sync.Mutex

var Db *gorm.DB
func InitCurrDb(cfg *Configuration) (*gorm.DB){
	mutex.Lock()
	defer mutex.Unlock()
	tmpDb := conns[cfg.Name]
	if tmpDb == nil {
		driver := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.UserName,
			cfg.Password,
			cfg.Host,
			cfg.Prefix + cfg.Name + "_" + cfg.Env)
		fmt.Printf("%#v", driver)
		log.Info("init db", "driver info", driver)
		openedDb, err := gorm.Open(cfg.Type, driver)
		if err != nil {
			//panic("db connect err:" + err.Error())
			log.Fatal("db connect err:" + err.Error())
		}
		tmpDb = openedDb
		tmpDb.DB().SetMaxIdleConns(cfg.MaxIdleConn)
		tmpDb.DB().SetMaxOpenConns(cfg.MaxOpenConn)
		// 避免久了不使用，导致连接被mysql断掉的问题
		tmpDb.DB().SetConnMaxLifetime(time.Hour * 2)
		conns[cfg.Prefix + cfg.Name + "_" + cfg.Env] = tmpDb
	}
	Db = tmpDb
	return tmpDb
}

// 得到当前的db信息
func GetDbConfig(cfg *Configuration) (*gorm.DB) {
	return InitCurrDb(cfg)
}

// 创建db
func CreateDb(cfg *Configuration) (error) {
	fmt.Printf("create db: %s", cfg.Name)
	log.Info("create db", "db", cfg.Name)
	if err := checkCfg(cfg); err != nil {
		return err
	}
	cStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		"information_schema")
	log.Info("create db", "create sql", cStr)
	fmt.Printf("%#v", cStr)
	openedDb, err := gorm.Open(cfg.Type, cStr)
	if err != nil {
		panic("create db ============= db connect err:" + err.Error())
	}
	createDbSQL := "CREATE DATABASE IF NOT EXISTS " + cfg.Prefix + cfg.Name + "_" + cfg.Env + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"

	if err = openedDb.Exec(createDbSQL).Error; err != nil {
		fmt.Println("create db, exec error:"+ err.Error())
		log.Info("create db, exec error:"+ err.Error())
		return errors.New(fmt.Sprintf("%s db create error", cfg.Name))
	}
	return nil
}

// 删除db
func DropDb(cfg *Configuration) error {
	fmt.Printf("drop db: %s", cfg.Name)
	log.Info("drop db: %s", cfg.Name)
	if err := checkCfg(cfg); err != nil {
		return err
	}
	openedDb, err := gorm.Open(cfg.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.UserName,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			"information_schema"))
	if err != nil {
		panic("drop db ============= db connect err:" + err.Error())
		log.Fatal("drop db ============= db connect err:" + err.Error())
	}
	dropDbSQL := "DROP DATABASE IF EXISTS " + cfg.Prefix + cfg.Name + "_" + cfg.Env + ";"
	if err = openedDb.Exec(dropDbSQL).Error; err != nil {
		fmt.Println("drop db, exec error:"+ err.Error())
		log.Fatal("drop db, exec error:"+ err.Error())
		return errors.New(fmt.Sprintf("%s db create error", cfg.Name))
	}
	return nil
}

func AutoMigrateDb(cfg *Configuration) error {
	db := InitCurrDb(cfg)
	db.AutoMigrate(&UserInfo{})
	return nil
}

func checkCfg(cfg *Configuration) error {
	if cfg.Host == "" {
		return errors.New("db host is empty")
	}
	if cfg.Password == "" {
		return errors.New("db password is empty")
	}
	if cfg.UserName == "" {
		return errors.New("db username is empty")
	}
	if cfg.Port <= 0 {
		return errors.New("db port is error")
	}
	return nil
}