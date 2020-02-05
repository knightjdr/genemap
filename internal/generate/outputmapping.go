package generate

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/spf13/afero"
)

func outputMapping(entries *uniprotEntries, folder string) {
	outfile := fmt.Sprintf("%s/genemap-%s.json", folder, time.Now().Format("2006-01-02"))

	mappingJSON, _ := json.MarshalIndent(entries, "", "\t")
	err := afero.WriteFile(fs.Instance, outfile, mappingJSON, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
