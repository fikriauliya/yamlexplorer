package renderer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resizer", func() {
	Describe("Transpose", func() {
		It("transpose a matrix", func() {
			sut := [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
			}

			res := transpose(sut)

			Expect(res).To(BeEquivalentTo([][]string{
				{"1", "4"},
				{"2", "5"},
				{"3", "6"},
			}))
		})
		Context("when the matrix is empty", func() {
			It("returns empty matrix", func() {
				sut := [][]string{{}}

				res := transpose(sut)

				Expect(res).To(BeEquivalentTo([][]string{{}}))
			})
		})
	})
	Describe("Len Matrix", func() {
		It("returns the length of each matrix element", func() {
			sut := [][]string{
				{"Hello", "Hi"},
				{"1", "The"},
			}

			res := lenMatrix(sut)

			Expect(res).To(BeEquivalentTo([][]float64{
				{5, 2},
				{1, 3},
			}))
		})
	})
	Describe("Resize", func() {
		It("returns a matrix with trimmed content based on median length of the column", func() {
			sut := [][]string{
				{"Hello", "Hi"},
				{"1", "The"},
			}

			res, _ := Resize(sut, 100)

			Expect(res).To(BeEquivalentTo(&[][]string{
				{"Hel", "Hi"},
				{"1", "Th"},
			}))
		})
	})
})
