package terms

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/genemap/pkg/fs"
)

var jsonText = `[
	{
		"Accession": ["P31946","A8K9K2","E1P616"], 
		"Biogrid": "113361",
		"EnsemblGene": ["ENSG00000166913"],
		"EnsemblProtein": ["ENSP00000300161","ENSP00000361930"],
		"Entrez": "7529",
		"HGNC": "12849",
		"ID": "1433B_HUMAN",
		"Name": "14-3-3 protein beta/alpha",
		"RefseqMRNA": ["NM_003404","NM_139323","XM_017028039"],
		"RefseqProtein": ["NP_003395","NP_647539","XP_016883528"],
		"Symbol": ["YWHAB"]
	},
	{
		"Accession": ["Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"],
		"Symbol": ["ACCS", "PHACS"]
	},
	{
		"Accession": ["Q9BUL8", "A8K515", "D3DNN5", "O14811"],
		"Symbol":    ["PDCD10", "CCM3", "TFAR15"]
	},
	{
		"Accession": ["Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"],
		"Symbol":    ["STK26", "MASK", "MST4"]
	}
]`

var _ = Describe("Get record value", func() {
	It("should return values", func() {
		r := record{
			Accession:      []string{"P31946", "A8K9K2", "E1P616"},
			Biogrid:        "113361",
			EnsemblGene:    []string{"ENSG00000166913"},
			EnsemblProtein: []string{"ENSP00000300161", "ENSP00000361930"},
			Entrez:         "7529",
			HGNC:           "12849",
			ID:             "1433B_HUMAN",
			RefseqMRNA:     []string{"NM_003404", "NM_139323", "XM_017028039"},
			RefseqProtein:  []string{"NP_003395", "NP_647539", "XP_016883528"},
			Symbol:         []string{"YWHAB"},
		}

		tests := []string{"Accession", "Biogrid", "EnsemblGene", "EnsemblProtein", "Entrez", "HGNC", "ID", "RefseqMRNA", "RefseqProtein", "Symbol", ""}
		expected := [][]string{
			[]string{"P31946", "A8K9K2", "E1P616"},
			[]string{"113361"},
			[]string{"ENSG00000166913"},
			[]string{"ENSP00000300161", "ENSP00000361930"},
			[]string{"7529"},
			[]string{"12849"},
			[]string{"1433B_HUMAN"},
			[]string{"NM_003404", "NM_139323", "XM_017028039"},
			[]string{"NP_003395", "NP_647539", "XP_016883528"},
			[]string{"YWHAB"},
			[]string{},
		}

		for i, test := range tests {
			Expect(r.getValue(test)).To(Equal(expected[i]), fmt.Sprintf("should return %s = %v", test, expected[i]))
		}
	})
})

var _ = Describe("Read mapping file", func() {
	It("should read json file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/genemap.json", []byte(jsonText), 0444)

		expected := records{
			record{
				Accession:      []string{"P31946", "A8K9K2", "E1P616"},
				Biogrid:        "113361",
				EnsemblGene:    []string{"ENSG00000166913"},
				EnsemblProtein: []string{"ENSP00000300161", "ENSP00000361930"},
				Entrez:         "7529",
				HGNC:           "12849",
				ID:             "1433B_HUMAN",
				RefseqMRNA:     []string{"NM_003404", "NM_139323", "XM_017028039"},
				RefseqProtein:  []string{"NP_003395", "NP_647539", "XP_016883528"},
				Symbol:         []string{"YWHAB"},
			},
			record{
				Accession: []string{"Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"},
				Symbol:    []string{"ACCS", "PHACS"},
			},
			record{
				Accession: []string{"Q9BUL8", "A8K515", "D3DNN5", "O14811"},
				Symbol:    []string{"PDCD10", "CCM3", "TFAR15"},
			},
			record{
				Accession: []string{"Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"},
				Symbol:    []string{"STK26", "MASK", "MST4"},
			},
		}
		Expect(readMappingFile("test/genemap.json")).To(Equal(expected))
	})
})
