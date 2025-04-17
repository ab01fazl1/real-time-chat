package cmd

import (
	"github.com/ab01fazl1/real-time-chat/internal/user"
	"github.com/gin-gonic/gin"
)


func Main() {
	router := gin.Default()

	router.POST("/register", user.Register)
	
	router.Run(":8080") // port 8080
}