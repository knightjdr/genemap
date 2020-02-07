// Package generate mapping file between various types.
package generate

import (
	"log"
)

// MappingFiles generates file for mapping between types.
func MappingFiles(options map[string]interface{}) {
	settings, err := parseFlags(options)
	if err != nil {
		log.Fatalln(err)
	}

	fetchUniprot(settings.folder)
	entries := parseUniprot(settings.folder)
	outputMapping(entries, settings.folder)
}
