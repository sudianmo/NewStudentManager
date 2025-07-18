package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var gormDB *gorm.DB //定义为全局变量，包级别的变量
var err error
var rdb *redis.Client

type Student struct {
	Name  string `json:"name"`
	Tel   int    `json:"tel"`
	Study string `json:"study"`
	Id    int    `json:"id" gorm:"primaryKey"`
}

func (Student) TableName() string {
	return "students"
}

func InitDB() error {
	dsn := "root:314159@tcp(127.0.0.1:3306)/Student_sql"
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := gormDB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)           //最大空闲连接数量
	sqlDB.SetMaxOpenConns(100)          //最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //连接最大生存时间

	fmt.Println("数据库连接池配置- MICaxdleConns;%d ,MaxOpenConns :%d\n,10,100")

	//驱动为Mysql，ds
	err := sqlDB.Ping()
	if err != nil {
		return err
	}

	if err := gormDB.AutoMigrate(&Student{}); err != nil {
		return err
	}
	fmt.Println("Successfully connected to DB")
	return nil

}

func InitRedis() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	if err := rdb.Ping(context.Background()); err != nil {
		return fmt.Errorf("Redis connect error: %v", err)
	}
	fmt.Println("Successfully connected to Redis")
	return nil
}
