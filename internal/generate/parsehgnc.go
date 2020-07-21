package generate

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/spf13/afero"
)

type jsonData struct {
	Response response
}

type response struct {
	NumFound int
	Docs     hgncRecords
}

type hgncRecords []hgncRecord

type hgncRecord struct {
	Accession   []string `json:"uniprot_ids,omitempty"`
	AliasSymbol []string `json:"alias_symbol,omitempty"`
	EnsemblGene string   `json:"ensembl_gene_id,omitempty"`
	Entrez      string   `json:"entrez_id,omitempty"`
	HGNC        string   `json:"hgnc_id,omitempty"`
	Name        string   `json:"name,omitempty"`
	PrevName    []string `json:"prev_name,omitempty"`
	PrevSymbol  []string `json:"prev_symbol,omitempty"`
	Refseq      []string `json:"refseq_accession,omitempty"`
	Symbol      string   `json:"symbol,omitempty"`
}

func (h *hgncRecords) formatRecords() {
	for i, record := range *h {
		(*h)[i].HGNC = strings.Split(record.HGNC, ":")[1]
	}
}

func parseHGNC(folder string) hgncRecords {
	hgncFile := fmt.Sprintf("%s/hgnc.json", folder)
	byteValue, err := afero.ReadFile(fs.Instance, hgncFile)
	if err != nil {
		log.Fatalln(err)
	}
	var data jsonData
	json.Unmarshal(byteValue, &data)

	data.Response.Docs.formatRecords()

	removeJSON(hgncFile)

	return data.Response.Docs
}

func removeJSON(file string) {
	err := fs.Instance.Remove(file)
	if err != nil {
		log.Println(err)
	}
}
