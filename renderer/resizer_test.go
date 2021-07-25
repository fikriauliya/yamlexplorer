package renderer

import (
	"github.com/fikriauliya/yamlexplorer/entity"
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

			Expect(res).To(Equal([][]string{
				{"1", "4"},
				{"2", "5"},
				{"3", "6"},
			}))
		})
		Context("when the matrix is empty", func() {
			It("returns empty matrix", func() {
				sut := [][]string{{}}

				res := transpose(sut)

				Expect(res).To(Equal([][]string{{}}))
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

			Expect(res).To(Equal([][]float64{
				{5, 2},
				{1, 3},
			}))
		})
	})
	Describe("Resize", func() {
		Context("total width is larger than total columns' width", func() {
			It("the last element takes the remaining space", func() {
				sut := entity.Table{
					Header: []string{"", ""},
					Body: [][]string{
						{"Hello", "Hi"},
						{"World", "The"},
					}}

				width, err := Resize(&sut, 100)

				Expect(err).NotTo(HaveOccurred())
				Expect(width).To(Equal([]int{5, 95}))
			})
		})
		Context("total width is smaller than total columns' width", func() {
			It("the width equals to the median length of the column", func() {
				sut := entity.Table{
					Header: []string{"", ""},
					Body: [][]string{
						{"Helloo", "Hi"},
						{"1", "The"},
					}}

				width, err := Resize(&sut, 5)

				Expect(err).NotTo(HaveOccurred())
				Expect(width).To(Equal([]int{3, 2}))
			})
			Context("the max length is only 2 larger than the median", func() {
				It("the width equals to the max length of the column", func() {
					sut := entity.Table{
						Header: []string{"", ""},
						Body: [][]string{
							{"Hello", "Hi"},
							{"12", "The"},
						}}

					width, err := Resize(&sut, 5)

					Expect(err).NotTo(HaveOccurred())
					Expect(width).To(Equal([]int{5, 0}))
				})
			})
		})
	})
})
