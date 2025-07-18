package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := InitDB(); err != nil {
		log.Fatal(err)
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	if err := InitRedis(); err != nil {
		log.Fatal(err)
	}
	//连接池不会自动关闭，连接池设计的目的是为了合理复用

	r := gin.Default()
	//gin框架的入口返回engine用于曹操作gin
	v1 := r.Group("/api/v1")
	{
		v1.GET("/students", GetStudents) //获取所有
		v1.GET("/students/:name", GetStudentByName)
		v1.PUT("/students/:name", UpdateStudent) //更新
		v1.DELETE("/students/:name", DeleteStudent)
		//删除学byID
		v1.POST("/students", CreateStudent)
	}

	r.Run()
}
