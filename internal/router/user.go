package router

import (
	"github.com/AlexandreMaycon/api-social-midia.git/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	r := route.Group("/user")
	{
		r.POST("/", handler.CreateUser)
		r.GET("/", handler.GetAllUsers)
		r.PUT("/", handler.UpdateUser)
		r.DELETE("/", handler.DeleteUser)
	}
}
