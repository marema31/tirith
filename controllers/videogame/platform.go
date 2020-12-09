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

// ListVGP is list videogame platforms endpoint handler
// @Summary list videogame platforms
// @Description list all videogame platforms
// @Tags platform
// @Accept  json
// @Produce  json
// @Success 200 {array} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /platform [get]
func ListVGP(c *gin.Context) {
	var p []videogame.VideoGamePlateform
	models.DB.Find(&p)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

type vGPInput struct {
	Name string `json:"name" binding:"required"`
}

// CreateVGP is create videogame platform endpoint handler
// @Summary create videogame platform
// @Description create a videogame platform
// @Tags platform
// @Accept  json
// @Produce  json
// @Param user body vGPInput true "platform attribute"
// @Success 201 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /platform/ [post]
func CreateVGP(c *gin.Context) {
	var input vGPInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videogameplatform := videogame.VideoGamePlateform{Name: input.Name}

	models.DB.Create(&videogameplatform)

	c.JSON(http.StatusCreated, gin.H{"data": videogameplatform})
}

// FindVGP is find videogame platform endpoint handler
// @Summary find videogame platform
// @Description find a videogame platform
// @Tags platform
// @Accept  json
// @Produce  json
// @Param id path int true "platform to find"
// @Success 200 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /platform/{id} [get]
func FindVGP(c *gin.Context) {
	var p videogame.VideoGamePlateform

	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

// UpdateVGP is update videogame platform endpoint handler
// @Summary update a videogame platform
// @Description create a videogame platform
// @Tags platform
// @Accept  json
// @Produce  json
// @Param id path int true "find platform"
// @Param user body vGPInput true "platform to update"
// @Success 200 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /platform/{id} [put]
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

// DeleteVGP is delete videogame platform endpoint handler
// @Summary delete a videogame platform
// @Description delete a videogame platform
// @Tags platform
// @Accept  json
// @Produce  json
// @Param id path int true "platform to delete"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /platform/{id} [delete]
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
