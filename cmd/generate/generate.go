// Package generate mapping file between various types.
package generate

import (
	"log"

	g "github.com/knightjdr/genemap/internal/generate"
)

func generate() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	g.MappingFiles(options)
}
