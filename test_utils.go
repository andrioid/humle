package humle

import (
	"strings"
	"testing"
)

type TestCase struct {
	Input    Node
	Expected string
}

type TestCases map[string]TestCase

func (testcases TestCases) Test(t *testing.T) {

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			b := strings.Builder{}
			WriteHTML(&b, tc.Input)
			if b.String() != tc.Expected {
				t.Errorf("💩 %v\n\033[32m%#v\033[0m\n\033[31m%#v\033[0m", name, tc.Expected, b.String())
			}
		})
	}
}
