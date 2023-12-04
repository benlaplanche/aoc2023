package day1_test

import (
	"github.com/benlaplanche/aoc2023/day1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day1", func() {
	var testCase string
	BeforeEach(func() {
		testCase = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	})

	Describe("Counting Stars", func() {
		Context("Line with two numbers", func() {
			It("Should return the correct values", func() {
				Expect(day1.Parse(testCase).To(HaveValue(Equal(1))))
			})
		})
	})

})
