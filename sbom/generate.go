package sbom

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func GenerateSbom() {
	// Set the number of CPU cores to use
	// runtime.GOMAXPROCS(runtime.NumCPU())

	targets := []string{"nginx", "harness/gitops-agent:v0.49.0", "harness/ssca-plugin", "uday4vunnam/go-ping:0.0.20", "node:20.11-alpine", "ritesports/fastify-api:dev"}
	outputFormats := []string{"syft-json", "spdx-json", "cyclonedx-json"}

	var wg sync.WaitGroup

	for _, target := range targets {
		targetDir := filepath.Join("temp", target)
		// os.MkdirAll(targetDir, 0755)

		for _, outputFormat := range outputFormats {
			cmd := getCommand(target, outputFormat, targetDir)
			cmdAllLayers := getAllLayersScopeCommand(target, outputFormat, targetDir)

			wg.Add(1)
			go runCommand(cmd, &wg)
			go runCommand(cmdAllLayers, &wg)
		}
	}

	wg.Wait()
	fmt.Println("All commands executed successfully.")
}

type Artifact struct {
	Locations []Location `json:"locations"`
	Type      string     `json:"type"`
	Purl      string     `json:"purl"`
	Name      string     `json:"name"`
	Version   string     `json:"version"`
}

type Location struct {
	Path string `json:"name"`
}

type ImageSBOM struct {
	Artifacts []Artifact `json:"artifacts"`
}

func runCommand(command string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %s - %v\n", command, err)
	}
}

func getCommand(target, outputFormat, outputDir string) string {
	return fmt.Sprintf("syft %s -o %s=%s/%s.json", target, outputFormat, outputDir, outputFormat)
}

func getAllLayersScopeCommand(target, outputFormat, outputDir string) string {
	return fmt.Sprintf("syft %s --scope all-layers -o %s=%s/%s.json", target, outputFormat, outputDir, "all-layers-"+outputFormat)
}
