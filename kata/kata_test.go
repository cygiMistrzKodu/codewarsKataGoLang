// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	. "fundamentals/kata"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("basic tests", func() {
		Expect(Xor(false, false)).To(Equal(false))
		Expect(Xor(true, false)).To(Equal(true))
		Expect(Xor(false, true)).To(Equal(true))
		Expect(Xor(true, true)).To(Equal(false))
	})
	It("nested tests", func() {
		Expect(Xor(false, Xor(false, false))).To(Equal(false))
		Expect(Xor(Xor(true, false), false)).To(Equal(true))
		Expect(Xor(Xor(true, true), false)).To(Equal(false))
		Expect(Xor(true, Xor(true, true))).To(Equal(true))
		Expect(Xor(Xor(false, false), Xor(false, false))).To(Equal(false))
		Expect(Xor(Xor(false, false), Xor(false, true))).To(Equal(true))
		Expect(Xor(Xor(true, false), Xor(false, false))).To(Equal(true))
		Expect(Xor(Xor(true, false), Xor(true, false))).To(Equal(false))
		Expect(Xor(Xor(true, true), Xor(true, false))).To(Equal(true))
		Expect(Xor(Xor(true, Xor(true, true)), Xor(Xor(true, true), false))).To(Equal(true))
	})
})
