package generate

import (
	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/knightjdr/genemap/pkg/zip"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Fetch uniprot", func() {
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

		fetchUniprot("test")

		Expect(afero.Exists(fs.Instance, "test/uniprot.dat.gz")).To(BeFalse(), "should remove gz file")
		Expect(afero.Exists(fs.Instance, "test/uniprot.dat")).To(BeTrue(), "should unzip gz file")

		bytes, _ := afero.ReadFile(fs.Instance, "test/uniprot.dat")
		Expect(string(bytes)).To(Equal(uniprotText), "should download uniprot DB")
	})
})
