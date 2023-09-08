package models

type Pokemon struct {
	ID            int       `json:"id"`
	Num           string    `json:"num"`
	Name          string    `json:"name"`
	Img           string    `json:"img"`
	Type          []string  `json:"type"`
	Height        string    `json:"height"`
	Weight        string    `json:"weight"`
	Candy         string    `json:"candy"`
	Egg           string    `json:"egg"`
	Multipliers   []float64 `json:"multipliers"`
	Weaknesses    []string  `json:"weaknesses"`
	CandyCount    int       `json:"candy_count"`
	SpawnChance   float64   `json:"spawn_chance"`
	AvgSpawns     float64   `json:"avg_spawns"`
	SpawnTime     string    `json:"spawn_time"`
	PrevEvolution []struct {
		Num  string `json:"num"`
		Name string `json:"name"`
	} `json:"prev_evolution"`
	NextEvolution []struct {
		Num  string `json:"num"`
		Name string `json:"name"`
	} `json:"next_evolution"`
}

type CreatePokemonRequest struct {
	ID   int      `json:"id" binding:"required"`
	Num  string   `json:"num" binding:"required"`
	Name string   `json:"name" binding:"required"`
	Type []string `json:"type"`
}

type UpdatePokemonRequest struct {
	Num  string   `json:"num" binding:"required"`
	Name string   `json:"name" binding:"required"`
	Type []string `json:"type"`
}
