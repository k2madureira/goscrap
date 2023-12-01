package getpage

import (
	"encoding/json"
	"log"
	"net/http"
)

type Page struct {
	Url string `json:"url"`
}

func GetPageController(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	page := new(Page)
	if err := json.NewDecoder(req.Body).Decode(page); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Please enter a correct Todo!!"))
		return
	}

	response := make(map[string]string)
	response["msg"] = SearchPage(page.Url)

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error handling JSON marshal")
		return
	}

	res.WriteHeader(200)
	res.Write(jsonResp)
}
