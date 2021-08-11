package router

import (
	"github.com/arielril/attack-graph-flow-weight-api/api/http/weightapi"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	rt := gin.Default()

	rt.POST("/v1/compute", weightapi.ComputeProbability)

	return rt
}
