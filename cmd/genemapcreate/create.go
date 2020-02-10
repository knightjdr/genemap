// Package main generates mapping file between various types.
package main

import (
	"log"

	"github.com/knightjdr/genemap/internal/generate"
)

func main() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	generate.MappingFilesCMD(options)
}
