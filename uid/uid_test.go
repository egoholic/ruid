package uid_test

import (
	. "github.com/egoholic/ruid/uid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UID", func() {
	var (
		label1 = "label1"
		label2 = "label2"
	)

	Describe("New()", func() {
		It("returns uid string", func() {
			uid1 := New([]string{label1})
			uid2 := New([]string{label2})
			Expect(uid1).To(HavePrefix(label1))
			Expect(uid2).To(HavePrefix(label2))
			Expect(uid1).NotTo(Equal(uid2))
		})
	})
})
