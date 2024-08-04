package kata_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
    . "fundamentals/kata"
)

var sampleTest = Describe("Sample Tests", func() {
	It("it should work with the sample tests", func() {
		Expect(UnluckyDays(2015)).To(Equal(3))
		Expect(UnluckyDays(1986)).To(Equal(1))
	})
})