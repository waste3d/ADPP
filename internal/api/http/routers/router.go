package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/waste3d/ADPP/internal/api/http/v1"
)

func InitRouters(handler *v1.Handler) *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.Use(gin.Recovery())

	router.POST("/v1/jobs", handler.CreateJobHandler)

	return router
}
