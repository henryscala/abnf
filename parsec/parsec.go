//refer to http://eprints.nottingham.ac.uk/237/1/monparsing.pdf
package parsec

type Any interface{}

type Pair struct {
	First, Second Any
}

type Cell struct {
	Head Any
	Tail string
}

type Parser func(input string) []Cell

//succeed without consuming the input str
func Result(a Any) func(string) []Cell {
	return func(input string) []Cell {
		return []Cell{Cell{Head: a, Tail: input}}
	}
}

//always fail regardless the input str
func Zero(input string) []Cell {
	return []Cell{}
}

//consumes the first char
func Item(input string) []Cell {
	if len(input) == 0 {
		return []Cell{}
	}
	return []Cell{Cell{Head: input[0], Tail: input[1:]}}
}

//apply parser one after another

func Seq(pa, pb Parser) Parser {
	f := func(a Any) Parser {
		return func(input string) []Cell {
			var list []Cell
			list = pb(input)
			if len(list) == 0 {
				return list
			}
			var cell Cell = list[0]

			return []Cell{Cell{Head: Pair{First: a, Second: cell.Head}, Tail: cell.Tail}}
		}
	}
	return Bind(pa, f)

}

//Parser a -> (a->Parser b) -> Parser b
func Bind(pa Parser, f func(a Any) Parser) Parser {
	return func(input string) []Cell {
		var list []Cell
		list = pa(input)
		if len(list) == 0 {
			return list
		}
		var pair Cell
		pair = list[0]

		pb := f(pair.Head)

		var list2 []Cell
		list2 = pb(pair.Tail)
		return list2

	}

}

//Satisfy a byte or char
func Sat(predicate func(byte) bool) Parser {
	return Bind(Item, func(b Any) Parser {
		if predicate(b.(byte)) {
			return Result(b)
		} else {
			return Zero
		}
	})
}

//whether it is a specific char
func Char(c byte) Parser {
	return Sat(func(achar byte) bool {
		return c == achar
	})
}

//whether it is digit
func Digit() Parser {
	return Sat(func(achar byte) bool {
		l := '0'
		u := '9'
		return rune(achar) >= l && rune(achar) <= u
	})
}

//whether it is a lower char
func Lower() Parser {
	return Sat(
		func(achar byte) bool {
			l := 'a'
			u := 'z'
			return rune(achar) >= l && rune(achar) <= u
		})
}

//whether it is a lower char
func Upper() Parser {
	return Sat(
		func(achar byte) bool {
			l := 'A'
			u := 'Z'
			return rune(achar) >= l && rune(achar) <= u
		})
}

func Plus(p1, p2 Parser) Parser {
	return func(input string) []Cell {
		list1 := p1(input)
		list2 := p2(input)

		if len(list1) == 0 {
			if len(list2) == 0 {
				return []Cell{}
			} else {
				return list2
			}
		} else {
			if len(list2) == 0 {
				return list1
			} else {
				return append(list1, list2...)
			}
		}

	}
}

func Letter() Parser {
	return Plus(Lower(), Upper())
}

func AlphaNum() Parser {
	return Plus(Letter(), Digit())
}
