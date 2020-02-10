// Package main reads gene identifiers from a file and converts between requested types.
package main

import (
	"log"

	"github.com/knightjdr/genemap/internal/terms"
)

func main() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	terms.MapperCMD(options)
}
