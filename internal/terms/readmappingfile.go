package terms

import (
	"encoding/json"
	"log"

	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/spf13/afero"
)

type records []record

type record struct {
	Accession      []string
	Biogrid        string
	EnsemblGene    []string
	EnsemblProtein []string
	Entrez         string
	HGNC           string
	ID             string
	Name           string
	RefseqMRNA     []string
	RefseqProtein  []string
	Symbol         []string
}

func (r record) getValue(field string) []string {
	if field == "Accession" {
		return r.Accession
	}
	if field == "Biogrid" {
		return []string{r.Biogrid}
	}
	if field == "EnsemblGene" {
		return r.EnsemblGene
	}
	if field == "EnsemblProtein" {
		return r.EnsemblProtein
	}
	if field == "Entrez" {
		return []string{r.Entrez}
	}
	if field == "HGNC" {
		return []string{r.HGNC}
	}
	if field == "ID" {
		return []string{r.ID}
	}
	if field == "Name" {
		return []string{r.Name}
	}
	if field == "RefseqMRNA" {
		return r.RefseqMRNA
	}
	if field == "RefseqProtein" {
		return r.RefseqProtein
	}
	if field == "Symbol" {
		return r.Symbol
	}

	return []string{}
}

func readMappingFile(mappingFile string) records {
	byteValue, err := afero.ReadFile(fs.Instance, mappingFile)
	if err != nil {
		log.Fatalln(err)
	}
	var records records
	json.Unmarshal(byteValue, &records)

	return records
}
