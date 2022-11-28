package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/config"
)

func (a *APIHandler) GetCost(c *gin.Context) {
	c.String(200, fmt.Sprintf("%d", config.SmokeCostPerOrder))
}
