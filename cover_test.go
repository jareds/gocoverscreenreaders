package gocoverscreenreaders

import "testing"

func TestAll(t *testing.T) {
	for i := 1; i < 10; i++ {
		partialCoverage(i)
		fullCoverage(i)
	}
}
