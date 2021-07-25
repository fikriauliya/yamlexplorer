package main

import (
	"fmt"

	"github.com/fikriauliya/yamlexplorer/parser"
	"github.com/fikriauliya/yamlexplorer/renderer"
)

func main() {
	path := "sample/index.yaml"
	table, err := parser.ParseYAML(path)
	if err != nil {
		fmt.Printf("Error parsing file: %s: %s\n", path, err)
		return
	}
	fmt.Printf("%+v", table)
	r := renderer.NewRenderer()
	r.Render(table)
}
