package parser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Describe("Extract Header", func() {
		It("returns only the header", func() {
			sut := make([]interface{}, 2)
			sut[0] = map[interface{}]interface{}{
				"name": "alice",
				"age":  "23"}
			sut[1] = map[interface{}]interface{}{
				"name": "bob",
				"age":  "30",
			}

			result := extractHeader([]interface{}(sut), nil)

			Expect(result).To(ConsistOf("name", "age"))
		})
		Context("when order is specified", func() {
			It("always returns with the same sequence as the given order", func() {
				sut := make([]interface{}, 2)
				sut[0] = map[interface{}]interface{}{
					"name": "alice",
					"age":  "23"}
				sut[1] = map[interface{}]interface{}{
					"name": "bob",
					"age":  "30",
				}
				order := []string{"age", "name"}

				result := extractHeader([]interface{}(sut), order)

				Expect(result).To(BeEquivalentTo([]string{"age", "name"}))
			})
		})
	})
	Describe("Extract Body", func() {
		It("returns only the body", func() {
			sut := make([]interface{}, 2)
			sut[0] = map[interface{}]interface{}{
				"name": "alice",
				"age":  "23"}
			sut[1] = map[interface{}]interface{}{
				"name": "bob",
				"age":  "30",
			}
			result := extractBody([]string{"name", "age"}, sut)

			Expect(result[0]).To(ConsistOf("alice", "23"))
			Expect(result[1]).To(ConsistOf("bob", "30"))
		})
		It("returns the body with the same order as the header", func() {
			sut := make([]interface{}, 2)
			sut[0] = map[interface{}]interface{}{
				"name": "alice",
				"age":  "23"}
			sut[1] = map[interface{}]interface{}{
				"name": "bob",
				"age":  "30",
			}
			result := extractBody([]string{"name", "age"}, sut)

			Expect(result[0]).To(BeEquivalentTo([]string{"alice", "23"}))
			Expect(result[1]).To(BeEquivalentTo([]string{"bob", "30"}))

			result = extractBody([]string{"age", "name"}, sut)

			Expect(result[0]).To(BeEquivalentTo([]string{"23", "alice"}))
			Expect(result[1]).To(BeEquivalentTo([]string{"30", "bob"}))
		})
	})
})
