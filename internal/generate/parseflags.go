package generate

import (
	"github.com/knightjdr/genemap/internal/pkg/flags"
)

type parameters struct {
	folder string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	folder := flags.SetString("folder", args, fileOptions, ".")
	folder = parseFolderName(folder)

	// Copy arguments from options file.
	options := parameters{
		folder: folder,
	}

	var err error
	return options, err
}

func parseFolderName(folder string) string {
	if folder[len(folder)-1:] == "/" {
		return folder[:len(folder)-1]
	}

	return folder
}
