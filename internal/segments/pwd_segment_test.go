package segments

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pwd{}", func() {
	var wd, _ = os.Getwd()
	Describe("Output()", func() {
		It("is expected to be the current working directory", func() {
			Expect(Pwd{}.Output()).To(Equal(wd))
		})

		PContext("When the length of PWD is greater than 1/4th of terminalWidth", func() {})
	})

	Describe("Len()", func() {
		It("is expected to be the length of the current working directory.", func() {
			Expect(Pwd{}.Len()).To(Equal(len(wd)))
		})
	})
})
