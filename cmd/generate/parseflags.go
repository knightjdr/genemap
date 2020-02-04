package generate

import (
	"encoding/json"

	"github.com/knightjdr/genemap/internal/pkg/flags"
	"github.com/knightjdr/genemap/internal/pkg/fs"
	"github.com/spf13/afero"
)

func parseFlags() (map[string]interface{}, error) {
	args := flags.Parse()
	optionsFile := flags.SetString("options", args, map[string]interface{}{}, "")

	var err error
	var options map[string]interface{}

	// Read options from file if specified.
	if optionsFile != "" {
		jsonFile, _ := afero.ReadFile(fs.Instance, optionsFile)
		var jsonData interface{}
		err = json.Unmarshal(jsonFile, &jsonData)
		options, _ = jsonData.(map[string]interface{})
	}

	return options, err
}
