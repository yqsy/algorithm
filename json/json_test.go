package json

import (
	"testing"
	"math"
	"fmt"
)

func isDoubleEqual(f1, f2 float64) bool {
	const TOLERANCE = 0.000001
	return math.Abs(f1-f2) < TOLERANCE
}

func checkNumber(v float64, ctx *Context, t *testing.T) {
	node, err := ctx.ParseNumber()
	if err != nil || !isDoubleEqual(node.GetNumber(), v) {
		t.Fatal(fmt.Sprintf("err : %v", ctx.json))
	}
}

func TestNumber(t *testing.T) {
	checkNumber(1.0, &Context{json: "1 "}, t)
	checkNumber(0.0, &Context{json: "-0 "}, t)
	checkNumber(0.0, &Context{json: "-0.0 "}, t)
	checkNumber(1, &Context{json: "1 "}, t)

	checkNumber(-1.0, &Context{json: "-1 "}, t)
	checkNumber(1.5, &Context{json: "1.5 "}, t)
	checkNumber(-1.5, &Context{json: "-1.5 "}, t)
	checkNumber(3.1416, &Context{json: "3.1416 "}, t)
	checkNumber(1E10, &Context{json: "1E10 "}, t)
	checkNumber(1e10, &Context{json: "1e10 "}, t)
	checkNumber(1E+10, &Context{json: "1E+10 "}, t)
	checkNumber(1E-10, &Context{json: "1E-10 "}, t)
	checkNumber(-1E10, &Context{json: "-1E10 "}, t)
	checkNumber(-1e10, &Context{json: "-1e10 "}, t)
	checkNumber(-1E+10, &Context{json: "-1E+10 "}, t)
	checkNumber(-1E-10, &Context{json: "-1E-10 "}, t)
	checkNumber(1.234E+10, &Context{json: "1.234E+10 "}, t)
	checkNumber(1.234E-10, &Context{json: "1.234E-10 "}, t)
	checkNumber(0.0, &Context{json: "1e-10000 "}, t) /* must underflow */

	checkNumber(1.0000000000000002, &Context{json: "1.0000000000000002 "}, t)           /* the smallest number > 1 */
	checkNumber(4.9406564584124654e-324, &Context{json: "4.9406564584124654e-324 "}, t) /* minimum denormal */
	checkNumber(-4.9406564584124654e-324, &Context{json: "-4.9406564584124654e-324 "}, t)
	checkNumber(2.2250738585072009e-308, &Context{json: "2.2250738585072009e-308 "}, t) /* Max subnormal double */
	checkNumber(-2.2250738585072009e-308, &Context{json: "-2.2250738585072009e-308 "}, t)
	checkNumber(2.2250738585072014e-308, &Context{json: "2.2250738585072014e-308 "}, t) /* Min normal positive double */
	checkNumber(-2.2250738585072014e-308, &Context{json: "-2.2250738585072014e-308 "}, t)
	checkNumber(1.7976931348623157e+308, &Context{json: "1.7976931348623157e+308 "}, t) /* Max double */
	checkNumber(-1.7976931348623157e+308, &Context{json: "-1.7976931348623157e+308 "}, t)

}
