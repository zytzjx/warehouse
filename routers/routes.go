package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zytzjx/warehouse/controller"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.Use(gin.Logger()) //gin.LoggerWithConfig(gin.LoggerConfig{})
	r.Use(gin.Recovery())
	r.StaticFile("/favicon.ico", "static/favicon.ico")
	r.GET("/index", controller.Home)
	r.GET("/index.html", controller.Home)

	r.GET("/", controller.CheckAuth, controller.LoginOrUI)
	r.GET("/login", controller.Login)
	r.POST("/login", controller.FDLogin)
	r.GET("/logout", controller.Logout)
	r.POST("/signup", controller.Signup)
	r.GET("/registration", controller.Registration)

	r.GET("/editui", controller.EditNew)
	r.GET("/devices", controller.FindAllHandset)
	r.POST("/device", controller.RequireAuth, controller.Device)
	r.POST("/updatenote", controller.RequireAuth, controller.UpdateNote)
	r.POST("/returnwarehouse", controller.RequireAuth, controller.ReturnWarehouse)

	r.GET("/admin", controller.RequireAuth, controller.AdminHome)
	r.POST("/deletedevices", controller.RequireAuth, controller.DeleteDevice)
	r.POST("/changeborrower", controller.RequireAuth, controller.ChangeBorrower)
	return r
}
