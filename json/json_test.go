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
	if err != nil || !isDoubleEqual(node, v) {
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

func checkString(v string, ctx *Context, t *testing.T) {
	value, err := ctx.ParseString()
	if err != nil || value != v {
		t.Fatal(fmt.Sprintf("err : %v", ctx.json))
	}
}

func TestString(t *testing.T) {
	checkString("", &Context{json: `"" `}, t)
	checkString("Hello", &Context{json: `"Hello" `}, t)
	checkString("Hello\nWorld", &Context{json: `"Hello\nWorld" `}, t)
	checkString("\" \\ / \b \f \n \r \t", &Context{json: `"\" \\ \/ \b \f \n \r \t" `}, t)
	checkString("Hello\u0000World", &Context{json: "\"Hello\\u0000World\" "}, t)
	checkString("\x24", &Context{json: `"\u0024" `}, t)
	checkString("\xE2\x82\xAC", &Context{json: `"\u20AC" `}, t)
	checkString("你好吗", &Context{json: `"你好吗" `}, t)
}

func TestObject(t *testing.T) {

	json := `
{
"n" : null  ,
"f" : false ,
"t" : true,
"i" : 123,
"s" : "abc", 
"a" : [ 1, 2, 3],
"o" : { "1" : 1, "2" : 2, "3" : 3}
}
`
	value, err := Decode(json)
	if value == nil || err != nil {
		t.Fatal("err")
	}

	if value.(map[string]interface{})["n"] != nil {
		t.Fatal("err")
	}

	if value.(map[string]interface{})["f"].(bool) {
		t.Fatal("err")
	}

	if !value.(map[string]interface{})["t"].(bool) {
		t.Fatal("err")
	}

	v := value.(map[string]interface{})["i"].(float64)

	if !isDoubleEqual(v, 123) {
		t.Fatal("err")
	}

	if value.(map[string]interface{})["s"].(string) != "abc" {
		t.Fatal("err")
	}

	a := value.(map[string]interface{})["a"].([]interface{})

	for i := 0; i < 3; i++ {
		if !isDoubleEqual(a[i].(float64), float64(i+1)) {
			t.Fatal("err")
		}
	}

	o := value.(map[string]interface{})["o"].(map[string]interface{})

	if !isDoubleEqual(o["3"].(float64), 3) {
		t.Fatal("err")
	}

	if !isDoubleEqual(o["2"].(float64), 2) {
		t.Fatal("err")
	}

	if !isDoubleEqual(o["1"].(float64), 1) {
		t.Fatal("err")
	}

	fmt.Println(Encode(value))
}

func TestSimpleKind(t *testing.T) {
	json := `null`

	value, err := Decode(json)
	if value != nil || err != nil {
		t.Fatal("err")
	}

	json = `true`

	value, err = Decode(json)
	if value == nil || err != nil {
		t.Fatal("err")
	}

	if !value.(bool) {
		t.Fatal("err")
	}

	json = `"shit"`

	value, err = Decode(json)
	if value == nil || err != nil {
		t.Fatal("err")
	}

	if value.(string) != "shit" {
		t.Fatal("err")
	}

	json = `123456`

	value, err = Decode(json)
	if value == nil || err != nil {
		t.Fatal("err")
	}

	if !isDoubleEqual(value.(float64), 123456) {
		t.Fatal("err")
	}

	json = `[1,2,3]`

	value, err = Decode(json)
	if value == nil || err != nil {
		t.Fatal("err")
	}

	if !isDoubleEqual(value.([]interface{})[0].(float64), 1) {
		t.Fatal("err")
	}
}
