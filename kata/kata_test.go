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
       Expect(QueueTime([]int{}, 1)).To(Equal(0))
       Expect(QueueTime([]int{1,2,3,4}, 1)).To(Equal(10))
       Expect(QueueTime([]int{2,2,3,3,4,4}, 2)).To(Equal(9))
       Expect(QueueTime([]int{1,2,3,4,5}, 100)).To(Equal(5))
     })
})

