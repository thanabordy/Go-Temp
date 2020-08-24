package router

import (
	"github.com/gin-gonic/gin"
)

// Gin Router to Gin framework
func Gin(engine *gin.Engine) {
	web(engine.Group("/"))
	api(engine.Group("/api"))
}