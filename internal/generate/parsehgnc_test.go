package generate

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/genemap/pkg/fs"
)

var hgncText = `{
	"responseHeader": {
		"status": 0,
		"QTime": 20
	},
	"response": {
		"numFound":42148,
		"start":0,
		"docs": [
			{
				"hgnc_id": "HGNC:5",
				"symbol": "A1BG",
				"name": "alpha-1-B glycoprotein",
				"entrez_id": "1",
				"refseq_accession": ["NM_130786"],
				"ensembl_gene_id": "ENSG00000121410",
				"uniprot_ids": ["P04217"]
			},
			{
				"hgnc_id": "HGNC:37133",
				"symbol": "A1BG-AS1",
				"name": "A1BG antisense RNA 1",
				"prev_symbol": ["NCRNA00181","A1BGAS","A1BG-AS"],
				"prev_name": [
					"non-protein coding RNA 181",
					"A1BG antisense RNA (non-protein coding)",
					"A1BG antisense RNA 1 (non-protein coding)"
				],
				"alias_symbol": ["FLJ23569"],
				"entrez_id": "503538",
				"refseq_accession": ["NR_015380"],
				"ensembl_gene_id": "ENSG00000268895"
			}
		]
	}
}`

var _ = Describe("Read hgnc file", func() {
	It("should read json file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/hgnc.json", []byte(hgncText), 0444)

		expected := hgncRecords{
			hgncRecord{
				Accession:   []string{"P04217"},
				EnsemblGene: "ENSG00000121410",
				Entrez:      "1",
				HGNC:        "5",
				Name:        "alpha-1-B glycoprotein",
				Refseq:      []string{"NM_130786"},
				Symbol:      "A1BG",
			},
			hgncRecord{
				AliasSymbol: []string{"FLJ23569"},
				EnsemblGene: "ENSG00000268895",
				Entrez:      "503538",
				HGNC:        "37133",
				Name:        "A1BG antisense RNA 1",
				PrevName: []string{
					"non-protein coding RNA 181",
					"A1BG antisense RNA (non-protein coding)",
					"A1BG antisense RNA 1 (non-protein coding)",
				},
				PrevSymbol: []string{"NCRNA00181", "A1BGAS", "A1BG-AS"},
				Refseq:     []string{"NR_015380"},
				Symbol:     "A1BG-AS1",
			},
		}
		Expect(parseHGNC("test")).To(Equal(expected))
	})
})
