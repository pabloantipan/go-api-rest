package handler

import (
	"fmt"
	"net/http"
	"poc/internal/domain"
	"poc/internal/handler/interfaces"
	serviceInterfaces "poc/internal/service/interfaces"

	"github.com/gin-gonic/gin"
)

type StatHandler struct {
	service serviceInterfaces.StatService
}

func NewStatHandler(s serviceInterfaces.StatService) interfaces.StatHandler {
	return &StatHandler{service: s}
}

// CreateStat godoc
// @Summary Create a new stat
// @Description Create a new stat with the provided input data
// @Tags stats
// @Accept json
// @Produce json
// @Param stat body domain.Stat true "Stat object"
// @Success 201 {object} domain.Stat
// @Failure 400 {object} gin.H
// @Router /stats [post]
func (h *StatHandler) Create(c *gin.Context) {
	var newStat domain.Stat
	if err := c.ShouldBindJSON(&newStat); err != nil {
		fmt.Println("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Creating stat 1: ", newStat)

	stat, err := h.service.Create(newStat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, stat)
}

// GetStats godoc
// @Summary Get all stats
// @Description Get all stats
// @Tags stats
// @Produce json
// @Success 200 {array} domain.Stat
// @Failure 500 {object} gin.H
// @Router /stats [get]
func (h *StatHandler) GetAll(c *gin.Context) {
	stats, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *StatHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	stat, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stat)
}

func (h *StatHandler) GetByUserID(c *gin.Context) {
	id := c.Param("id")
	stat, err := h.service.GetByUserID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stat)
}

func (h *StatHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var updatedStat domain.Stat
	if err := c.ShouldBindJSON(&updatedStat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedStat.ID = id

	stat, err := h.service.Update(updatedStat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stat)
}

func (h *StatHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
