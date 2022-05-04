package router

import (
	"IM-chat/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	app := gin.Default()
	app.Use(Cors())
	v1 := app.Group("/")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.GET("user/selectUserById", api.UserSelect)
		v1.GET("message/selectMessageByUserIdAndPage", api.SelectMessageByUserIdAndPage)

		v1.GET("ws/:id", api.HandleWebsocketConnect)

	}
	//v2 := app.Group("/message")
	//{
	//	v2.GET("/selectMessageByUserIdAndPage", service.MessageSelect)
	//}
	return app
}
