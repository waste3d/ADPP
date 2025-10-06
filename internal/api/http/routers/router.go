package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/waste3d/ADPP/internal/api/http/v1"
	"github.com/waste3d/ADPP/internal/storage/postgres"
)

func InitRouters(jobStorage *postgres.Storage, handler *v1.Handler) *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.Use(gin.Recovery())

	router.POST("/v1/jobs", handler.CreateJobHandler)

	return router
}
