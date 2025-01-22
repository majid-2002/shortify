package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"shortify/controllers" 
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	controllers.SetupRoutes(r, db)
}