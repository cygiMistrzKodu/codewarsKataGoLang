

package kata_test
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)

func dotest(c, e int, expected string) {
     Expect(Derive(c, e)).To(Equal(expected), "With coefficient = %d, exponent = %d", c, e)
}

var _ = Describe("Tests", func() {
     It("Sample tests", func() {
          dotest(7, 8, "56x^7")
          dotest(5, 9, "45x^8")
     })
})
