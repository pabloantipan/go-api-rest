package interfaces

import (
	"github.com/gin-gonic/gin"
)

type ProfileHandler interface {
	GetByID(c *gin.Context)
}
