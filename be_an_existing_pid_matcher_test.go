package matchermaker_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/williammartin/matchermaker"
)

var _ = Describe("BeAnExistingPid", func() {
	var pid int

	When("the pid does not exist", func() {
		BeforeEach(func() {
			pid = 1000000000
		})

		It("should fail", func() {
			Expect(pid).NotTo(BeAnExistingPid())
		})
	})

	When("the pid does exist", func() {
		BeforeEach(func() {
			pid = os.Getpid()
		})

		It("should succeed", func() {
			Expect(pid).To(BeAnExistingPid())
		})
	})

	When("not passed an int", func() {
		It("errors", func() {
			_, err := BeAnExistingPid().Match(nil)
			Expect(err).To(MatchError("BeAnExistingPid matcher expects an int, got <nil>"))
		})
	})

	It("should provide a useful failure message", func() {
		Expect(BeAnExistingPid().FailureMessage(1)).To(Equal("Expected pid 1 to be an existing process"))
	})

	It("should provide a useful negated failure message", func() {
		Expect(BeAnExistingPid().NegatedFailureMessage(1)).To(Equal("Expected pid 1 not to be an existing process"))
	})
})
