package terms

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
})

var _ = Describe("Parseflags", func() {
	Context("all command line arguments", func() {
		It("should parse arguments", func() {
			os.Args = []string{
				"cmd",
				"-fromType", "Symbol",
				"-idFile", "ids.txt",
				"-mapFile", "mapping.json",
				"-outFile", "out.txt",
				"-toType", "Accession",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				fromType: "Symbol",
				idFile:   "ids.txt",
				mapFile:  "mapping.json",
				outFile:  "out.txt",
				toType:   "Accession",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})

	Context("only required command line arguments", func() {
		It("should set defaults", func() {
			os.Args = []string{
				"cmd",
				"-fromType", "Symbol",
				"-idFile", "ids.txt",
				"-mapFile", "mapping.json",
				"-toType", "Accession",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.outFile).To(Equal("conversion.txt"), "should set default output file")
		})
	})

	Context("missing required command line arguments", func() {
		It("should report error", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{}

			_, err := parseFlags(fileOptions)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("argument passed via input file", func() {
		It("should set variables from file", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{
				"fromType": "Symbol",
				"idFile":   "file-ids.txt",
				"mapFile":  "mapping.json",
				"outFile":  "file-out.txt",
				"toType":   "Accession",
			}

			expected := parameters{
				fromType: "Symbol",
				idFile:   "file-ids.txt",
				mapFile:  "mapping.json",
				outFile:  "file-out.txt",
				toType:   "Accession",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
