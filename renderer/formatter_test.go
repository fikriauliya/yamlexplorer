package renderer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formatter", func() {
	Describe("Trim", func() {
		It("trims to the max lenght", func() {
			sut := "Hello"

			res := trim(sut, 2)

			Expect(res).To(BeIdenticalTo("He"))
		})
		Context("when max value < 0", func() {
			It("returns empty string", func() {
				sut := "Hello"

				res := trim(sut, -1)

				Expect(res).To(BeIdenticalTo(""))
			})
		})
		Context("when max value == 0", func() {
			It("returns empty string", func() {
				sut := "Hello"

				res := trim(sut, 0)

				Expect(res).To(BeIdenticalTo(""))
			})
		})
	})
})
