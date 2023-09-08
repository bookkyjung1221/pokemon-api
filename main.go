package main

import (
	"fmt"
	"log"

	"github.com/bookkyjung1221/pokemon/config"
	"github.com/bookkyjung1221/pokemon/router"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Welcome to Pokemon API")

	config.LoadENV()

	err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	config.Seed()

	defer config.Close()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	router.PokemonRoutes(r)

	r.Run()
}
