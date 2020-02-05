package generate

import (
	"fmt"
	"log"

	"github.com/knightjdr/genemap/pkg/download"
	"github.com/knightjdr/genemap/pkg/zip"
)

func fetchUniprot(folder string) {
	downloadUniprot(folder)
	unzipUniprot(folder)
}

func downloadUniprot(folder string) {
	url := "ftp.uniprot.org:21"
	source := "/pub/databases/uniprot/current_release/knowledgebase/taxonomic_divisions/uniprot_sprot_human.dat.gz"
	target := fmt.Sprintf("%s/uniprot.dat.gz", folder)

	err := download.FTP(url, source, target)
	if err != nil {
		log.Fatalln(err)
	}
}

func unzipUniprot(folder string) {
	source := fmt.Sprintf("%s/uniprot.dat.gz", folder)
	target := fmt.Sprintf("%s/uniprot.dat", folder)

	zip.Gunzip(source, target)
}
