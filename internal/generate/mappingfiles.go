// Package generate mapping file between various types.
package generate

import (
	"log"
)

// MappingFiles generates file for mapping between types.
func MappingFiles(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	fetchUniprot(options.folder)
	entries := parseUniprot(options.folder)
	outputMapping(entries, options.folder)
}
