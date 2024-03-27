package sbom

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func Nomalize() {
	targets := []string{"nginx", "harness/gitops-agent:v0.49.0", "harness/ssca-plugin", "uday4vunnam/go-ping:0.0.20", "node:20.11-alpine", "ritesports/fastify-api:dev"}

	var wg sync.WaitGroup

	for _, target := range targets {
		source := filepath.Join("temp", target, "all-layers-syft-json.json")
		target := filepath.Join("temp", target, "all-layers-syft-json-transformed.json")

		wg.Add(1)
		go transformJSON(&wg, source, target)

	}

	wg.Wait()
	fmt.Println("All commands executed successfully.")
}

func transformJSON(wg *sync.WaitGroup, importJSON, exportJSON string) error {
	defer wg.Done()

	// Open and decode the JSON file
	inputFile, err := os.Open(importJSON)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var imageSBOM ImageSBOM
	err = json.NewDecoder(inputFile).Decode(&imageSBOM)
	if err != nil {
		return err
	}

	// Filter components based on type and property
	filteredComponents := make([]Artifact, 0)
	for _, comp := range imageSBOM.Artifacts {
		if comp.Type != "operating-system" && (comp.Locations == nil || len(comp.Locations) < 2) {
			comp.Locations = nil
			filteredComponents = append(filteredComponents, comp)
		}
	}

	// Update imageSBOM with filtered components
	imageSBOM.Artifacts = filteredComponents

	// Write the output JSON
	outputFile, err := os.Create(exportJSON)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(imageSBOM)
	if err != nil {
		return err
	}

	return nil
}
