package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test/internal/config"
)

// DB 用于存储 GORM 数据库连接实例
var DB *gorm.DB

// SetDB 设置数据库连接
func SetDB(db *gorm.DB) {
	DB = db
}

type datastore struct {
	db *gorm.DB
}

var DataStore Factory

// GetFactory 返回一个数据工厂
func GetFactory() (Factory, error) {
	// 获取 mysql 配置信息
	mysqlConf := config.Conf.Mysql

	// 连接到 MySQL 数据库
	db, err := getGormDB(mysqlConf)
	if err != nil {
		return nil, err
	}

	// 创建 datastore 实例
	DataStore = &datastore{
		db: db,
	}

	return DataStore, nil
}

// getGormDB 初始化并返回 GORM 数据库连接
func getGormDB(mysqlConfig config.Mysql) (*gorm.DB, error) {
	// 构建 MySQL 数据源名称（DSN）
	dsn := mysqlConfig.User + ":" + mysqlConfig.Pass_word + "@tcp(" + mysqlConfig.Host + ":" + mysqlConfig.Port + ")/" + mysqlConfig.Db_name +
		"?charset=" + mysqlConfig.Charset + "&parseTime=true&loc=Local&collation=utf8mb4_unicode_ci"

	// 使用 GORM 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 设置数据库连接池（可选）
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(1500)

	// 可选：启用 SQL 打印
	// db = db.Debug()  // 开启调试模式（打印 SQL）

	return db, nil
}
