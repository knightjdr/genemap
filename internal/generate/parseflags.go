package generate

import (
	"errors"
	"strings"

	"github.com/knightjdr/generate/internal/pkg/flags"
)

type parameters struct {
	folder
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	folder := flags.SetString("folder", args, fileOptions, "./")

	// Copy arguments from options file.
	options := parameters{
		folder: folder,
	}

	return options, err
}
