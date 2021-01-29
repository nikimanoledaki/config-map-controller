package acceptance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = Describe("Acceptance", func() {
	Describe("client", func() {
		It("does sth", func() {
			client := fake.Clientset{}
			Expect(client).NotTo(BeNil())
		})
	})
})
