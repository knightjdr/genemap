package generate

import (
	"github.com/knightjdr/genemap/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

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
ID   PDC10_HUMAN             Reviewed;         212 AA.
AC   Q9BUL8; A8K515; D3DNN5; O14811;
DT   11-OCT-2005, integrated into UniProtKB/Swiss-Prot.
DT   01-JUN-2001, sequence version 1.
DT   11-DEC-2019, entry version 152.
DE   RecName: Full=Programmed cell death protein 10;
DE   AltName: Full=Cerebral cavernous malformations 3 protein;
DE   AltName: Full=TF-1 cell apoptosis-related protein 15;
GN   Name=PDCD10; Synonyms=CCM3, TFAR15;
OS   Homo sapiens (Human).
DR   Ensembl; ENST00000392750; ENSP00000376506; ENSG00000114209.
DR   Ensembl; ENST00000461494; ENSP00000420021; ENSG00000114209.
DR   Ensembl; ENST00000470131; ENSP00000417202; ENSG00000114209.
DR   Ensembl; ENST00000473645; ENSP00000418317; ENSG00000114209.
DR   Ensembl; ENST00000497056; ENSP00000420553; ENSG00000114209.
DR   GeneID; 11235; -.
DR   RefSeq; NP_009148.2; NM_007217.3.
DR   RefSeq; NP_665858.1; NM_145859.1.
DR   RefSeq; NP_665859.1; NM_145860.1.
DR   RefSeq; XP_005247143.1; XM_005247086.4.
DR   HGNC; HGNC:8761; PDCD10.
//
ID   STK26_HUMAN             Reviewed;         416 AA.
AC   Q9P289; B2RAU2; Q3ZB77; Q8NC04; Q9BXC3; Q9BXC4;
DT   16-AUG-2005, integrated into UniProtKB/Swiss-Prot.
DT   01-OCT-2001, sequence version 2.
DT   11-DEC-2019, entry version 165.
DE   RecName: Full=Serine/threonine-protein kinase 26 {ECO:0000305};
DE            EC=2.7.11.1 {ECO:0000269|PubMed:11641781};
DE   AltName: Full=MST3 and SOK1-related kinase {ECO:0000303|PubMed:11741893};
DE   AltName: Full=Mammalian STE20-like protein kinase 4 {ECO:0000303|PubMed:11641781};
DE            Short=MST-4 {ECO:0000305};
DE            Short=STE20-like kinase MST4 {ECO:0000305};
DE   AltName: Full=Serine/threonine-protein kinase MASK {ECO:0000305};
GN   Name=STK26 {ECO:0000312|HGNC:HGNC:18174};
GN   Synonyms=MASK {ECO:0000303|PubMed:11741893},
GN   MST4 {ECO:0000303|PubMed:11641781};
//
`

var _ = Describe("Parse uniprot", func() {
	It("should parse entries from uniprot file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/uniprot.dat",
			[]byte(uniprotText),
			0444,
		)

		expected := &uniprotRecords{
			uniprotRecord{
				Accession:      []string{"P31946", "A8K9K2", "E1P616"},
				Biogrid:        "113361",
				EnsemblGene:    []string{"ENSG00000166913"},
				EnsemblProtein: []string{"ENSP00000300161", "ENSP00000361930"},
				Entrez:         "7529",
				HGNC:           "12849",
				ID:             "1433B_HUMAN",
				Name:           "14-3-3 protein beta/alpha",
				RefseqMRNA:     []string{"NM_003404", "NM_139323", "XM_017028039"},
				RefseqProtein:  []string{"NP_003395", "NP_647539", "XP_016883528"},
				Reviewed:       true,
				Symbol:         []string{"YWHAB"},
			},
			uniprotRecord{
				Accession:      []string{"Q96QU6", "B4E219", "Q8WUL4", "Q96LX5"},
				EnsemblGene:    []string{"ENSG00000110455"},
				EnsemblProtein: []string{"ENSP00000263776"},
				Entrez:         "84680",
				HGNC:           "23989",
				ID:             "1A1L1_HUMAN",
				Name:           "1-aminocyclopropane-1-carboxylate synthase-like protein 1",
				RefseqMRNA:     []string{"NM_001127219", "NM_032592"},
				RefseqProtein:  []string{"NP_001120691", "NP_115981"},
				Reviewed:       true,
				Symbol:         []string{"ACCS", "PHACS"},
			},
			uniprotRecord{
				Accession:      []string{"Q9BUL8", "A8K515", "D3DNN5", "O14811"},
				EnsemblGene:    []string{"ENSG00000114209"},
				EnsemblProtein: []string{"ENSP00000376506", "ENSP00000417202", "ENSP00000418317", "ENSP00000420021", "ENSP00000420553"},
				Entrez:         "11235",
				HGNC:           "8761",
				ID:             "PDC10_HUMAN",
				Name:           "Programmed cell death protein 10",
				RefseqMRNA:     []string{"NM_007217", "NM_145859", "NM_145860", "XM_005247086"},
				RefseqProtein:  []string{"NP_009148", "NP_665858", "NP_665859", "XP_005247143"},
				Reviewed:       true,
				Symbol:         []string{"PDCD10", "CCM3", "TFAR15"},
			},
			uniprotRecord{
				Accession:      []string{"Q9P289", "B2RAU2", "Q3ZB77", "Q8NC04", "Q9BXC3", "Q9BXC4"},
				EnsemblGene:    []string{},
				EnsemblProtein: []string{},
				ID:             "STK26_HUMAN",
				Name:           "Serine/threonine-protein kinase 26",
				RefseqMRNA:     []string{},
				RefseqProtein:  []string{},
				Reviewed:       true,
				Symbol:         []string{"STK26", "MASK", "MST4"},
			},
		}
		Expect(parseUniprot("test")).To(Equal(expected))
	})
})
