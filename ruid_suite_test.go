package ruid_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRuid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ruid Suite")
}
