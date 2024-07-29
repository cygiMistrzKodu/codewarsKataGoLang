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
       Expect(BinToDec("0")).To(Equal(0))
       Expect(BinToDec("1")).To(Equal(1))
       Expect(BinToDec("10")).To(Equal(2))
       Expect(BinToDec("11")).To(Equal(3))
       Expect(BinToDec("101010")).To(Equal(42))
       Expect(BinToDec("1001001")).To(Equal(73))
     })
})
