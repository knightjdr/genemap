package generate

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
				"-folder", "outputfolder",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				folder: "outputfolder",
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
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.folder).To(Equal("."), "should set default output folder")
		})
	})

	Context("argument passed via input file", func() {
		It("should set variables from file", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{
				"folder": "outputfolder-alternative",
			}

			expected := parameters{
				folder: "outputfolder-alternative",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})

var _ = Describe("Parse folder name", func() {
	It("should remove trailing slash", func() {
		folder := "path/folder/"

		expected := "path/folder"
		Expect(parseFolderName(folder)).To(Equal(expected))
	})

	It("should return name that does not end in a slash", func() {
		folder := "path/folder"

		expected := "path/folder"
		Expect(parseFolderName(folder)).To(Equal(expected))
	})
})
