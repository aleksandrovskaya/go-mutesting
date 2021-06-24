package numbers

import (
	"testing"

	"github.com/vasiliyyudin/go-mutesting/test"
)

func TestMutatorNumbersIncrementer(t *testing.T) {
	test.Mutator(
		t,
		MutatorNumbersIncrementer,
		"../../testdata/numbers/incrementer.go",
		2,
	)
}
