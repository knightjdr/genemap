package flags

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert", func() {
	Context("float", func() {
		It("should return input float", func() {
			input := 0.01
			expected := 0.01
			Expect(convertFloat(input)).To(Equal(expected))
		})

		It("should convert an int to a float", func() {
			input := 1
			expected := float64(1)
			Expect(convertFloat(input)).To(Equal(expected))
		})

		It("should convert a string to a float", func() {
			input := "0.01"
			expected := 0.01
			Expect(convertFloat(input)).To(Equal(expected))
		})

		It("should return 0 for unrecognized type", func() {
			input := int64(1)
			expected := float64(0)
			Expect(convertFloat(input)).To(Equal(expected))
		})

		It("should return 0 for nil", func() {
			expected := float64(0)
			Expect(convertFloat(nil)).To(Equal(expected))
		})
	})

	Context("int", func() {
		It("should return input int", func() {
			input := 1
			expected := 1
			Expect(convertInt(input)).To(Equal(expected))
		})

		It("should convert a float to an int", func() {
			input := 1.4
			expected := 1
			Expect(convertInt(input)).To(Equal(expected))
		})

		It("should convert a string to an int", func() {
			input := "1"
			expected := 1
			Expect(convertInt(input)).To(Equal(expected))
		})

		It("should return 0 for unrecognized type", func() {
			input := int64(1)
			expected := 0
			Expect(convertInt(input)).To(Equal(expected))
		})

		It("should return 0 for nil", func() {
			expected := 0
			Expect(convertInt(nil)).To(Equal(expected))
		})
	})

	Context("string", func() {
		It("should return input string", func() {
			input := "a"
			expected := "a"
			Expect(convertString(input)).To(Equal(expected))
		})

		It("should return empty string for nil", func() {
			expected := ""
			Expect(convertString(nil)).To(Equal(expected))
		})
	})
})
