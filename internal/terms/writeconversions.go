package terms

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/spf13/afero"
)

func writeConversions(mapper *Mapper, settings parameters) {
	var buffer bytes.Buffer

	writeHeader(&buffer, settings)
	writeBody(&buffer, mapper)

	afero.WriteFile(fs.Instance, settings.outFile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer, settings parameters) {
	header := fmt.Sprintf("%s\t%s\tPossible conversions\n", settings.fromType, settings.toType)
	buffer.WriteString(header)
}

func writeBody(buffer *bytes.Buffer, mapper *Mapper) {
	outputOrder := getOutputOrder(mapper)

	for _, id := range outputOrder {
		target, targets := getMappingTargets(id, mapper)
		buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\n", id, target, targets))
	}
}

func getOutputOrder(mapper *Mapper) []string {
	ids := make([]string, 0)

	for id := range mapper.Converted {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	return append(ids, mapper.Unconverted...)
}

func getMappingTargets(id string, mapper *Mapper) (string, string) {
	if _, ok := mapper.Converted[id]; ok {
		return mapper.Converted[id], strings.Join(mapper.PossibleConversions[id], ", ")
	}
	return "", ""
}
