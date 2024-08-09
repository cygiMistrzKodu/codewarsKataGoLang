package kata_test
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)
var _ = Describe("Split Strings", func() {
    It("Basic tests", func() {
        // Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
        Expect(Solution("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
        Expect(Solution("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
    })
})