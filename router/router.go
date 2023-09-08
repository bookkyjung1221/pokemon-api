package router

import (
	"github.com/bookkyjung1221/pokemon/controllers"
	"github.com/gin-gonic/gin"
)

func PokemonRoutes(routes *gin.Engine) {

	routes.GET("/healthcheck", controllers.Healthcheck)

	pokemon := routes.Group("/api/v1")
	{
		pokemon.GET("/pokemons", controllers.GetPokemons)
		pokemon.GET("/pokemon/:id", controllers.GetPokemon)
		pokemon.POST("/pokemon", controllers.CreatePokemon)
		pokemon.PUT("/pokemon/:id", controllers.UpdatePokemon)
		pokemon.DELETE("/pokemons", controllers.DeletePokemons)
		pokemon.DELETE("/pokemon/:id", controllers.DeletePokemon)
	}
}
