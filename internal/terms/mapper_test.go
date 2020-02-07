package terms

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/genemap/pkg/fs"
)

var _ = Describe("Mapper", func() {
	It("should load JSON file", func() {
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
		}

		mapper := CreateMapper()
		mapper.Load("test/genemap.json")
		Expect(mapper.records).To(Equal(expected))
	})

	It("should return error when map has not be loaded for conversion", func() {
		ids := []string{"ACCS", "CCM3", "MST4", "STK24", "YWHAB"}

		mapper := CreateMapper()
		mapper.FromType = "Symbol"
		mapper.ToType = "Accession"
		err := mapper.Convert(ids)

		Expect(err).To(MatchError("No mapping file loaded"))
	})

	It("should return error when the map has no records", func() {
		ids := []string{"ACCS", "CCM3", "MST4", "STK24", "YWHAB"}

		mapper := CreateMapper()
		mapper.FromType = "Symbol"
		mapper.ToType = "Accession"
		mapper.records = records{}
		err := mapper.Convert(ids)

		Expect(err).To(MatchError("No mapping file loaded"))
	})

	It("should convert IDs", func() {
		ids := []string{"ACCS", "CCM3", "MST4", "STK24", "YWHAB"}

		expectedConverted := map[string]string{
			"ACCS":  "Q96QU6",
			"CCM3":  "Q9BUL8",
			"MST4":  "Q9P289",
			"YWHAB": "P31946",
		}
		expectedPossibleConverions := map[string][]string{
			"ACCS":  []string{"Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"},
			"CCM3":  []string{"Q9BUL8", "A8K515", "D3DNN5", "O14811"},
			"MST4":  []string{"Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"},
			"YWHAB": []string{"P31946", "A8K9K2", "E1P616"},
		}
		expectedUnconverted := []string{"STK24"}

		mapper := CreateMapper()
		mapper.FromType = "Symbol"
		mapper.ToType = "Accession"
		mapper.records = records{
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
			record{
				Accession: []string{"P31946", "A8K9K2", "E1P616"},
				Symbol:    []string{"YWHAB"},
			},
		}
		err := mapper.Convert(ids)

		Expect(err).To(BeNil(), "should not return an error: %e", err)
		Expect(mapper.Converted).To(Equal(expectedConverted), "should convert ids to a single value")
		Expect(mapper.PossibleConverions).To(Equal(expectedPossibleConverions), "should return all possible conversion for ids")
		Expect(mapper.Unconverted).To(Equal(expectedUnconverted), "should return ids that could not be converted")
	})
})
