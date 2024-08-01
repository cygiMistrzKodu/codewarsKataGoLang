// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>


package kata_test
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)


var _ = Describe("Tests", func() {
     It("Sample tests", func() {
       arr := []int{3, 5, 8, 13}
       Expect(EachCons(arr, 1)).To(Equal([][]int{{3}, {5}, {8}, {13}}))
       Expect(EachCons(arr, 2)).To(Equal([][]int{{3, 5}, {5, 8}, {8, 13}}))
       Expect(EachCons(arr, 3)).To(Equal([][]int{{3, 5, 8}, {5, 8, 13}}))
       Expect(EachCons([]int{}, 3)).To(BeEmpty())
     })
})