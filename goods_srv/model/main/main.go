package main

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"shop_srvs/goods_srv/model"
	"time"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)

	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	// 1.参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3307)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	// 2.设置全局的logger，这个logger可以记录每条sql语句的执行
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢SQL阈值
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	// 3.连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 4.迁移schema
	//_ = db.AutoMigrate(&model.Category{}, &model.Brands{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})
	_ = db.AutoMigrate(&model.Goods{})

}
