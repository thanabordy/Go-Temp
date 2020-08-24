package router

import (
	"app/app/http/controller/example"

	"github.com/gin-gonic/gin"
)

func api(router *gin.RouterGroup) {
	router.GET("/testFunc", example.FuncTest)

	router.GET("/select", example.SelectUser)	//SelectAll
	router.GET("/select/:id", example.SelectUserByid)	// SelectOne Param
	router.POST("/create/user", example.CreateUser)	// Create
	router.POST("/update/user", example.UpdateUser)	// Updata 
	router.POST("/update/user/:id", example.UpdateUserByid)	// Updata Param
	router.DELETE("/delete/user", example.DeleteUser)	// Delete
	router.DELETE("/delete/user/:id", example.DeleteUserByid) 	// Delete Param
}
