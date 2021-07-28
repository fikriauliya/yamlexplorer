package parser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Parser", func() {
	Describe("Extract Header", func() {
		It("returns only the header", func() {
			sut := yaml.MapSlice{
				yaml.MapItem{
					Key: "data",
					Value: []interface{}{
						yaml.MapSlice{
							yaml.MapItem{Key: "name", Value: "alice"},
							yaml.MapItem{Key: "age", Value: "23"},
						},
						yaml.MapSlice{
							yaml.MapItem{Key: "name", Value: "bob"},
							yaml.MapItem{Key: "age", Value: "30"},
						},
					},
				},
			}

			result, err := extractHeader(sut)

			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(ConsistOf("name", "age"))
		})
		It("always returns with the same sequence as the input", func() {
			sut := yaml.MapSlice{
				yaml.MapItem{
					Key: "data",
					Value: []interface{}{
						yaml.MapSlice{
							yaml.MapItem{Key: "age", Value: "23"},
							yaml.MapItem{Key: "name", Value: "alice"},
						},
						yaml.MapSlice{
							yaml.MapItem{Key: "age", Value: "30"},
							yaml.MapItem{Key: "name", Value: "bob"},
						},
					},
				},
			}

			result, err := extractHeader(sut)

			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]string{"age", "name"}))
		})
	})
	Describe("Extract Body", func() {
		It("returns only the body", func() {
			sut := yaml.MapSlice{
				yaml.MapItem{
					Key: "data",
					Value: []interface{}{
						yaml.MapSlice{
							yaml.MapItem{Key: "name", Value: "alice"},
							yaml.MapItem{Key: "age", Value: "23"},
						},
						yaml.MapSlice{
							yaml.MapItem{Key: "name", Value: "bob"},
							yaml.MapItem{Key: "age", Value: "30"},
						},
					},
				},
			}
			result, err := extractBody([]string{"name", "age"}, sut)

			Expect(err).ToNot(HaveOccurred())
			Expect(result[0]).To(ConsistOf("alice", "23"))
			Expect(result[1]).To(ConsistOf("bob", "30"))
		})
		It("returns the body with the same order as the header", func() {
			sut := yaml.MapSlice{
				yaml.MapItem{
					Key: "data",
					Value: []interface{}{
						yaml.MapSlice{
							yaml.MapItem{Key: "name", Value: "alice"},
							yaml.MapItem{Key: "age", Value: "23"},
						},
						yaml.MapSlice{
							yaml.MapItem{Key: "name", Value: "bob"},
							yaml.MapItem{Key: "age", Value: "30"},
						},
					},
				},
			}
			result, err := extractBody([]string{"name", "age"}, sut)

			Expect(err).ToNot(HaveOccurred())
			Expect(result[0]).To(BeEquivalentTo([]string{"alice", "23"}))
			Expect(result[1]).To(BeEquivalentTo([]string{"bob", "30"}))

			result, err = extractBody([]string{"age", "name"}, sut)

			Expect(err).ToNot(HaveOccurred())
			Expect(result[0]).To(BeEquivalentTo([]string{"23", "alice"}))
			Expect(result[1]).To(BeEquivalentTo([]string{"30", "bob"}))
		})
	})
})
