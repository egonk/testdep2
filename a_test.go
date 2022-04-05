package testdep2

import (
	"testing"

	"github.com/egonk/testdep1"
)

func TestA(t *testing.T) {
	str := testdep1.A() + "," + A()
	if str != "testdep1,testdep2" {
		t.Error(str)
	}
}
