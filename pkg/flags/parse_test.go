package flags

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {
	It("should return command line arguments as interface", func() {
		os.Args = []string{
			"cmd",
			"-optiona=a",
			"--optionb", "1",
		}

		expected := map[string]interface{}{
			"optiona": "a",
			"optionb": "1",
		}
		Expect(Parse()).To(Equal(expected))
	})
})
