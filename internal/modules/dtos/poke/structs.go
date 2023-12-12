package poke

type Body struct {
	Url   string `json:"url"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Response struct {
	Page       int    `json:"page"`
	TotalPages int    `json:"totalPages"`
	TotalItems int    `json:"totalItems"`
	Pokemons   []Body `json:"pokemons"`
}

type ListRequest struct {
	Page int `json:"page"`
}

type Poke struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Region string `json:"region"`
	Image  string `json:"image"`
}
