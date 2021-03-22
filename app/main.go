package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	engine:= gin.Default()
	db := gormConnect()
	defer db.Close()
	engine.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3000")
}

func gormConnect() *gorm.DB {
	host := "127.0.0.1"
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "root"
	PROTOCOL := "tcp( " + host + ":3306)"
	DBNAME   := "sendQ_database"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
