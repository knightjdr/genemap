package terms

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/genemap/pkg/fs"
)

func readIDFile(filename string) []string {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	ids := make([]string, 0)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		ids = append(ids, line[0])
	}

	return ids
}
