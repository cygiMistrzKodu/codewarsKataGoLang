package kata_test
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)
var _ = Describe("basic tests", func() {
    It("Testing for attitude", func() { Expect(WordsToMarks("attitude")).To(Equal(100)) })
    It("Testing for friends", func() { Expect(WordsToMarks("friends")).To(Equal(75)) })
    It("Testing for family", func() { Expect(WordsToMarks("family")).To(Equal(66)) })
    It("Testing for selfness", func() { Expect(WordsToMarks("selfness")).To(Equal(99)) })
    It("Testing for knowledge", func() { Expect(WordsToMarks("knowledge")).To(Equal(96)) })
})