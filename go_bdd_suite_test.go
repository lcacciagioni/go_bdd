package go_bdd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoBdd Suite")
}
