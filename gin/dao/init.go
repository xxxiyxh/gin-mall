package dao

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)                 // 合适的最大连接数
	sqlDB.SetMaxIdleConns(20)                  // 合适的空闲连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 合适的连接生命周期

	_db = db

	_ = _db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      // 主库连接
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 从库连接
			Policy:   dbresolver.RandomPolicy{},                                    // 或者其他策略，如轮询等
		}))
	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
