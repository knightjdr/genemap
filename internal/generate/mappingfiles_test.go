package generate

import (
	"fmt"
	"time"

	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/knightjdr/genemap/pkg/zip"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Mapping files", func() {
	It("should fetch file from uniprot", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		oldFTP := FTP
		defer func() { FTP = oldFTP }()
		FTP = func(url, sourceFile, targetFile string) error {
			zip.Gzip(uniprotText, "test/uniprot.dat.gz")
			return nil
		}

		MappingFiles("test")

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
			"\t},\n" +
			"\t{\n" +
			"\t\t\"Accession\": [\n\t\t\t\"Q9BUL8\",\n\t\t\t\"A8K515\",\n\t\t\t\"D3DNN5\",\n\t\t\t\"O14811\"\n\t\t],\n" +
			"\t\t\"Biogrid\": \"\",\n" +
			"\t\t\"EnsemblGene\": [\n\t\t\t\"ENSG00000114209\"\n\t\t],\n" +
			"\t\t\"EnsemblProtein\": [\n\t\t\t\"ENSP00000376506\",\n\t\t\t\"ENSP00000417202\",\n\t\t\t\"ENSP00000418317\",\n\t\t\t\"ENSP00000420021\",\n\t\t\t\"ENSP00000420553\"\n\t\t],\n" +
			"\t\t\"Entrez\": \"11235\",\n" +
			"\t\t\"HGNC\": \"8761\",\n" +
			"\t\t\"ID\": \"PDC10_HUMAN\",\n" +
			"\t\t\"Name\": \"Programmed cell death protein 10\",\n" +
			"\t\t\"RefseqMRNA\": [\n\t\t\t\"NM_007217\",\n\t\t\t\"NM_145859\",\n\t\t\t\"NM_145860\",\n\t\t\t\"XM_005247086\"\n\t\t],\n" +
			"\t\t\"RefseqProtein\": [\n\t\t\t\"NP_009148\",\n\t\t\t\"NP_665858\",\n\t\t\t\"NP_665859\",\n\t\t\t\"XP_005247143\"\n\t\t],\n" +
			"\t\t\"Symbol\": [\n\t\t\t\"PDCD10\",\n\t\t\t\"CCM3\",\n\t\t\t\"TFAR15\"\n\t\t]\n" +
			"\t},\n" +
			"\t{\n" +
			"\t\t\"Accession\": [\n\t\t\t\"Q9P289\",\n\t\t\t\"B2RAU2\",\n\t\t\t\"Q3ZB77\",\n\t\t\t\"Q8NC04\",\n\t\t\t\"Q9BXC3\",\n\t\t\t\"Q9BXC4\"\n\t\t],\n" +
			"\t\t\"Biogrid\": \"\",\n" +
			"\t\t\"EnsemblGene\": [],\n" +
			"\t\t\"EnsemblProtein\": [],\n" +
			"\t\t\"Entrez\": \"\",\n" +
			"\t\t\"HGNC\": \"\",\n" +
			"\t\t\"ID\": \"STK26_HUMAN\",\n" +
			"\t\t\"Name\": \"Serine/threonine-protein kinase 26\",\n" +
			"\t\t\"RefseqMRNA\": [],\n" +
			"\t\t\"RefseqProtein\": [],\n" +
			"\t\t\"Symbol\": [\n\t\t\t\"STK26\",\n\t\t\t\"MASK\",\n\t\t\t\"MST4\"\n" +
			"\t\t]\n" +
			"\t}\n" +
			"]"

		outfile := fmt.Sprintf("test/genemap-%s.json", time.Now().Format("2006-01-02"))
		bytes, _ := afero.ReadFile(fs.Instance, outfile)
		Expect(string(bytes)).To(Equal(expected))
	})

	It("should fetch file from uniprot from command line", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		oldFTP := FTP
		defer func() { FTP = oldFTP }()
		FTP = func(url, sourceFile, targetFile string) error {
			zip.Gzip(uniprotText, "test/uniprot.dat.gz")
			return nil
		}

		options := map[string]interface{}{
			"folder": "test",
		}

		MappingFilesCMD(options)

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
			"\t},\n" +
			"\t{\n" +
			"\t\t\"Accession\": [\n\t\t\t\"Q9BUL8\",\n\t\t\t\"A8K515\",\n\t\t\t\"D3DNN5\",\n\t\t\t\"O14811\"\n\t\t],\n" +
			"\t\t\"Biogrid\": \"\",\n" +
			"\t\t\"EnsemblGene\": [\n\t\t\t\"ENSG00000114209\"\n\t\t],\n" +
			"\t\t\"EnsemblProtein\": [\n\t\t\t\"ENSP00000376506\",\n\t\t\t\"ENSP00000417202\",\n\t\t\t\"ENSP00000418317\",\n\t\t\t\"ENSP00000420021\",\n\t\t\t\"ENSP00000420553\"\n\t\t],\n" +
			"\t\t\"Entrez\": \"11235\",\n" +
			"\t\t\"HGNC\": \"8761\",\n" +
			"\t\t\"ID\": \"PDC10_HUMAN\",\n" +
			"\t\t\"Name\": \"Programmed cell death protein 10\",\n" +
			"\t\t\"RefseqMRNA\": [\n\t\t\t\"NM_007217\",\n\t\t\t\"NM_145859\",\n\t\t\t\"NM_145860\",\n\t\t\t\"XM_005247086\"\n\t\t],\n" +
			"\t\t\"RefseqProtein\": [\n\t\t\t\"NP_009148\",\n\t\t\t\"NP_665858\",\n\t\t\t\"NP_665859\",\n\t\t\t\"XP_005247143\"\n\t\t],\n" +
			"\t\t\"Symbol\": [\n\t\t\t\"PDCD10\",\n\t\t\t\"CCM3\",\n\t\t\t\"TFAR15\"\n\t\t]\n" +
			"\t},\n" +
			"\t{\n" +
			"\t\t\"Accession\": [\n\t\t\t\"Q9P289\",\n\t\t\t\"B2RAU2\",\n\t\t\t\"Q3ZB77\",\n\t\t\t\"Q8NC04\",\n\t\t\t\"Q9BXC3\",\n\t\t\t\"Q9BXC4\"\n\t\t],\n" +
			"\t\t\"Biogrid\": \"\",\n" +
			"\t\t\"EnsemblGene\": [],\n" +
			"\t\t\"EnsemblProtein\": [],\n" +
			"\t\t\"Entrez\": \"\",\n" +
			"\t\t\"HGNC\": \"\",\n" +
			"\t\t\"ID\": \"STK26_HUMAN\",\n" +
			"\t\t\"Name\": \"Serine/threonine-protein kinase 26\",\n" +
			"\t\t\"RefseqMRNA\": [],\n" +
			"\t\t\"RefseqProtein\": [],\n" +
			"\t\t\"Symbol\": [\n\t\t\t\"STK26\",\n\t\t\t\"MASK\",\n\t\t\t\"MST4\"\n" +
			"\t\t]\n" +
			"\t}\n" +
			"]"

		outfile := fmt.Sprintf("test/genemap-%s.json", time.Now().Format("2006-01-02"))
		bytes, _ := afero.ReadFile(fs.Instance, outfile)
		Expect(string(bytes)).To(Equal(expected))
	})
})
