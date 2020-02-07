package terms

import "sort"

func convertIDs(m *Mapper, ids []string) {
	m.Converted = make(map[string]string, 0)
	m.PossibleConverions = make(map[string][]string, 0)

	lookup := createLookupTable(m)
	lookupAndConvertIDs(m, lookup, ids)
}

func createLookupTable(m *Mapper) *map[string]record {
	lookup := &map[string]record{}

	for _, mapRecord := range m.records {
		for _, accession := range mapRecord.getValue(m.FromType) {
			(*lookup)[accession] = mapRecord
		}
	}

	return lookup
}

func lookupAndConvertIDs(m *Mapper, lookup *map[string]record, ids []string) {
	for _, id := range ids {
		if _, ok := (*lookup)[id]; !ok {
			m.Unconverted = append(m.Unconverted, id)
		} else {
			targetValues := (*lookup)[id].getValue(m.ToType)
			m.PossibleConverions[id] = targetValues
			m.Converted[id] = targetValues[0]

		}
	}

	sort.Strings(m.Unconverted)
}
