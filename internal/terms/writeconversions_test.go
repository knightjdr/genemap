package terms

import (
	"github.com/knightjdr/genemap/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write conversion file", func() {
	It("should write file with conversions", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		mapper := &Mapper{
			Converted: map[string]string{
				"ACCS":  "Q96QU6",
				"CCM3":  "Q9BUL8",
				"MST4":  "Q9P289",
				"YWHAB": "P31946",
			},
			PossibleConversions: map[string][]string{
				"ACCS":  []string{"Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"},
				"CCM3":  []string{"Q9BUL8", "A8K515", "D3DNN5", "O14811"},
				"MST4":  []string{"Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"},
				"YWHAB": []string{"P31946", "A8K9K2", "E1P616"},
			},
			Unconverted: []string{"STK24"},
		}
		settings := parameters{
			fromType: "Symbol",
			outFile:  "test/out.txt",
			toType:   "Accession",
		}

		expected := "Symbol\tAccession\tPossible conversions\n" +
			"ACCS\tQ96QU6\tQ96QU6, B4E219, Q8WUL4, Q96LX5\n" +
			"CCM3\tQ9BUL8\tQ9BUL8, A8K515, D3DNN5, O14811\n" +
			"MST4\tQ9P289\tQ9P289, B2RAU2, Q3ZB77, Q8NC04, Q9BXC3, Q9BXC4\n" +
			"YWHAB\tP31946\tP31946, A8K9K2, E1P616\n" +
			"STK24\t\t\n"

		writeConversions(mapper, settings)
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
