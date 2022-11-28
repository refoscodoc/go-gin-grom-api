package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/refoscodoc/go-gin-grom-api/dtos"
	"github.com/refoscodoc/go-gin-grom-api/models"
	"github.com/refoscodoc/go-gin-grom-api/services"
	"net/http"
)

func FindGuitars(c *gin.Context) {
	var guitars []models.Guitar
	services.DB.Find(&guitars)

	c.JSON(http.StatusOK, gin.H{"data": guitars})
}

func CreateGuitar(c *gin.Context) {
	var input dtos.CreateGuitar
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guitar := models.Guitar{Manufacturer: input.Manufacturer, Model: input.Model, Year: input.Year}
	services.DB.Create(&guitar)

	c.JSON(http.StatusOK, gin.H{"data": guitar})
}
