package main

import (
	"log"
	"tft-tracker/internal/riotapi"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	riotapi.GetSummonerByName("Mortdog", "na1")
}
