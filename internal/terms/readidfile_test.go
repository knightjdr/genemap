package terms

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/genemap/pkg/fs"
	"github.com/spf13/afero"
)

var idText = `ACCS
CCM3
MST4
STK24
YWHAB
`

var _ = Describe("Read ID file", func() {
	It("should read IDs from file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/ids.txt", []byte(idText), 0444)

		expected := []string{"ACCS", "CCM3", "MST4", "STK24", "YWHAB"}
		Expect(readIDFile("test/ids.txt")).To(Equal(expected))
	})
})
