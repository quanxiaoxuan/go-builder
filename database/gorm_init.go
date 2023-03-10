package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// 初始化Gorm
func InitGormDB(config *Config) {
	if DB == nil {
		var err error
		DB, err = CreateGormDB(config)
		if err != nil {
			log.Error("=== 数据库连接失败 : ", err)
		} else {
			log.Info("=== 数据库连接成功 ===")
		}
	}
}

// 创建数据库连接
func CreateGormDB(config *Config) (gormDB *gorm.DB, err error) {
	gormDB = GetGormDB(GetDSN(config), config.Dialect)
	var sqlDB *sql.DB
	sqlDB, err = gormDB.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetMaxOpenConns(60)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 是否打印SQL
	if config.Debug {
		gormDB = gormDB.Debug()
	}
	return
}

const (
	Mysql    = "mysql"
	Postgres = "postgres"
)

func GetDSN(config *Config) (dsn string) {
	switch strings.ToLower(config.Dialect) {
	case Mysql:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.Database) +
			"?clientFoundRows=false&parseTime=true&timeout=1800s&charset=utf8&collation=utf8_general_ci&loc=Local"
	case Postgres:
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host,
			config.Port,
			config.UserName,
			config.Password,
			config.Database)
	}
	return
}

// 根据dsn生成gormDB
func GetGormDB(dsn, dialect string) (gormDb *gorm.DB) {
	var err error
	switch strings.ToLower(dialect) {
	case Mysql:
		gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	case Postgres:
		gormDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	}
	if err != nil {
		panic(err)
	}
	return
}
