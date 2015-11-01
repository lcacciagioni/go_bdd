package time_test

import (
	. "github.com/lcacciagioni/go_bdd/time"

	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Time", func() {
	It("Must return the current time", func() {
		Expect(CurrentTime("")).To(BeEquivalentTo(time.Now().Format(time.UnixDate)))
	})
	It("Must return a full string plus the current time", func() {
		Expect(CurrentTime("The time is: ")).To(BeEquivalentTo("The time is: " + time.Now().Format(time.UnixDate)))
	})
})
