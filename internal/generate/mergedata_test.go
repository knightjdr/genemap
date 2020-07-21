package generate

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Merge entries", func() {
	It("should merge uniprot and hgnc entries", func() {
		hgnc := hgncRecords{
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
		uniprot := &uniprotRecords{
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
				Accession:      []string{"P04217"},
				EnsemblGene:    []string{},
				EnsemblProtein: []string{},
				Entrez:         "",
				HGNC:           "5",
				ID:             "",
				Name:           "",
				RefseqMRNA:     []string{},
				RefseqProtein:  []string{},
				Reviewed:       true,
				Symbol:         []string{"A1BG"},
			},
			uniprotRecord{
				Accession:      []string{"P1XXXXX"},
				EnsemblGene:    []string{},
				EnsemblProtein: []string{},
				Entrez:         "",
				HGNC:           "37133",
				ID:             "",
				Name:           "",
				RefseqMRNA:     []string{},
				RefseqProtein:  []string{},
				Reviewed:       true,
				Symbol:         []string{},
			},
		}

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
				Accession:      []string{"P04217"},
				EnsemblGene:    []string{},
				EnsemblProtein: []string{},
				Entrez:         "1",
				HGNC:           "5",
				ID:             "",
				Name:           "",
				RefseqMRNA:     []string{},
				RefseqProtein:  []string{},
				Reviewed:       true,
				Symbol:         []string{"A1BG"},
			},
			uniprotRecord{
				Accession:      []string{"P1XXXXX"},
				EnsemblGene:    []string{},
				EnsemblProtein: []string{},
				Entrez:         "503538",
				HGNC:           "37133",
				ID:             "",
				Name:           "",
				RefseqMRNA:     []string{},
				RefseqProtein:  []string{},
				Reviewed:       true,
				Symbol:         []string{"A1BG-AS1"},
			},
		}

		Expect(mergeData(uniprot, hgnc)).To(Equal(expected))
	})
})
