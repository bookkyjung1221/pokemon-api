package controllers

import (
	"net/http"

	"github.com/bookkyjung1221/pokemon/models"
	"github.com/bookkyjung1221/pokemon/services"
	"github.com/gin-gonic/gin"
)

func GetPokemons(c *gin.Context) {
	pokemons, err := services.GetPokemons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  "Pokemons not found.",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gin.H{"pokemons": pokemons},
	})

}

func GetPokemon(c *gin.Context) {
	pokemon, err := services.GetPokemon(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  "Pokemon not found.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gin.H{"pokemon": pokemon},
	})

}

func CreatePokemon(c *gin.Context) {

	var req models.CreatePokemonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	err := services.CreatePokemon(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Pokemon created successfully",
	})

}

func UpdatePokemon(c *gin.Context) {

	var req models.UpdatePokemonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	err := services.UpdatePokemon(c.Param("id"), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pokemon updated successfully",
	})

}

func DeletePokemon(c *gin.Context) {
	err := services.DeletePokemon(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pokemon deleted successfully",
	})

}

func DeletePokemons(c *gin.Context) {
	err := services.DeletePokemons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pokemons deleted successfully",
	})

}
