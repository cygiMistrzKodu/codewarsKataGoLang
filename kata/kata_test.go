package kata_test

import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
  "fmt"
)

func doTest(n int, expected int) {
  It(fmt.Sprintf("n = %d", n), func() {
    Expect(NextHigher(n)).To(Equal(expected))
  })
}

var _ = Describe("Testing NextHigher", func() {
  Context("Basic tests", func() {
    doTest(1022, 1279)
    doTest(128, 256)
    doTest(1, 2)
    doTest(127, 191)
    doTest(1253343, 1253359)
  })
})
