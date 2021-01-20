package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication start application
func StartApplication() {
	MapUrls()
	router.Run(":8080")
}
