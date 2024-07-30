// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)

func dotest(text, char, expected string) {
     Expect(Contamination(text, char)).To(Equal(expected), "With text = \"%s\", char = \"%s\"", text, char)
}

var _ = Describe("Tests", func() {
     It("Sample tests", func() {
        dotest("abc","z", "zzz")
        dotest("","z", "")
        dotest("abc","", "")
        dotest("_3ebzgh4","&", "&&&&&&&&")
        dotest("//case"," ", "      ")
     })
})
