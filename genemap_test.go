package genemap_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	. "github.com/knightjdr/genemap"
	"github.com/knightjdr/genemap/internal/generate"
	"github.com/knightjdr/genemap/internal/terms"
	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/knightjdr/genemap/pkg/zip"
)

var jsonText = `[
	{
		"Accession": ["Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"], 
		"Symbol": ["ACCS", "PHACS"]
	},
	{
		"Accession": ["Q9BUL8", "A8K515", "D3DNN5", "O14811"], 
		"Symbol": ["PDCD10", "CCM3", "TFAR15"]
	},
	{
		"Accession": ["Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"], 
		"Symbol": ["STK26", "MASK", "MST4"]
	},
	{
		"Accession": ["P31946", "A8K9K2", "E1P616"], 
		"Symbol": ["YWHAB"]
	}
]`

var uniprotText = `
ID   1433B_HUMAN             Reviewed;         246 AA.
AC   P31946; A8K9K2; E1P616;
DT   01-JUL-1993, integrated into UniProtKB/Swiss-Prot.
DT   23-JAN-2007, sequence version 3.
DT   11-DEC-2019, entry version 225.
DE   RecName: Full=14-3-3 protein beta/alpha;
GN   Name=YWHAB;
DR   Ensembl; ENST00000353703; ENSP00000300161; ENSG00000166913. [P31946-1]
DR   Ensembl; ENST00000372839; ENSP00000361930; ENSG00000166913. [P31946-1]
DR   BioGrid; 113361; 404.
DR   GeneID; 7529; -.
DR   RefSeq; NP_003395.1; NM_003404.4. [P31946-1]
DR   RefSeq; NP_647539.1; NM_139323.3. [P31946-1]
DR   RefSeq; XP_016883528.1; XM_017028039.1. [P31946-1]
DR   HGNC; HGNC:12849; YWHAB.
//
ID   1A1L1_HUMAN             Reviewed;         501 AA.
AC   Q96QU6; B4E219; Q8WUL4; Q96LX5;
DT   05-FEB-2008, integrated into UniProtKB/Swiss-Prot.
DT   01-DEC-2001, sequence version 1.
DT   11-DEC-2019, entry version 133.
DE   RecName: Full=1-aminocyclopropane-1-carboxylate synthase-like protein 1;
DE            Short=ACC synthase-like protein 1;
GN   Name=ACCS; Synonyms=PHACS;
DR   Ensembl; ENST00000263776; ENSP00000263776; ENSG00000110455. [Q96QU6-1]
DR   GeneID; 84680; -.
DR   RefSeq; NP_001120691.1; NM_001127219.1. [Q96QU6-1]
DR   RefSeq; NP_115981.1; NM_032592.3. [Q96QU6-1]
DR   HGNC; HGNC:23989; ACCS.
//
`

var _ = Describe("Generate map", func() {
	It("should load map and convert id between types", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		oldFTP := generate.FTP
		defer func() { generate.FTP = oldFTP }()
		generate.FTP = func(url, sourceFile, targetFile string) error {
			zip.Gzip(uniprotText, "test/uniprot.dat.gz")
			return nil
		}

		expected := "[\n" +
			"\t{\n" +
			"\t\t\"Accession\": [\n\t\t\t\"P31946\",\n\t\t\t\"A8K9K2\",\n\t\t\t\"E1P616\"\n\t\t],\n" +
			"\t\t\"Biogrid\": \"113361\",\n" +
			"\t\t\"EnsemblGene\": [\n\t\t\t\"ENSG00000166913\"\n\t\t],\n" +
			"\t\t\"EnsemblProtein\": [\n\t\t\t\"ENSP00000300161\",\n\t\t\t\"ENSP00000361930\"\n\t\t],\n" +
			"\t\t\"Entrez\": \"7529\",\n" +
			"\t\t\"HGNC\": \"12849\",\n" +
			"\t\t\"ID\": \"1433B_HUMAN\",\n" +
			"\t\t\"Name\": \"14-3-3 protein beta/alpha\",\n" +
			"\t\t\"RefseqMRNA\": [\n\t\t\t\"NM_003404\",\n\t\t\t\"NM_139323\",\n\t\t\t\"XM_017028039\"\n\t\t],\n" +
			"\t\t\"RefseqProtein\": [\n\t\t\t\"NP_003395\",\n\t\t\t\"NP_647539\",\n\t\t\t\"XP_016883528\"\n\t\t],\n" +
			"\t\t\"Symbol\": [\n\t\t\t\"YWHAB\"\n\t\t]\n" +
			"\t},\n" +
			"\t{\n" +
			"\t\t\"Accession\": [\n\t\t\t\"Q96QU6\",\n\t\t\t\"B4E219\",\n\t\t\t\"Q8WUL4\",\n\t\t\t\"Q96LX5\"\n\t\t],\n" +
			"\t\t\"Biogrid\": \"\",\n" +
			"\t\t\"EnsemblGene\": [\n\t\t\t\"ENSG00000110455\"\n\t\t],\n" +
			"\t\t\"EnsemblProtein\": [\n\t\t\t\"ENSP00000263776\"\n\t\t],\n" +
			"\t\t\"Entrez\": \"84680\",\n" +
			"\t\t\"HGNC\": \"23989\",\n" +
			"\t\t\"ID\": \"1A1L1_HUMAN\",\n" +
			"\t\t\"Name\": \"1-aminocyclopropane-1-carboxylate synthase-like protein 1\",\n" +
			"\t\t\"RefseqMRNA\": [\n\t\t\t\"NM_001127219\",\n\t\t\t\"NM_032592\"\n\t\t],\n" +
			"\t\t\"RefseqProtein\": [\n\t\t\t\"NP_001120691\",\n\t\t\t\"NP_115981\"\n\t\t],\n" +
			"\t\t\"Symbol\": [\n\t\t\t\"ACCS\",\n\t\t\t\"PHACS\"\n\t\t]\n" +
			"\t}\n" +
			"]"

		Generate("test")

		outfile := fmt.Sprintf("test/genemap-%s.json", time.Now().Format("2006-01-02"))
		bytes, _ := afero.ReadFile(fs.Instance, outfile)
		Expect(string(bytes)).To(Equal(expected))

	})
})

var _ = Describe("Create mapping struct", func() {
	It("should create struct", func() {
		expected := &terms.Mapper{}

		Expect(CreateMapper()).To(Equal(expected))
	})
})

var _ = Describe("Convert IDs", func() {
	It("should load map and convert id between types", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/genemap.json", []byte(jsonText), 0444)

		ids := []string{"ACCS", "CCM3", "MST4", "STK24", "YWHAB"}

		expectedConverted := map[string]string{
			"ACCS":  "Q96QU6",
			"CCM3":  "Q9BUL8",
			"MST4":  "Q9P289",
			"YWHAB": "P31946",
		}
		expectedPossibleConversions := map[string][]string{
			"ACCS":  []string{"Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"},
			"CCM3":  []string{"Q9BUL8", "A8K515", "D3DNN5", "O14811"},
			"MST4":  []string{"Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"},
			"YWHAB": []string{"P31946", "A8K9K2", "E1P616"},
		}
		expectedUnconverted := []string{"STK24"}

		mapper := CreateMapper()
		mapper.Load("test/genemap.json")
		mapper.FromType = "Symbol"
		mapper.ToType = "Accession"
		err := mapper.Convert(ids)

		Expect(err).To(BeNil(), "should not return an error: %e", err)
		Expect(mapper.Converted).To(Equal(expectedConverted), "should convert ids to a single value")
		Expect(mapper.PossibleConversions).To(Equal(expectedPossibleConversions), "should return all possible conversion for ids")
		Expect(mapper.Unconverted).To(Equal(expectedUnconverted), "should return ids that could not be converted")
	})
})
