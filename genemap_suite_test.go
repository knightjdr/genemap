package genemap_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGenemap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Genemap Suite")
}
