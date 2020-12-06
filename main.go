package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marema31/tirith/controllers"
	"github.com/marema31/tirith/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	controllers.Register(r)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
