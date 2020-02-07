package terms

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate types", func() {
	It("should return an error when FromType is not specified", func() {
		m := Mapper{
			ToType: "accession",
		}

		expected := "Must specify FromType identifier for conversion"
		Expect(validateTypes(&m)).To(MatchError(expected))
	})

	It("should return an error when ToType is not specified", func() {
		m := Mapper{
			FromType: "accession",
		}

		expected := "Must specify ToType identifier for conversion"
		Expect(validateTypes(&m)).To(MatchError(expected))
	})

	It("should return an error when FromType is not valid identifer", func() {
		m := Mapper{
			FromType: "unknown",
			ToType:   "accession",
		}

		expected := "FromType identifier (\"unknown\") is invalid"
		Expect(validateTypes(&m)).To(MatchError(expected))
	})

	It("should return an error when ToType is not valid identifer", func() {
		m := Mapper{
			FromType: "accession",
			ToType:   "unknown",
		}

		expected := "ToType identifier (\"unknown\") is invalid"
		Expect(validateTypes(&m)).To(MatchError(expected))
	})

	It("should return an error when conversion types are the same", func() {
		m := Mapper{
			FromType: "accession",
			ToType:   "accession",
		}

		expected := "Conversion types are identifical (\"Accession\")"
		Expect(validateTypes(&m)).To(MatchError(expected))
	})

	It("should ignore case when checking conversion types", func() {
		m := Mapper{
			FromType: "AccESsion",
			ToType:   "BIOGRID",
		}

		Expect(validateTypes(&m)).To(BeNil())
	})
})
