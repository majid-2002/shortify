package controllers

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"shortify/models"
)

var db *gorm.DB
var err error

func SetupRoutes(r *gin.Engine, database *gorm.DB) {
	db = database

	r.GET("/", HomeHandler)
	r.POST("/shorten", CreateShortURL)
	r.GET("/:shortened", RedirectShortURL)
	r.PUT("/:shortened", UpdateShortURL)
	r.DELETE("/:shortened", DeleteShortURL)
	r.GET("/:shortened/stats", GetStats)
}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Shortify!",
	})
}

func CreateShortURL(c *gin.Context) {
	var url models.URL

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url.Shortened = generateShortenedURL(url.Original)

	if err := db.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"original":  url.Original,
		"shortened": url.Shortened,
	})
}

func generateShortenedURL(original string) string {
	return fmt.Sprintf("%x", time.Now().UnixNano())
}


func RedirectShortURL(c *gin.Context) {
	shortened := c.Param("shortened")

	var url models.URL
	if err := db.Where("shortened = ?", shortened).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	url.AccessCount++
	db.Save(&url)

	c.Redirect(http.StatusMovedPermanently, url.Original)
}

func UpdateShortURL(c *gin.Context) {
	shortened := c.Param("shortened")

	var url models.URL
	if err := db.Where("shortened = ?", shortened).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	var input models.URL
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url.Original = input.Original
	db.Save(&url)

	c.JSON(http.StatusOK, gin.H{
		"original":  url.Original,
		"shortened": url.Shortened,
	})
}


func DeleteShortURL(c *gin.Context) {
	shortened := c.Param("shortened")

	if err := db.Where("shortened = ?", shortened).Delete(&models.URL{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL deleted"})
}


func GetStats(c *gin.Context) {
	shortened := c.Param("shortened")

	var url models.URL
	if err := db.Where("shortened = ?", shortened).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortened":   url.Shortened,
		"access_count": url.AccessCount,
	})
}