package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Does not work on amd64 architecture")
	}
	t.Fail()
}
