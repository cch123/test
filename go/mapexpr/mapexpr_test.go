package mapexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	m      map[string]string
	expr   string
	result bool
}

func TestMapExpr(t *testing.T) {
	cases := []testCase{
		testCase{
			m: map[string]string{
				"age": "13", "name": "14",
			},
			expr:   `age == 13 && name == 14`,
			result: true,
		},
		testCase{
			m: map[string]string{
				"age": "3", "name": "14",
			},
			expr:   `age == 13 && name == 14`,
			result: false,
		},
	}

	for _, cas := range cases {
		res, err := Eval(cas.m, cas.expr)
		assert.Nil(t, err)
		assert.Equal(t, res, cas.result)
	}
}
