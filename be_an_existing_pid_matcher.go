package matchermaker

import (
	"fmt"

	"github.com/onsi/gomega/types"
	"golang.org/x/sys/unix"
)

func BeAnExistingPid() types.GomegaMatcher {
	return &BeAnExistingPidMatcher{}
}

type BeAnExistingPidMatcher struct{}

func (matcher *BeAnExistingPidMatcher) Match(actual interface{}) (bool, error) {
	pid, ok := actual.(int)
	if !ok {
		return false, fmt.Errorf("BeAnExistingPid matcher expects an int, got %#v", actual)
	}

	err := unix.Kill(pid, unix.Signal(0))
	return (err == nil), nil
}

func (matcher *BeAnExistingPidMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected pid %#v to be an existing process", actual)
}

func (matcher *BeAnExistingPidMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected pid %#v not to be an existing process", actual)
}
