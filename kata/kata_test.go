// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)

func dotest(x1, x2 int, expected [3]int) {
     Expect(Quadratic(x1, x2)).To(Equal(expected), "With x1 = %d, x2 = %d", x1, x2)
}

var _ = Describe("Tests", func() {
     It("Sample tests", func() {
      dotest(0, 1, [3]int{1, -1, 0})
      dotest(4, 9, [3]int{1, -13, 36})
      dotest(2, 6, [3]int{1, -8, 12})
      dotest(-5, -4, [3]int{1, 9, 20})
     })
})
