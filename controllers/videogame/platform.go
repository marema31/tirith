package videogame

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marema31/tirith/models"
	"github.com/marema31/tirith/models/videogame"
)

func RegisterPlatform(r *gin.RouterGroup) {

	r.GET("/", ListVGP)
	r.POST("/", CreateVGP)
	r.GET("/:id", FindVGP)
	r.PATCH("/:id", UpdateVGP)
	r.DELETE("/:id", DeleteVGP)
}

func ListVGP(c *gin.Context) {
	var p []videogame.VideoGamePlateform
	models.DB.Find(&p)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

type vGPInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateVGP(c *gin.Context) {
	var input vGPInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videogameplatform := videogame.VideoGamePlateform{Name: input.Name}

	models.DB.Create(&videogameplatform)

	c.JSON(http.StatusOK, gin.H{"data": videogameplatform})
}

func FindVGP(c *gin.Context) {
	var p videogame.VideoGamePlateform

	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

func UpdateVGP(c *gin.Context) {
	var p videogame.VideoGamePlateform
	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input vGPInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&p).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

func DeleteVGP(c *gin.Context) {
	// Get model if exist
	var p videogame.VideoGamePlateform
	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&p)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
