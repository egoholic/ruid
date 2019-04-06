package uid_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uid Suite")
}
