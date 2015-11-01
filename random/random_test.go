package random_test

import (
	"regexp"

	. "github.com/lcacciagioni/go_bdd/random"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Random", func() {
	r := regexp.MustCompile(`This is a random number \d+`)
	It("Must return a string plus a random number", func() {
		Expect(r.MatchString(Random())).To(BeTrue())
	})
})
