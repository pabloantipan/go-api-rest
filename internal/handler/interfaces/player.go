package interfaces

import (
	"github.com/gin-gonic/gin"
)

type PlayerHandler interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
