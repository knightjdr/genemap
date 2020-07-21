package generate

func mergeData(uniprot *uniprotRecords, hgnc hgncRecords) *uniprotRecords {
	entries := &uniprotRecords{}
	hgncMap := mapHGNC(hgnc)
	
	for _, record := range *uniprot {
		entry := record
		if entry.Entrez == "" {
			if mappedValue, ok := hgncMap[entry.HGNC]; ok {
				entry.Entrez = mappedValue["entrez"]
			}
		}
		if len(entry.Symbol) == 0 {
			if mappedValue, ok := hgncMap[entry.HGNC]; ok {
				entry.Symbol = []string{mappedValue["symbol"]}
			}
		}
		*entries = append(*entries, entry)
	}
	
	return entries
}

func mapHGNC(hgnc hgncRecords) map[string]map[string]string {
	hgncMap := make(map[string]map[string]string, 0)

	for _, record := range hgnc {
		hgncMap[record.HGNC] = map[string]string{
			"entrez": record.Entrez,
			"symbol": record.Symbol,
		}
	}

	return hgncMap
}