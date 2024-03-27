package main

import (
	"Dexun/config"
	"Dexun/controller"
	"Dexun/middlewares"
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
		//v1.POST("/adminlogin", controller.AdminLogin)

		v1.POST("/register", controller.Register)
		v1.POST("/login", controller.Login)
		v1.GET("/getcode", controller.GetCode)
		v1.GET("/logout", controller.Logout)
		v1.POST("/role", controller.GetAccountRole)
		v1.POST("/certtype", controller.GetAccountCertType)

		v1.GET("/account/:account", controller.Info)

		v1.POST("/accounts", controller.AccountList)

		v1.POST("/cert/person", controller.PersonalCert)
		v1.POST("/cert/corp", controller.CorpCert)
		v1.POST("/cert/card", controller.CardCert)
		v1.POST("/certs/person", middlewares.AdminRequired, controller.PersonalCertList)
		v1.POST("/certs/corp", middlewares.AdminRequired, controller.CorpCertList)
		v1.POST("/cert/corp/:id/status", middlewares.AdminRequired, controller.UpdateCorpCertStatus)
		v1.POST("/cert/person/:id/status", middlewares.AdminRequired, controller.UpdatePersonCertStatus)
		v1.POST("/cert/person/status", controller.GetPersonStatus)
		v1.POST("/cert/corp/status", controller.GetCorpStatus)

	}
	app.Run(fmt.Sprintf(":%d", *httpPort))
}
