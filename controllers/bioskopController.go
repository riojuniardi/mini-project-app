package controllers

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop

	err := c.ShouldBindJSON(&bioskop)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&bioskop)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Data berhasil dibuat",
		"bioskop": bioskop,
	})
}

func GetBioskops(c *gin.Context) {
	var bioskops models.Bioskop

	config.DB.Find(&bioskops)
	c.JSON(http.StatusOK, gin.H{
		"message": "Data menampilkan semua data bioskop",
		"bioskop": bioskops,
	})
}

func GetBioskopById(c *gin.Context) {
	var bioskop models.Bioskop
	paramsId := c.Param("id")

	var eventData = config.DB.First(&bioskop, paramsId).Error
	if eventData != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Bioskop tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ditail data bioskop berhasil di ambil",
		"bioskop": bioskop,
	})
}

func UpdateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	paramsId := c.Param("id")

	var bioskopData = config.DB.First(&bioskop, paramsId).Error
	if bioskopData != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Bioskop tidak ditemukan",
		})
		return
	}

	var input models.Bioskop
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Model(&bioskop).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhsil di update",
		"bioskop": bioskop,
	})
}

func DeleteEvent(context *gin.Context) {
	var bioskop models.Bioskop
	paramsId := context.Param("id")

	var eventData = config.DB.First(&bioskop, paramsId).Error
	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event d=tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&bioskop)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data berhsil di Delete",
		"bioskop": bioskop,
	})
}
