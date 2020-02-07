# Map between gene identifiers in GO.

## Usage

### Generate file for mapping.

The first step in mapping is to download reviewed UniProt entries and create a JSON file containing
supported conversion identifiers. Currently this only supports human genes. By default the file
will be created in the current directory but a target directory can be passed as a flag/option. The
created mapping file will have the creation date in the name formatted like `genemap-DD-MM-YYYY.json`.

#### Import

```
import "github.com/knightjdr/genemap"

func main() {
  options := map[string]interface{}{
    folder: "./my-folder/",
  }

  genemap.Generate(options)
}
```

#### CLI

The map generation step can be run from the command line

```
> go get github.com/knightjdr/genemap
> cd $GOPATH/src/github.com/knightjdr/prohits-viz-analysis
> go install ./...
> genemap --folder="./my-folder/"
```