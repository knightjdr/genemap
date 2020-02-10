package terms

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert ids", func() {
	It("should convert ids", func() {
		ids := []string{"ACCS", "CCM3", "MST4", "STK24", "YWHAB"}
		idRecords := records{
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
		m := &Mapper{
			FromType: "Symbol",
			ToType:   "Accession",
			records:  idRecords,
		}

		expected := &Mapper{
			Converted: map[string]string{
				"ACCS":  "Q96QU6",
				"CCM3":  "Q9BUL8",
				"MST4":  "Q9P289",
				"YWHAB": "P31946",
			},
			FromType: "Symbol",
			PossibleConversions: map[string][]string{
				"ACCS":  []string{"Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"},
				"CCM3":  []string{"Q9BUL8", "A8K515", "D3DNN5", "O14811"},
				"MST4":  []string{"Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"},
				"YWHAB": []string{"P31946", "A8K9K2", "E1P616"},
			},
			records:     idRecords,
			ToType:      "Accession",
			Unconverted: []string{"STK24"},
		}

		convertIDs(m, ids)
		Expect(m).To(Equal(expected))
	})
})
