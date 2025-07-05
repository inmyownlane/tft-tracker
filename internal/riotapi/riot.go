package riotapi

import (
	"fmt"
	"os"
)

func GetSummonerByName(name string, region string) {
	apiKey := os.Getenv("RIOT_API_KEY")

	if apiKey == "" {
		fmt.Println("No API key found")
		return
	}

	fmt.Println("Name:", name)
	fmt.Println("Region:", region)
	fmt.Println("API Key:", apiKey)
}
