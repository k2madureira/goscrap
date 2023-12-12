package server

import (
	"net/http"

	pokesRouter "github.com/k2madureira/goscrap/internal/modules/pokes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	pokes := r.Group("/pokes")
	pokesRouter.Router(pokes)

	r.Use(cors.Default())
	return r
}
