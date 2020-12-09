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

// ListVG is list videogame  endpoint handler
// @Summary list videogame
// @Description list all videogame
// @Tags videogame
// @Accept  json
// @Produce  json
// @Success 200 {array} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /videogame [get]
func ListVG(c *gin.Context) {
	var g []videogame.VideoGame
	models.DB.Find(&g)

	c.JSON(http.StatusOK, gin.H{"data": g})
}

type VGInput struct {
	Title       string `json:"title" binding:"required"`
	Support     string `json:"support" binding:"required"`
	ExtensionOf uint   `json:"extensionOf"`
	Platform    uint   `json:"platform" binding:"required"`
	Type        uint   `json:"type" binding:"required"`
}

// CreateVG is create videogame  endpoint handler
// @Summary create videogame
// @Description create a videogame
// @Tags videogame
// @Accept  json
// @Produce  json
// @Param user body VGInput true " attribute"
// @Success 201 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /videogame/ [post]
func CreateVG(c *gin.Context) {
	var input VGInput
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

	c.JSON(http.StatusCreated, gin.H{"data": videogame})
}

// FindVG is find videogame  endpoint handler
// @Summary find videogame
// @Description find a videogame
// @Tags videogame
// @Accept  json
// @Produce  json
// @Param id path int true " to find"
// @Success 200 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /videogame/{id} [get]
func FindVG(c *gin.Context) {
	var vg videogame.VideoGame

	if err := models.DB.Where("id = ?", c.Param("id")).First(&vg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vg})
}

// UpdateVG is update videogame  endpoint handler
// @Summary update a videogame
// @Description create a videogame
// @Tags videogame
// @Accept  json
// @Produce  json
// @Param id path int true "find "
// @Param user body VGInput true " to update"
// @Success 200 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /videogame/{id} [put]
func UpdateVG(c *gin.Context) {
	var vg videogame.VideoGame
	if err := models.DB.Where("id = ?", c.Param("id")).First(&vg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input VGInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&vg).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": vg})
}

// DeleteVG is delete videogame  endpoint handler
// @Summary delete a videogame
// @Description delete a videogame
// @Tags videogame
// @Accept  json
// @Produce  json
// @Param id path int true " to delete"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /videogame/{id} [delete]
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
