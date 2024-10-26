package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	models_character "github.com/JaanJah/auto-mwi-combat-simulator/models/character"
)

func main() {
	characters := readCharacters("input")
	fmt.Printf("Amount of characters to simulate: %d", len(characters))
}

func readCharacters(directory string) []models_character.Character {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var characters []models_character.Character

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(directory, file.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				continue
			}
			var character models_character.Character

			unmarshal_err := json.Unmarshal([]byte(data), &character)
			if unmarshal_err != nil {
				log.Printf("Failed to parse file %s: %v", filePath, unmarshal_err)
			}

			characters = append(characters, character)
		}
	}

	return characters
}
