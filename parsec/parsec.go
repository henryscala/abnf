//refer to http://theorangeduck.com/page/you-could-have-invented-parser-combinators
package parsec

import (
	"strings"
)

var EmptyStringSlice []string = []string{}
var EmptyAnySlice []interface{} = []interface{}{}

type Parser func(input string) (result []string, remaining string)

type AppliedParser func(input string) (result []interface{}, remaining string)

type Apply func(inbound string) interface{}

func Literal(s string) Parser {
	return func(input string) (result []string, remaining string) {

		if !strings.HasPrefix(input, s) {
			return EmptyStringSlice, input
		}
		result = append(result, s)
		remaining = input[len(s):]
		return
	}
}

func Alternative(p1, p2 Parser) Parser {
	return func(input string) (result []string, remaining string) {
		result, remaining = p1(input)
		if len(result) != 0 {
			return
		}
		result, remaining = p2(input)
		return
	}
}

func Concat(p1, p2 Parser) Parser {
	return func(input string) (result []string, remaining string) {
		result, remaining = p1(input)
		if len(result) == 0 { //failed
			return EmptyStringSlice, input
		}
		result2, remaining2 := p2(remaining)
		if len(result2) == 0 {
			return EmptyStringSlice, input
		}

		result[0] += result2[0]
		return result, remaining2
	}
}

func Map(f Apply, p Parser) AppliedParser {
	return func(input string) (result []interface{}, remaining string) {
		var result1 []string
		result1, remaining = p(input)
		if len(result1) == 0 {
			return EmptyAnySlice, input
		}
		result = append(result, f(result1[0]))
		return result, remaining
	}
}
