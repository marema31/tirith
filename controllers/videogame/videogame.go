package videogame

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marema31/tirith/models"
	"github.com/marema31/tirith/models/videogame"
)

func Register(r *gin.RouterGroup) {

	r.GET("/", ListVG)
	r.POST("/", CreateVG)
	r.GET("/:id", FindVG)
	r.PATCH("/:id", UpdateVG)
	r.DELETE("/:id", DeleteVG)
}

func ListVG(c *gin.Context) {
	var g []videogame.VideoGame
	models.DB.Find(&g)

	c.JSON(http.StatusOK, gin.H{"data": g})
}

type vGInput struct {
	Title       string `json:"title"`
	Support     string `json:"support"`
	ExtensionOf uint   `json:"extensionOf"`
	Platform    uint   `json:"platform"`
	Type        uint   `json:"type"`
}

func CreateVG(c *gin.Context) {
	var input vGInput
	var vg videogame.VideoGame

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	support := videogame.PHYSICAL
	if input.Support == "demat" {
		support = videogame.DEMAT
	}

	videogame := videogame.VideoGame{Title: input.Title, Support: support, Platform: input.Platform, Type: input.Type}

	if input.ExtensionOf != 0 {
		if err := models.DB.Where("id = ?", input.ExtensionOf).First(&vg).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record for extension not found!"})
			return
		}
		videogame.ExtensionOf = input.ExtensionOf
	}

	models.DB.Create(&videogame)

	c.JSON(http.StatusOK, gin.H{"data": videogame})
}

func FindVG(c *gin.Context) {
	var vg videogame.VideoGame

	if err := models.DB.Where("id = ?", c.Param("id")).First(&vg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vg})
}

func UpdateVG(c *gin.Context) {
	var vg videogame.VideoGame
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input vGInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&vg).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": vg})
}

func DeleteVG(c *gin.Context) {
	// Get model if exist
	var vg videogame.VideoGame
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&vg)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
