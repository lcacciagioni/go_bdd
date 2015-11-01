package hello_test

import (
	. "github.com/lcacciagioni/go_bdd/hello"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hello", func() {
	Describe("The function", func() {
		It("Must return hello world", func() {
			Expect(Hello()).To(BeEquivalentTo("Hello world!!!"))
		})
	})
})
