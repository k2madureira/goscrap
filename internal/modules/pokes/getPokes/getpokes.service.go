package getpage

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	poke "github.com/k2madureira/goscrap/internal/modules/dtos/poke"
)

func List(url string, page int) poke.Response {
	var response poke.Response
	var pokeData []poke.Body
	c := colly.NewCollector()

	c.OnHTML("div.storefront-sorting", func(e *colly.HTMLElement) {

		str := e.ChildText("p")
		strSplited := strings.Split(str, "of ")
		strSplited = strings.Split(strSplited[1], " results")
		totalItems, err := strconv.Atoi(strSplited[0])
		if err != nil {
			panic(err)
		}
		response.TotalItems = totalItems

		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			digit := el.ChildText("a")
			if di, err := strconv.Atoi(digit); err == nil {
				response.TotalPages = di
			}

		})

	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		p := poke.Body{}

		p.Url = e.ChildAttr("a", "href")
		p.Image = e.ChildAttr("img", "src")
		p.Name = e.ChildText("h2")
		p.Price = e.ChildText(".price")

		pokeData = append(pokeData, p)
	})

	c.Visit(url)

	response.Pokemons = pokeData
	response.Page = page

	return response
}
