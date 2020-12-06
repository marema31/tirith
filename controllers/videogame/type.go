package videogame

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marema31/tirith/models"
	"github.com/marema31/tirith/models/videogame"
)

func RegisterType(r *gin.RouterGroup) {

	r.GET("/", ListVGT)
	r.POST("/", CreateVGT)
	r.GET("/:id", FindVGT)
	r.PATCH("/:id", UpdateVGT)
	r.DELETE("/:id", DeleteVGT)
}

func ListVGT(c *gin.Context) {
	var p []videogame.VideoGamePlateform
	models.DB.Find(&p)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

type vGTInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateVGT(c *gin.Context) {
	var input vGTInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videogametype := videogame.VideoGamePlateform{Name: input.Name}

	models.DB.Create(&videogametype)

	c.JSON(http.StatusOK, gin.H{"data": videogametype})
}

func FindVGT(c *gin.Context) {
	var p videogame.VideoGamePlateform

	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

func UpdateVGT(c *gin.Context) {
	var p videogame.VideoGamePlateform
	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input vGTInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&p).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

func DeleteVGT(c *gin.Context) {
	// Get model if exist
	var p videogame.VideoGamePlateform
	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&p)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
