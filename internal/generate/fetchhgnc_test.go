package generate

import (
	"github.com/knightjdr/genemap/pkg/fs"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Fetch hgnc", func() {
	It("should fetch file from hgnc", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		oldHTTP := HTTP
		defer func() { HTTP = oldHTTP }()
		HTTP = func(url string, headers map[string]string, targetFile string) error {
			afero.WriteFile(fs.Instance, "test/hgnc.json", []byte(hgncText), 0644)
			return nil
		}

		fetchHGNC("test")

		bytes, _ := afero.ReadFile(fs.Instance, "test/hgnc.json")
		Expect(string(bytes)).To(Equal(hgncText), "should download hgnc DB")
	})
})
