package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	getpage "github.com/k2madureira/goscrap/internal/modules/pages/getPage"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/{id}", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("web: index router"))
	})
	router.Post("/", func(res http.ResponseWriter, req *http.Request) {
		getpage.GetPageController(res, req)
	})

	return router
}
