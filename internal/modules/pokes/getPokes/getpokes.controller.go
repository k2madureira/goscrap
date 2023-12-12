package getpage

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	strPage := c.Query("page")
	page, err := strconv.Atoi(strPage)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprint("https://scrapeme.live/shop/page/", page)
	response := List(url, page)
	c.JSON(http.StatusOK, response)
}
