package imagereader

import (
	"archive/tar"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadImage() {
	cmd := exec.Command("bash", "-c", "docker save nginx -o image.tar | gzip > image.tar.gz")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing docker save command:", err)
		return
	}
	fmt.Println(string(output))

	// tar xf image.tar
	file, err := os.Open("/Users/uday/Downloads/go-client/image.tar")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading tar header:", err)
			continue
		}

		if header.FileInfo().IsDir() {
			continue
		}

		fileName := header.Name
		if strings.HasSuffix(fileName, ".json") {
			file, err := os.Create(filepath.Base(fileName))
			if err != nil {
				fmt.Println("Error creating file:", err)
				continue
			}
			defer file.Close()

			_, err = io.Copy(file, tarReader)
			if err != nil {
				fmt.Println("Error writing file:", err)
			}
		}
	}

	// for file in *.json; do
	//     rm -f leak.json
	//     echo $file
	//     cat $file | jq -r '."moby.buildkit.buildinfo.v1"' | base64 -d > base_image.json
	// done
	files, err := filepath.Glob("*.json")
	if err != nil {
		fmt.Println("Error finding JSON files:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file)

		jsonData, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		var jsonObj map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonObj)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			continue
		}

		buildInfo, ok := jsonObj["moby.buildkit.buildinfo.v1"]
		if !ok {
			fmt.Println("'moby.buildkit.buildinfo.v1' not found")
			continue
		}

		buildInfoBytes, err := base64.StdEncoding.DecodeString(buildInfo.(string))
		if err != nil {
			fmt.Println("Error decoding base64:", err)
			continue
		}

		err = os.Remove("leak.json")
		if err != nil {
			fmt.Println("Error removing leak.json:", err)
		}

		err = os.WriteFile("base_image.json", buildInfoBytes, 0644)
		if err != nil {
			fmt.Println("Error writing base_image.json:", err)
		}
	}
}
