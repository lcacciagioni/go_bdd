package nextAfter_test

import (
	. "github.com/lcacciagioni/go_bdd/nextAfter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/rand"
)

var _ = Describe("NextAfter", func() {
	It("Must return the next after 2 numbers", func() {
		num1 := rand.Intn(1000)
		num2 := num1 + rand.Intn(1000)
		result := num1 + (2 * (num2 - num1))
		Expect(NextAfter(num1, num2)).To(BeEquivalentTo(result))
	})
	It("Must throw an exception if number 2 is smaller than number1", func() {
		num1 := rand.Intn(1000)
		num2 := num1 + rand.Intn(1000)
		_, err := NextAfter(num2, num1)
		Expect(err).To(HaveOccurred())
	})
	It("Must not throw any exception if the numbers are ok", func() {
		num1 := rand.Intn(1000)
		num2 := num1 + rand.Intn(1000)
		_, err := NextAfter(num1, num2)
		Expect(err).NotTo(HaveOccurred())
	})
})
