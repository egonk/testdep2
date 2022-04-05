package b

import "github.com/egonk/testdep1"

func B() {
	return testdep1.A() + "b"
}
