package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marema31/tirith/controllers/videogame"
)

func Register(r *gin.Engine) {
	videogame.Register(r.Group("/videogame"))
	videogame.RegisterPlatform(r.Group("/platform"))
	videogame.RegisterType(r.Group("/type"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "welcome to tirith"})
	})

}
