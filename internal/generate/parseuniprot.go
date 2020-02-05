package generate

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/slice"
	"github.com/knightjdr/genemap/pkg/fs"
)

type uniprotEntries []uniprotEntry

type uniprotEntry struct {
	Accession       []string `json:"accession"`
	EnsemblGene     []string `json:"ensemblg"`
	EnsembleProtein []string `json:"ensemblp"`
	Entrez          string   `json:"entrez"`
	HGNC            string   `json:"hgnc"`
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	RefseqMRNA      []string `json:"refseqm"`
	RefseqProtein   []string `json:"refseqp"`
	Reviewed        bool     `json:"-"`
	Symbol          []string `json:"symbol"`
}

func parseUniprot(folder string) *uniprotEntries {
	file, err := fs.Instance.Open(fmt.Sprintf("%s/uniprot.dat", folder))
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	entries := &uniprotEntries{}
	entry := uniprotEntry{}
	re := createRe()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "ID") {
			entry = uniprotEntry{
				Accession: []string{},
				ID:        parseValue(re["id"], line),
				Reviewed:  isReviewed(line),
			}
		} else if strings.HasPrefix(line, "AC") {
			entry.Accession = append(entry.Accession, parseArrayValue(re["accession"], line)...)
		} else if strings.HasPrefix(line, "DE   RecName") {
			entry.Name = parseValue(re["name"], line)
		} else if strings.HasPrefix(line, "DR   Ensembl") {
			entry.EnsemblGene = append(entry.EnsemblGene, parseArrayValue(re["ensembl"], line)...)
		} else if strings.HasPrefix(line, "DR   GeneID") {
			entry.Entrez = parseValue(re["entrez"], line)
		} else if strings.HasPrefix(line, "DR   HGNC") {
			entry.HGNC = parseValue(re["hgnc"], line)
		} else if strings.HasPrefix(line, "DR   RefSeq") {
			entry.RefseqMRNA = append(entry.RefseqMRNA, parseArrayValue(re["refseq"], line)...)
		} else if strings.HasPrefix(line, "GN") {
			entry.Symbol = append(entry.Symbol, parseArrayValue(re["symbol"], line)...)
		} else if strings.HasPrefix(line, "//") {
			addEntry(entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return entries
}

func createRe() map[string]*regexp.Regexp {
	accession := regexp.MustCompile(`(\w+);`)
	ensembl := regexp.MustCompile(`(ENS[GP]\d+)`)
	entrez := regexp.MustCompile(`GeneID; (\w+);`)
	hgnc := regexp.MustCompile(`HGNC:(\d+)`)
	id := regexp.MustCompile(`^ID   (\w+)`)
	name := regexp.MustCompile(`RecName: Full=([^;{]+)`)
	refseq := regexp.MustCompile(`(\w[MP]_\d+)`)
	symbol := regexp.MustCompile(`(?:Name=|Synonyms=|^GN\s{3})([^;{]+)`)
	return map[string]*regexp.Regexp{
		"accession": accession,
		"ensembl":   ensembl,
		"entrez":    entrez,
		"hgnc":      hgnc,
		"id":        id,
		"name":      name,
		"refseq":    refseq,
		"symbol":    symbol,
	}
}

func isReviewed(line string) bool {
	return strings.Contains(line, "Reviewed")
}

func parseValue(re *regexp.Regexp, line string) string {
	matches := re.FindStringSubmatch(line)
	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}

func parseArrayValue(re *regexp.Regexp, line string) []string {
	matches := re.FindAllStringSubmatch(line, -1)
	if len(matches) > 0 {
		values := make([]string, len(matches))
		for i, match := range matches {
			values[i] = match[1]
		}
		return values
	}
	return []string{}
}

func addEntry(entries *uniprotEntries, entry uniprotEntry) {
	if entry.Reviewed {
		trimName(&entry)
		splitSymbols(&entry)
		separateEnsembl(&entry)
		separateRefseq(&entry)
		*entries = append(*entries, entry)
	}
}

func trimName(entry *uniprotEntry) {
	(*entry).Name = strings.TrimSpace((*entry).Name)
}

func splitSymbols(entry *uniprotEntry) {
	symbols := make([]string, 0)

	for _, symbol := range entry.Symbol {
		symbols = append(symbols, strings.Split(symbol, ", ")...)
	}

	for i, symbol := range symbols {
		symbol = strings.Replace(symbol, "Name=", "", 1)
		symbol = strings.Replace(symbol, "Synonyms=", "", 1)
		symbols[i] = strings.TrimSpace(symbol)
	}

	(*entry).Symbol = symbols
}

func separateEnsembl(entry *uniprotEntry) {
	gene := entry.EnsemblGene
	protein := make([]string, 0)

	for i := len(gene) - 1; i >= 0; i-- {
		if strings.HasPrefix(gene[i], "ENSP") {
			protein = append(protein, gene[i])
			gene = append(gene[:i], gene[i+1:]...)
		}
	}

	(*entry).EnsemblGene = slice.UniqueStrings(gene)
	(*entry).EnsembleProtein = slice.UniqueStrings(protein)

	sort.Strings((*entry).EnsemblGene)
	sort.Strings((*entry).EnsembleProtein)
}

func separateRefseq(entry *uniprotEntry) {
	mrna := entry.RefseqMRNA
	protein := make([]string, 0)

	for i := len(mrna) - 1; i >= 0; i-- {
		if strings.HasPrefix(mrna[i], "NP") || strings.HasPrefix(mrna[i], "XP") {
			protein = append(protein, mrna[i])
			mrna = append(mrna[:i], mrna[i+1:]...)
		}
	}

	(*entry).RefseqMRNA = slice.UniqueStrings(mrna)
	(*entry).RefseqProtein = slice.UniqueStrings(protein)

	sort.Strings((*entry).RefseqMRNA)
	sort.Strings((*entry).RefseqProtein)
}