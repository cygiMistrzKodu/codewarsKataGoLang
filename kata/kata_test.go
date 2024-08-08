package kata_test
 
import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "fundamentals/kata"
)
 
var _ = Describe("Tests using hard-coded strings", func() {
  It("Should return a correctly translated string", func() {
     Expect(ToNato("If you can read")).To(Equal("India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"))
     Expect(ToNato("Did not see that coming")).To(Equal("Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"))
     Expect(ToNato("go for it!")).To(Equal("Golf Oscar Foxtrot Oscar Romeo India Tango !"))
  })
})
 