package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	inventories := readInventories("input")
	fmt.Println(len(inventories))
}

func readInventories(directory string) []string {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var inventories []string

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(directory, file.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				continue
			}
			inventories = append(inventories, string(data))
		}
	}

	return inventories
}
