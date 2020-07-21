// Package generate mapping file between various types.
package generate

import (
	"log"
)

// MappingFiles generates file for mapping between types.
func MappingFiles(folder string) {
	fetchUniprot(folder)
	fetchHGNC(folder)
	entries := parseUniprot(folder)
	hgnc := parseHGNC(folder)
	entries = mergeData(entries, hgnc)
	outputMapping(entries, folder)
}

// MappingFilesCMD is a wrapper for the CLI.
func MappingFilesCMD(options map[string]interface{}) {
	settings, err := parseFlags(options)
	if err != nil {
		log.Fatalln(err)
	}

	MappingFiles(settings.folder)
}
