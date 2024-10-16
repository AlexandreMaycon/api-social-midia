package router

import (
	"github.com/gin-gonic/gin"
)

func StartServer(port string) {
	r := gin.Default()

	UserRoutes(r)

	r.Run(port)
}
