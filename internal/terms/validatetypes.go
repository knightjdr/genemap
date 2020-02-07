package terms

import (
	"errors"
	"fmt"
	"strings"
)

func validateTypes(m *Mapper) error {
	err := confirmTypesSpecified(m)
	if err != nil {
		return err
	}

	err = confirmValidType(m)
	if err != nil {
		return err
	}

	return nil
}

func confirmTypesSpecified(m *Mapper) error {
	if m.FromType == "" {
		return errors.New("Must specify FromType identifier for conversion")
	}
	if m.ToType == "" {
		return errors.New("Must specify ToType identifier for conversion")
	}

	return nil
}

func confirmValidType(m *Mapper) error {
	acceptedTypes := map[string]string{
		"accession":   "Accession",
		"biogrid":     "Biogrid",
		"ensemblgene": "EnsemblGene",
		"ensemblp":    "EnsemblProtein",
		"entrez":      "Entrez",
		"hgnc":        "HGNC",
		"id":          "ID",
		"refseqm":     "RefseqMRNA",
		"refseqp":     "RefseqProtein",
		"symbol":      "Symbol",
	}

	m.FromType = strings.ToLower(m.FromType)
	m.ToType = strings.ToLower(m.ToType)

	if _, ok := acceptedTypes[strings.ToLower(m.FromType)]; !ok {
		return fmt.Errorf("FromType identifier (\"%s\") is invalid", m.FromType)
	}
	if _, ok := acceptedTypes[strings.ToLower(m.ToType)]; !ok {
		return fmt.Errorf("ToType identifier (\"%s\") is invalid", m.ToType)
	}

	m.FromType = acceptedTypes[m.FromType]
	m.ToType = acceptedTypes[m.ToType]

	if m.FromType == m.ToType {
		return fmt.Errorf("Conversion types are identifical (\"%s\")", m.FromType)
	}

	return nil
}
