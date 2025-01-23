package main

import (
	"path/filepath"

	"C"
	"encoding/json"
	"fmt"
	"github.com/adrg/frontmatter"
	"os"
)

type metadata struct {
	Publish bool     `yaml:"publish"`
	Title   string   `yaml:"title"`
	Path    string   `yaml:"-"` /* manually set this field */
	Tags    []string `yaml:"tags"`
}

func getFile(path string) (*os.File, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	output, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func processMetadata1(path string) (metadata, error) {
	meta := metadata{}

	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Print("Error while resolving filepath: ", path)
		return meta, err
	}

	file, err := getFile(path)
	if err != nil {
		fmt.Print("Error while retrieving file handle: ", err.Error())
		return meta, err
	}
	defer file.Close()

	/* fill meta with frontmatter */
	_, err = frontmatter.Parse(file, &meta)
	// handleFrontmatterParseError(path, err)

	/* manually insert filepath */
	meta.Path = path

	/* convert to json */

	return meta, nil
}

//export ProcessFrontmatter
func ProcessFrontmatter(path string) *C.char {
	a, _ := processMetadata1(path)
	jso, _ := json.Marshal(a)
	return C.CString(string(jso))
}

func main() {}
