package getpage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/k2madureira/goscrap/internal/modules/dtos/poke"
)

func Handle(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	req := new(dto.Poke)
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusInternalServerError, "Error in parseJson")
		return
	}

	response := Create(req)
	c.JSON(http.StatusCreated, response)
}
