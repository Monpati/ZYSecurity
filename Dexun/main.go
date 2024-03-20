package main

import (
	"Dexun/config"
	"Dexun/controller"
	"Dexun/utils"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	httpPort := flag.Int("port", 8080, "port")
	flag.Parse()

	db := config.InitDB()
	defer db.Close()
	config.RedisInit()

	app := gin.Default()
	app.Use(utils.Cors())

	v1 := app.Group("v1")
	{
		v1.POST("/adminlogin", controller.AdminLogin)

		v1.POST("/register", controller.Register)
		v1.POST("/login", controller.Login)
		v1.GET("/getcode", controller.GetCode)
		v1.GET("/logout", controller.Logout)

		v1.GET("/account/:account", controller.Info)

		v1.POST("/accounts", controller.AccountList)
	}
	app.Run(fmt.Sprintf(":%d", *httpPort))
}
