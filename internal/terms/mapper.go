// Package terms maps gene identifiers between types.
package terms

import "errors"

// Mapper structure for setting options for mapping and retrieving results.
type Mapper struct {
	Converted          map[string]string
	FromType           string
	PossibleConverions map[string][]string
	records            records
	ToType             string
	Unconverted        []string
}

// CreateMapper creates a structure for mapping between geen identifiers.
func CreateMapper() *Mapper {
	return &Mapper{}
}

// Load mapping file in JSON format.
func (m *Mapper) Load(file string) {
	m.records = readMappingFile(file)
}

// Convert gene identifiers from Map.FromType to Map.ToType. Types must be one of (case insensitive)
// accession (uniprot, eg P31946), biogrid, ensemblgene, ensemblprotein,
// entrez (gene ID), hgnc, id (uniprot, eg 1433B_HUMAN), refseqmrna, refseqprotein or symbol.
func (m *Mapper) Convert(ids []string) error {
	if m.records == nil || len(m.records) == 0 {
		return errors.New("No mapping file loaded")
	}

	err := validateTypes(m)
	if err != nil {
		return err
	}

	convertIDs(m, ids)

	return nil
}
