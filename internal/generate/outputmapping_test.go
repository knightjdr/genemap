package generate

import (
	"fmt"
	"time"

	"github.com/knightjdr/genemap/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Output mapping", func() {
	It("should parse entries from uniprot file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		entries := &uniprotEntries{
			uniprotEntry{
				Accession:       []string{"P31946", "A8K9K2", "E1P616"},
				Biogrid:         113361,
				EnsemblGene:     []string{"ENSG00000166913"},
				EnsembleProtein: []string{"ENSP00000300161", "ENSP00000361930"},
				Entrez:          7529,
				HGNC:            12849,
				ID:              "1433B_HUMAN",
				Name:            "14-3-3 protein beta/alpha",
				RefseqMRNA:      []string{"NM_003404", "NM_139323", "XM_017028039"},
				RefseqProtein:   []string{"NP_003395", "NP_647539", "XP_016883528"},
				Reviewed:        true,
				Symbol:          []string{"YWHAB"},
			},
		}

		expected := "[\n" +
			"\t{\n" +
			"\t\t\"accession\": [\n\t\t\t\"P31946\",\n\t\t\t\"A8K9K2\",\n\t\t\t\"E1P616\"\n\t\t],\n" +
			"\t\t\"biogrid\": 113361,\n" +
			"\t\t\"ensemblg\": [\n\t\t\t\"ENSG00000166913\"\n\t\t],\n" +
			"\t\t\"ensemblp\": [\n\t\t\t\"ENSP00000300161\",\n\t\t\t\"ENSP00000361930\"\n\t\t],\n" +
			"\t\t\"entrez\": 7529,\n" +
			"\t\t\"hgnc\": 12849,\n" +
			"\t\t\"id\": \"1433B_HUMAN\",\n" +
			"\t\t\"name\": \"14-3-3 protein beta/alpha\",\n" +
			"\t\t\"refseqm\": [\n\t\t\t\"NM_003404\",\n\t\t\t\"NM_139323\",\n\t\t\t\"XM_017028039\"\n\t\t],\n" +
			"\t\t\"refseqp\": [\n\t\t\t\"NP_003395\",\n\t\t\t\"NP_647539\",\n\t\t\t\"XP_016883528\"\n\t\t],\n" +
			"\t\t\"symbol\": [\n\t\t\t\"YWHAB\"\n\t\t]\n" +
			"\t}\n" +
			"]"

		outputMapping(entries, "test")

		date := time.Now().Format("2006-01-02")
		bytes, _ := afero.ReadFile(fs.Instance, fmt.Sprintf("test/genemap-%s.json", date))
		Expect(string(bytes)).To(Equal(expected))
	})
})
