package terms_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTerms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Terms Suite")
}
