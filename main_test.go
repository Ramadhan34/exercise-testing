package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	main "github.com/Ramadhan34/exercise-testing"
)

var _ = Describe("PopulationData", func() {
	It("should correctly transform the data", func() {
		data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
		expected := []map[string]interface{}{
			{"name": "Budi", "age": 23, "address": "Jakarta"},
			{"name": "Joko", "age": 30, "address": "Bandung", "isMarried": true},
			{"name": "Susi", "age": 25, "address": "Bogor", "height": 165.42},
		}

		result := main.PopulationData(data)
		Expect(result).To(Equal(expected))
	})

	It("should handle incorrect data format", func() {
		data := []string{"Budi;23;Jakarta;abc;true", "Joko;30;Bandung;165.42;xyz"}
		expected := []map[string]interface{}{
			{"name": "Budi", "age": 23, "address": "Jakarta", "isMarried": true},
			{"name": "Joko", "age": 30, "address": "Bandung"},
		}

		result := main.PopulationData(data)
		Expect(result).To(Equal(expected))
	})

	It("should handle empty data", func() {
		data := []string{}
		expected := []map[string]interface{}{}

		result := main.PopulationData(data)
		Expect(result).To(Equal(expected))
	})
})
