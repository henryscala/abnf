package parsec

import (
	"strconv"
	"testing"
)

func TestParsecLiteral(t *testing.T) {
	input := "1+2+3"
	parser := Literal("1+")
	result, remaining := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)

}

func TestParsecAlternative(t *testing.T) {
	input := "abc"
	parser1 := Literal("A")
	parser2 := Literal("a")
	parser := Alternative(parser1, parser2)

	result, remaining := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
}

func TestParsecConcat(t *testing.T) {
	input := "abc"
	parser1 := Literal("a")
	parser2 := Literal("b")
	parser := Concat(parser1, parser2)

	result, remaining := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
}

func TestParsecApply(t *testing.T) {
	input := "1+2"

	parser1 := Literal("1")

	toInt := func(str string) interface{} {
		i, _ := strconv.Atoi(str)
		return i
	}

	parser := Map(toInt, parser1)

	result, remaining := parser(input)
	t.Log("result is int", result[0].(int))
	t.Log("remaining", remaining)
}
