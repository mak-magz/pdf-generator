package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mak-magz/pdf-generator/internal/pdf"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	greeter := pdf.Greet

	router.GET("/", greeter)

	return router
}
