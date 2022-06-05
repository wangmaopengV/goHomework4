package api

import (
	"goHomework4/internal/service"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/test", service.TestHandler)
}
