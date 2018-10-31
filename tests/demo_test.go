package demo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DemoTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "First Test")
}

var _ = Describe("First Test", func() {
	It("Should Pass", func() {
		Expect(100).To(Equal(100))
		Expect(true).To(Equal(true))
	})
})
