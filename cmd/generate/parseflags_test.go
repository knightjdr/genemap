package generate

import (
	"os"

	"github.com/knightjdr/genemap/internal/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var jsonText = `{
	"module": "test-module"
}`

var _ = BeforeSuite(func() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
})

var _ = Describe("Parseflags", func() {
	It("should parse arguments from file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/options.json",
			[]byte(jsonText),
			0444,
		)

		os.Args = []string{
			"cmd",
			"-options", "test/options.json",
		}

		expected := map[string]interface{}{
			"module": "test-module",
		}
		options, err := parseFlags()
		Expect(err).To(BeNil(), "should not return an error with complete options file")
		Expect(options).To(Equal(expected), "should return JSON as interface")
	})
})
