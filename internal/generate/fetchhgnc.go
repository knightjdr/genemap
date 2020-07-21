package generate

import (
	"fmt"
	"log"

	"github.com/knightjdr/genemap/pkg/download"
)

// HTTP download function
var HTTP = download.HTTP

func fetchHGNC(folder string) {
	url := "http://rest.genenames.org/fetch/status/Approved"
	target := fmt.Sprintf("%s/hgnc.json", folder)

	headers := map[string]string{
		"Accept": "application/json",
	}

	err := HTTP(url, headers, target)
	if err != nil {
		log.Fatalln(err)
	}
}
