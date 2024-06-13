

package kata_test
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)


var _ = Describe("Tests", func() {
     It("Sample tests", func() {
       Expect(Between(1, 4)).To(Equal([]int{1, 2, 3, 4}))
       Expect(Between(-2, 2)).To(Equal([]int{-2, -1, 0, 1, 2}))
     })
})
