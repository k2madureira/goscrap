package pokes

import (
	"github.com/gin-gonic/gin"
	createPokeController "github.com/k2madureira/goscrap/internal/modules/pokes/createPoke"
	getPokesController "github.com/k2madureira/goscrap/internal/modules/pokes/getPokes"
)

func Router(r *gin.RouterGroup) {

	r.GET("/", getPokesController.Handle)
	r.POST("/", createPokeController.Handle)

}
