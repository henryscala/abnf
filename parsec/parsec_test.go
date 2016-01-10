package parsec

import (
	"testing"
)

func TestParsec(t *testing.T) {
	input := "1+2+3"

	var cache []Any

	f := func(a Any) Parser {
		cache = append(cache, a)
		return Char(byte('+'))
	}

	f2 := func(a Any) Parser {
		cache = append(cache, a)
		return Digit()
	}

	parser := Bind(Bind(Digit(), f), f2)

	result := parser(input)

	t.Log("result", result)
	t.Log("cache", cache)

}
