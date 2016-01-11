package parsec

import (
	"strconv"
	"testing"
)

func TestParsecAnyChar(t *testing.T) {
	input := "abc"
	parser := Concat(AnyChar(), AnyChar())
	result, remaining, ok := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
	if result != "ab" {
		t.Error("nok")
	}
	if remaining != "c" {
		t.Error("nok")
	}
}

func TestParsecDigit(t *testing.T) {
	input := "a1"
	parser := Digit()
	result, remaining, ok := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if ok {
		t.Error("nok")
	}

	input = "1a"
	parser = Digit()
	result, remaining, ok = parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
}

func TestParsecLetter(t *testing.T) {
	input := "a1"
	parser := Letter()
	result, remaining, ok := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}

	input = "1a"
	parser = Letter()
	result, remaining, ok = parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if ok {
		t.Error("nok")
	}
}

func TestParsecLiteral(t *testing.T) {
	input := "1+2+3"
	parser := Literal("1+")
	result, remaining, ok := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
}

func TestParsecAlternative(t *testing.T) {
	input := "abc"
	parser1 := Literal("A")
	parser2 := Literal("a")
	parser := Alternative(parser1, parser2)

	result, remaining, ok := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
}

func TestParsecConcat(t *testing.T) {
	input := "abc"
	parser1 := Literal("a")
	parser2 := Literal("b")
	parser := Concat(parser1, parser2)

	result, remaining, ok := parser(input)
	t.Log("result", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
}

func TestParsecApply(t *testing.T) {
	input := "1+2"

	parser1 := Literal("1")

	toInt := func(str string) interface{} {
		i, _ := strconv.Atoi(str)
		return i
	}

	parser := Map(toInt, parser1)

	result, remaining, ok := parser(input)
	t.Log("result is int", result.(int))
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
}

func TestRepeatZeroOne(t *testing.T) {
	input := "abc"
	parser1 := Literal("a")
	parser := RepeatZeroOne(parser1)
	result, remaining, ok := parser(input)
	t.Log("result is ", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}

	input = "bbc"
	parser1 = Literal("a")
	parser = RepeatZeroOne(parser1)
	result, remaining, ok = parser(input)
	t.Log("result is ", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
}

func TestRepeatZeroMore(t *testing.T) {
	input := "abc"

	parser := RepeatZeroMore(AnyChar())
	result, remaining, ok := parser(input)
	t.Log("result is ", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
	if result != "abc" {
		t.Error("nok")
	}
}

func TestRepeatOneMore(t *testing.T) {
	input := "123a"

	parser := RepeatOneMore(Digit())
	result, remaining, ok := parser(input)
	t.Log("result is ", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
	if result != "123" {
		t.Error("nok")
	}
	if remaining != "a" {
		t.Error("nok")
	}
}

func TestRepeatAmbiguous(t *testing.T) {
	input := "a123"
	parser1 := RepeatZeroMore(AnyChar())
	parser2 := RepeatOneMore(Digit())

	parser := Concat(parser1, parser2)
	result, remaining, ok := parser(input)
	t.Log("result is ", result)
	t.Log("remaining", remaining)
	t.Log("ok", ok)
	if !ok {
		t.Error("nok")
	}
}
