package router

import (
	gin "github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	
	initRoutes(router)

	router.Run()
}