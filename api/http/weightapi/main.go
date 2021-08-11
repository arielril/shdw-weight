package weightapi

import (
	"net/http"

	"github.com/arielril/attack-graph-flow-weight-api/internal/weight"
	"github.com/gin-gonic/gin"
)

func ComputeProbability(c *gin.Context) {
	result := weight.Compute()

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
