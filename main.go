package main

import (
	// "fmt"
	"log"
	"net/http"

	"bwastartup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connection Established")

	// var users []user.User
	// length := len(users)

	// fmt.Println("Total Rows: ", length)

	// db.Find(&users)

	// length = len(users)
	// fmt.Println("Total Rows: ", length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("================================")
	// }
	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
} 

//handler dari db
func handler (c *gin.Context){
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK,users)

	// handler
	// service
	// repository
	// entity
}