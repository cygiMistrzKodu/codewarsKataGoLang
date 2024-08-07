package kata_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
  . "fundamentals/kata"
  "testing"
)

func TestDominator(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("What dominates your array?", func() {

	Context("Fixed tests", func() {
		It("should pass fixed tests", func() {
			Expect(Dominator([]int{3, 4, 3, 2, 3, 1, 3, 3})).To(Equal(3))
			Expect(Dominator([]int{1, 2, 3, 4, 5})).To(Equal(-1))
			Expect(Dominator([]int{1, 1, 1, 2, 2, 2})).To(Equal(-1))
			Expect(Dominator([]int{1, 1, 1, 2, 2, 2, 2})).To(Equal(2))
		})
	})

})