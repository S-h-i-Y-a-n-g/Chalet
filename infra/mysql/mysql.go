package mysql

import (
	"chalet/pkg/entity"
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	idle        = 300
	connections = 500
)

var db *gorm.DB

func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

func Init(mysqlConfig entity.MysqlConfig) {
	connStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Db)
	log.Println(connStr)

	var err error
	db, err = gorm.Open(mysql.Open(connStr))
	if err != nil {
		log.Panicf("connect mysql failed: %v\n", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(idle)
	sqlDB.SetMaxOpenConns(connections)
	sqlDB.SetConnMaxLifetime(time.Second)

	log.Println("mysql init done")
}
