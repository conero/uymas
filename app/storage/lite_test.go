package storage

import (
	"testing"
)

func TestLiteralExpression(t *testing.T) {
	val, er := LiteralExpression("3*(5-4*0.23+(10/2))")
	t.Log(val, er)
}

func TestExpNoBracket(t *testing.T) {
	var result, exp string

	exp = `10/2`
	result = ExpNoBracket(exp)
	t.Log(exp, result)

	exp = `3.45*2.14\
`
	result = ExpNoBracket(exp)
	t.Log(exp, result)
}
