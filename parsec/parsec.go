//refer to http://theorangeduck.com/page/you-could-have-invented-parser-combinators
package parsec

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Parser func(input string) (result string, remaining string, ok bool)

type AppliedParser func(input string) (result interface{}, remaining string, ok bool)

type Apply func(inbound string) interface{}

func Success() Parser {
	return func(input string) (result string, remaining string, ok bool) {
		return "", input, true
	}
}

func Fail() Parser {
	return func(input string) (result string, remaining string, ok bool) {
		return "", input, false
	}
}

func Literal(s string) Parser {
	return func(input string) (result string, remaining string, ok bool) {

		if !strings.HasPrefix(input, s) {
			return "", input, false
		}
		result = s
		remaining = input[len(s):]
		ok = true
		return
	}
}

func Set(f func(string) bool) Parser {
	return func(input string) (result string, remaining string, ok bool) {
		r, size := utf8.DecodeRuneInString(input)
		if r == utf8.RuneError {
			return "", input, false
		}
		if !f(string(r)) {
			return "", input, false
		}
		return string(r), input[size:], true
	}
}

func isDigit(s string) bool {
	if len(s) != 1 {
		panic("not exptected")
	}
	if s[0] >= '0' && s[0] <= '9' {
		return true
	}
	return false
}

func isLower(s string) bool {
	if len(s) < 1 {
		panic("not exptected")
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return true
	}
	return false
}

func isUpper(s string) bool {
	if len(s) < 1 {
		panic("not exptected")
	}
	if s[0] >= 'A' && s[0] <= 'Z' {
		return true
	}
	return false
}

func UpperLetter() Parser {
	return Set(isUpper)
}

func LowerLetter() Parser {
	return Set(isLower)
}

func Letter() Parser {
	return Alternative(LowerLetter(), UpperLetter())
}

func Digit() Parser {
	return Set(isDigit)
}

func AnyChar() Parser {
	return Set(func(string) bool {
		return true //always
	})
}

func Alternative(ps ...Parser) Parser {
	p := Fail()
	for _, p1 := range ps {

		p = alternative2(p, p1)
	}
	return p

}

func alternative2(p1, p2 Parser) Parser {
	return func(input string) (result string, remaining string, ok bool) {
		result, remaining, ok = p1(input)
		if ok { //succeed
			return
		}
		result, remaining, ok = p2(input)
		return
	}
}

func Concat(ps ...Parser) Parser {
	p := Success()
	for _, p1 := range ps {
		p = concat2(p, p1)
	}
	return p
}

func concat2(p1, p2 Parser) Parser {
	return func(input string) (result string, remaining string, ok bool) {
		result, remaining, ok = p1(input)
		if !ok { //failed
			return "", input, false
		}
		result2, remaining2, ok2 := p2(remaining)
		if !ok2 {
			return "", input, false
		}

		return result + result2, remaining2, true
	}
}

func Map(f Apply, p Parser) AppliedParser {
	return func(input string) (result interface{}, remaining string, ok bool) {
		var result1 string
		result1, remaining, ok = p(input)
		if !ok {
			return "", input, false
		}
		result = f(result1)

		return result, remaining, true
	}
}

//a?
func RepeatZeroOne(p Parser) Parser {
	return Repeat(p, 0, 1)
}

//a+
func RepeatOneMore(p Parser) Parser {
	return Repeat(p, 1, -1)
}

//a*
func RepeatZeroMore(p Parser) Parser {
	return Repeat(p, 0, -1)
}

func Repeat(p Parser, min, max int) Parser {
	if max >= 0 {
		if max < min {
			panic("not expected")
		}
	} else {
		// max < 0 means infinity
	}
	return func(input string) (result string, remaining string, ok bool) {
		ok = true
		result1, remaining1, ok1 := result, remaining, ok
		input1 := input
		for i := 0; i < min; i++ {
			result1, remaining1, ok1 = p(input1)
			if !ok1 {
				fmt.Println("Repeat1-", min, max, "", input, false)
				return "", input, false
			}
			input1 = remaining1
			result += result1
		}

		for i := min; i < max; i++ {
			result1, remaining1, ok1 = p(input1)
			if !ok1 {
				fmt.Println("Repeat2-", min, max, result, input1, true)
				return result, input1, true //return true up till now
			}
			input1 = remaining1
			result += result1
		}

		if max >= 0 {
			remaining = remaining1
			fmt.Println("Repeat3-", min, max, result, remaining, true)
			return result, remaining, true //return true up till now
		}

		//max < 0 means infinity
		for i := min; true; i++ {
			if len(input1) == 0 {
				break
			}
			result1, remaining1, ok1 = p(input1)
			fmt.Println("Repat3.5-", "result1", result1, "remaining1", remaining1, "ok1", ok1)
			if !ok1 {
				fmt.Println("Repeat4-", min, max, result, input1, true)
				return result, input1, true //return true up till now
			}
			input1 = remaining1
			result += result1
		}

		fmt.Println("Repeat5-", min, max, result, input1, true)
		return result, input1, true
	}
}
