// Package main generates mapping file between various types.
package main

import (
	"log"

	g "github.com/knightjdr/genemap/internal/generate"
)

func main() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	g.MappingFiles(options)
}
