package kata_test
import (
  ."github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)

func dotest(n int, expected bool) {
     Expect(AmIWilson(n)).To(Equal(expected), "With n = %d", n)
}

var _ = Describe("Tests", func() {
     It("Sample tests", func() {
       dotest(0, false)
       dotest(1, false)
       dotest(5, true)
       dotest(8, false)
       dotest(9, false)
       dotest(101, false)
       dotest(147, false)
       dotest(151, false)
     })
})