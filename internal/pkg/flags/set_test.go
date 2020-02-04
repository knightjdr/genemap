package flags

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	Context("float", func() {
		It("should return command line argument", func() {
			args := map[string]interface{}{
				"arg": "1.5",
			}
			fileOptions := map[string]interface{}{
				"arg": "1.0",
			}

			Expect(SetFloat("arg", args, fileOptions, 0.5)).To(Equal(1.5))
		})

		It("should return file parameter", func() {
			args := map[string]interface{}{}
			fileOptions := map[string]interface{}{
				"arg": "1.0",
			}

			Expect(SetFloat("arg", args, fileOptions, 0.5)).To(Equal(1.0))
		})

		It("should return default value", func() {
			args := map[string]interface{}{}
			fileOptions := map[string]interface{}{}

			Expect(SetFloat("arg", args, fileOptions, 0.5)).To(Equal(0.5))
		})
	})

	Context("int", func() {
		It("should return command line argument", func() {
			args := map[string]interface{}{
				"arg": "3",
			}
			fileOptions := map[string]interface{}{
				"arg": "2",
			}

			Expect(SetInt("arg", args, fileOptions, 1)).To(Equal(3))
		})

		It("should return file parameter", func() {
			args := map[string]interface{}{}
			fileOptions := map[string]interface{}{
				"arg": "2",
			}

			Expect(SetInt("arg", args, fileOptions, 1)).To(Equal(2))
		})

		It("should return default value", func() {
			args := map[string]interface{}{}
			fileOptions := map[string]interface{}{}

			Expect(SetInt("arg", args, fileOptions, 1)).To(Equal(1))
		})
	})

	Context("string", func() {
		It("should return command line argument", func() {
			args := map[string]interface{}{
				"arg": "a",
			}
			fileOptions := map[string]interface{}{
				"arg": "b",
			}

			Expect(SetString("arg", args, fileOptions, "c")).To(Equal("a"))
		})

		It("should return file parameter", func() {
			args := map[string]interface{}{}
			fileOptions := map[string]interface{}{
				"arg": "b",
			}

			Expect(SetString("arg", args, fileOptions, "c")).To(Equal("b"))
		})

		It("should return default value", func() {
			args := map[string]interface{}{}
			fileOptions := map[string]interface{}{}

			Expect(SetString("arg", args, fileOptions, "c")).To(Equal("c"))
		})
	})
})
