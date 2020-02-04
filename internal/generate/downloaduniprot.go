package generate

import (
	"fmt"

	"github.com/knightjdr/genemap/internal/pkg/download"
)

func downloadUniprot(folder string) {
	url := "ftp://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/taxonomic_divisions/uniprot_sprot_human.dat.gz"
	download.File(url, fmt.Sprintf("%s/uniprot.dat.gz", folder))
}
