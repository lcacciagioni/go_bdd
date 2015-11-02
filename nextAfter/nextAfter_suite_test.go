package nextAfter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNextAfter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NextAfter Suite")
}
