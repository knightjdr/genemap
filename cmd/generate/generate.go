// Package generate mapping file between various types.
package generate

func generate() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	generate.MappingFiles(options)
}
