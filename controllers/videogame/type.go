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

// ListVGT is list videogame types endpoint handler
// @Summary list videogame types
// @Description list all videogame types
// @Tags type
// @Accept  json
// @Produce  json
// @Success 200 {array} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /type [get]
func ListVGT(c *gin.Context) {
	var p []videogame.VideoGameType
	models.DB.Find(&p)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

type VGTInput struct {
	Name string `json:"name" binding:"required"`
}

// CreateVGT is create videogame type endpoint handler
// @Summary create videogame type
// @Description create a videogame type
// @Tags type
// @Accept  json
// @Produce  json
// @Param user body VGTInput true "type attribute"
// @Success 201 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /type/ [post]
func CreateVGT(c *gin.Context) {
	var input VGTInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videogametype := videogame.VideoGameType{Name: input.Name}

	models.DB.Create(&videogametype)

	c.JSON(http.StatusCreated, gin.H{"data": videogametype})
}

// FindVGT is find videogame type endpoint handler
// @Summary find videogame type
// @Description find a videogame type
// @Tags type
// @Accept  json
// @Produce  json
// @Param id path int true "type to find"
// @Success 200 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /type/{id} [get]
func FindVGT(c *gin.Context) {
	var p videogame.VideoGameType

	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

// UpdateVGT is update videogame type endpoint handler
// @Summary update a videogame type
// @Description create a videogame type
// @Tags type
// @Accept  json
// @Produce  json
// @Param id path int true "find type"
// @Param user body VGTInput true "type to update"
// @Success 200 {object} videogame.VideoGamePlateform
// @Failure 400 {object} string
// @Router /type/{id} [put]
func UpdateVGT(c *gin.Context) {
	var p videogame.VideoGameType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input VGTInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&p).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

// DeleteVGT is delete videogame type endpoint handler
// @Summary delete a videogame type
// @Description delete a videogame type
// @Tags type
// @Accept  json
// @Produce  json
// @Param id path int true "type to delete"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /type/{id} [delete]
func DeleteVGT(c *gin.Context) {
	// Get model if exist
	var p videogame.VideoGameType
	if err := models.DB.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&p)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
