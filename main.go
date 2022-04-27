package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Note struct {
	Id      int
	Type    string
	Content string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	// // dsn := os.Getenv(key:"DBconnectionStr")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// fmt.Println(db, err)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
