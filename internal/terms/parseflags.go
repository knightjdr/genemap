package terms

import (
	"errors"
	"strings"

	"github.com/knightjdr/genemap/pkg/flags"
)

type parameters struct {
	fromType string
	idFile   string
	mapFile  string
	outFile  string
	toType   string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	fromType := flags.SetString("fromType", args, fileOptions, "")
	idFile := flags.SetString("idFile", args, fileOptions, "")
	mapFile := flags.SetString("mapFile", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "conversion.txt")
	toType := flags.SetString("toType", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		fromType: fromType,
		idFile:   idFile,
		mapFile:  mapFile,
		outFile:  outFile,
		toType:   toType,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.fromType == "" {
		messages = append(messages, "missing fromType conversion identifier")
	}
	if options.idFile == "" {
		messages = append(messages, "missing text file with gene identifiers for conversion")
	}
	if options.mapFile == "" {
		messages = append(messages, "missing gene mapping file")
	}
	if options.toType == "" {
		messages = append(messages, "missing toType conversion identifier")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
