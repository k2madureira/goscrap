package repositories

import (
	"errors"
	"time"

	"github.com/k2madureira/goscrap/helper"
	models "github.com/k2madureira/goscrap/internal/models"
	"github.com/k2madureira/goscrap/internal/validations"
	"gorm.io/gorm"
)

type PokemonRepository interface {
	Create(pokemon models.Pokemon)
	Update(pokemon models.Pokemon)
	Delete(pokemonId int)
	FindById(pokemonId int) (pokemon models.Pokemon, err error)
	FindAll() []models.Pokemon
}

type PokemonRepositoryImpl struct {
	Db *gorm.DB
}

func NewPokemonRepositoryImpl(Db *gorm.DB) PokemonRepository {
	return &PokemonRepositoryImpl{Db: Db}
}

func (p *PokemonRepositoryImpl) Create(pokemons models.Pokemon) {
	result := p.Db.Create(&pokemons)
	helper.ErrorPanic(result.Error)
}

func (p *PokemonRepositoryImpl) Delete(pokemonId int) {
	var pokemon models.Pokemon
	result := p.Db.Where("id = ?", pokemonId).Delete(&pokemon)
	helper.ErrorPanic(result.Error)
}

func (p *PokemonRepositoryImpl) FindAll() []models.Pokemon {
	var pokemons []models.Pokemon
	result := p.Db.Find(&pokemons)
	helper.ErrorPanic(result.Error)
	return pokemons
}

func (p *PokemonRepositoryImpl) FindById(pokemonId int) (pokemons models.Pokemon, err error) {
	var pokemon models.Pokemon
	result := p.Db.Find(&pokemons, pokemonId)
	if result != nil {
		return pokemon, nil
	} else {
		return pokemon, errors.New("pokemon-not-found")
	}
}

func (p *PokemonRepositoryImpl) Update(pokemons models.Pokemon) {
	var updatePokemon = validations.UpdatePokemonRequest{
		Id:        pokemons.Id,
		Name:      pokemons.Name,
		Type:      pokemons.Type,
		Region:    pokemons.Region,
		UpdatedAt: time.Now(),
	}
	result := p.Db.Model(&pokemons).Updates(updatePokemon)
	helper.ErrorPanic(result.Error)
}
