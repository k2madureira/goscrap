package getpage

import (
	dto "github.com/k2madureira/goscrap/internal/modules/dtos/poke"
)

func Create(body *dto.Poke) dto.Poke {
	var response dto.Poke

	response.Id = 1
	response.Name = body.Name
	response.Type = body.Type
	response.Region = body.Region
	response.Image = body.Image

	return response
}
