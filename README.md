# Map between gene identifiers in GO

## Currently supported identifiers

| Identifier | Setting name | Examples |
|-------------------|----------------|-----------------|
| BioGRID           | Biogrid        | 116400          |
| Ensembl gene      | EnsemblGene    | ENSG00000114209 |
| Ensembl protein   | EnsemblProtein | ENSP00000376506 |
| Entrez gene ID    | Entrez         | 11235           |
| gene symbol       | Symbol         | PDCD10, CCM3    |
| HGNC              | HGNC           | 8761            |
| Refseq mRNA       | RefseqMRNA     | NM_007217       |
| Refseq protein    | RefseqProtein  | NP_009148       |
| UniProt Accession | Accession      | Q9BUL8          |
| UniProt ID        | ID             | PDC10_HUMAN     |


## Installation for CLI (optional)

```
> go get -v github.com/knightjdr/genemap
> cd $GOPATH/src/github.com/knightjdr/genemap
> go install ./...
```

## Usage

### Generate file for mapping.

The first step in mapping is to download reviewed UniProt entries and create a JSON file containing
supported conversion identifiers. Currently this only supports human genes. By default the file
will be created in the current directory but a target directory can be passed as a flag/option. The
created mapping file will have the creation date in the name formatted like `genemap-YYYY-MM-DD.json`.
This allows you to have your own database for mapping and update it as you see fit.

#### Import

```
import "github.com/knightjdr/genemap"

func main() {
  genemap.Generate("/path/to/output/folder/")
}
```

#### CLI

The map generation step can be run from the command line

```
> genemapcreate --folder="/path/to/output/folder/"
```

### Convert identifiers

You can use the mapping file created in the previous step to convert between supported gene identifiers. The
`Mapper` struct is used for setting mapping parameters and will contain the results of the
mapping.

#### Import

```
import "github.com/knightjdr/genemap"

func main() {
  mapper := genemap.CreateMapper()
  mapper.Load("/path/to/map/genemap.json")
  mapper.FromType = "Symbol"
  mapper.ToType = "Accession"

  err := mapper.Convert(ids)
  if err != nil {
    log.Fatal("Error converting IDs")
  }

  fmt.Println(mapper.Converted)
}
```

The mapper struct looks like so:

```
type Mapper struct {
	Converted          map[string]string
	FromType           string
	PossibleConversions map[string][]string
	ToType             string
	Unconverted        []string
}
```

`Converted` will have a one-to-one mapping of identifiers between the requested types and
`Unconverted` will contain a list of identifiers that could not be converted.
`PossibleConversions` will contain all possible mappings that the supplied identifier could be
converted to. For example, if you request a mapping to UniProt identifiers, `Converted` will
have a mapping to a reviewed UniProt entry while `PossibleConversions` will have all possible consistent
UniProt entries, whether reviewed or not. When converting to UniProt accession or ID, the `Converted`
map will always map to the reviewed entry. Similarly, when converting to gene Symbol, the entry
will always map to the official gene symbol. When mapping to other conversion types, the `Converted`
map will simply contain one of the possible identifiers it can be converted to if there are multiple.

#### CLI

ID conversion can be done from the command line. It requires a mapping file (.json format) and a text
file with IDs to convert (one entry per line). It will output a text file with the converted identifiers.
The output file name will be called `conversion.txt` if one is not supplied.

```
> genemapconvert --fromType="Symbol" --toType=="Accession" --idFile="ids.txt" --mapFile="genemap.json" --outfile="my-out-file.txt"
```