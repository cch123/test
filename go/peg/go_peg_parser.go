package main

import (
	"fmt"
	"strconv"

	gopeg "github.com/yhirose/go-peg"
)

func main() {

	// Create a PEG parser
	parser, _ := gopeg.NewParser(`
    # Simple calculator
    EXPR         ←  ATOM (BINOP ATOM)*
    ATOM         ←  NUMBER / '(' EXPR ')'
    BINOP        ←  < [-+/*] >
    NUMBER       ←  < [0-9]+ >
    %whitespace  ←  [ \t]*
    ---
    # Expression parsing option
    %expr  = EXPR   # Rule to apply 'precedence climbing method' to
    %binop = L + -  # Precedence level 1
    %binop = L * /  # Precedence level 2
	`)

	// Setup semantic actions
	g := parser.Grammar
	g["EXPR"].Action = func(v *gopeg.Values, d gopeg.Any) (gopeg.Any, error) {
		val := v.ToInt(0)
		if v.Len() > 1 {
			ope := v.ToStr(1)
			rhs := v.ToInt(2)
			switch ope {
			case "+":
				val += rhs
			case "-":
				val -= rhs
			case "*":
				val *= rhs
			case "/":
				val /= rhs
			}
		}
		return val, nil
	}
	g["BINOP"].Action = func(v *gopeg.Values, d gopeg.Any) (gopeg.Any, error) {
		return v.Token(), nil
	}
	g["NUMBER"].Action = func(v *gopeg.Values, d gopeg.Any) (gopeg.Any, error) {
		return strconv.Atoi(v.Token())
	}

	// Parse
	//input := " 1 + 2 * 3 * (4 - 5 + 6) / 7 - 8 "
	input := " 1+ 2 * 3 / 2 + (1+ 4) * 5"
	val, _ := parser.ParseAndGetValue(input, nil)

	fmt.Println(val) // Output: -3
}
