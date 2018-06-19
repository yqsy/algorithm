package json

import (
	"errors"
	"strconv"
	"bytes"
	"unicode/utf8"
	"fmt"
)

var (
	Escapee = map[byte]byte{'"': '"', '\\': '\\', '/': '/', 'b': '\b', 'f': '\f', 'n': '\n', 'r': '\r', 't': '\t'}
	BoolMap = map[bool]string{true: "true", false: "false"}
)

func isDigit1To9(c byte) bool {
	return c >= '1' && c <= '9'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'

}

func Encode(v interface{}) string {
	switch v.(type) {
	case nil:
		return "null"
	case bool:
		return BoolMap[v.(bool)]
	case string:
		return `"` + v.(string) + `"`
	case float64:
		return fmt.Sprintf("%v", v.(float64))
	case []interface{}:
		a := v.([]interface{})
		prettify := "["
		for i := 0; i < len(a); i++ {
			prettify += Encode(a[i]) + ","
		}
		if len(prettify) > 0 && prettify[len(prettify)-1] == ',' {
			prettify = prettify[:len(prettify)-1]
		}
		return prettify + "]"
	case map[string]interface{}:
		o := v.(map[string]interface{})
		prettify := "{"
		for k, v := range o {
			prettify += fmt.Sprintf(`"%v": %v,`, k, Encode(v))
		}
		if len(prettify) > 0 && prettify[len(prettify)-1] == ',' {
			prettify = prettify[:len(prettify)-1]
		}
		return prettify + "}"
	default:
		panic("only support nil,bool,string,float64,[]interface{} and map[string]interface{}")
	}
}

func Decode(json string) (interface{}, error) {
	ctx := Context{json: json}
	return ctx.ParseValue()
}

type Context struct {
	json string
}

func (ctx *Context) RemoveWhite() {
	i := 0
	for i < len(ctx.json) && ctx.json[i] <= ' ' {
		i++
	}
	ctx.json = ctx.json[i:]
}

func (ctx *Context) RemoveACharacter(c byte) error {
	if len(ctx.json) < 1 || ctx.json[0] != c {
		return errors.New("syntax error")
	}
	ctx.json = ctx.json[1:]
	return nil
}

func (ctx *Context) PeekACharacter() (byte, error) {
	if len(ctx.json) < 1 {
		return '0', errors.New("syntax error")
	}
	return ctx.json[0], nil
}

// get string from ""
func (ctx *Context) GetString() (string, error) {
	if err := ctx.RemoveACharacter('"'); err != nil {
		return "", err
	}

	var buf bytes.Buffer
	for i := 0; i < len(ctx.json); i++ {
		if ctx.json[i] == '"' {
			ctx.json = ctx.json[i+1:]
			return buf.String(), nil
		}

		if ctx.json[i] == '\\' {
			i++

			if i >= len(ctx.json) {
				break
			}

			if ele, ok := Escapee[ctx.json[i]]; ok {
				buf.Write([]byte{ele})
			} else if ctx.json[i] == 'u' {
				if !(i+4 < len(ctx.json)) {
					break
				}
				r, err := strconv.ParseUint(ctx.json[i+1:i+5], 16, 64)
				if err != nil {
					break
				}
				tmp := make([]byte, 8)
				n := utf8.EncodeRune(tmp, rune(r))
				buf.Write(tmp[:n])
				i += 4
			} else {
				break
			}
		} else {
			buf.Write([]byte{ctx.json[i]})
		}
	}

	return "", errors.New("syntax error")
}

// get specify word
func (ctx *Context) GetWord(s string) (string, error) {
	if len(ctx.json) < len(s) {
		return "", errors.New("syntax error")
	}
	rtn := ctx.json[:len(s)]
	if rtn != s {
		return "", errors.New("syntax error")
	}
	ctx.json = ctx.json[len(s):]
	return rtn, nil
}

// http://www.json.org/string.gif
func (ctx *Context) ParseString() (string, error) {
	return ctx.GetString()
}

// null true false
func (ctx *Context) ParseWord(specifyStr string, i interface{}) (interface{}, error) {
	if _, err := ctx.GetWord(specifyStr); err != nil {
		return nil, err
	} else {
		return i, nil
	}
}

// http://www.json.org/number.gif
func (ctx *Context) ParseNumber() (float64, error) {
	p := 0
	pValid := func() bool { return p < len(ctx.json) }

	if pValid() && ctx.json[p] == '-' {
		p++
	}

	if pValid() && ctx.json[p] == '0' {
		p++
	} else {
		if !pValid() || !isDigit1To9(ctx.json[p]) {
			return 0.0, errors.New("syntax error")
		}
		p++
		for pValid() && isDigit(ctx.json[p]) {
			p++
		}
	}

	if pValid() && ctx.json[p] == '.' {
		p++
		if !pValid() || !isDigit(ctx.json[p]) {
			return 0.0, errors.New("syntax error")
		}
		for pValid() && isDigit(ctx.json[p]) {
			p++
		}
	}

	if pValid() && (ctx.json[p] == 'e' || ctx.json[p] == 'E') {
		p++
		if pValid() && (ctx.json[p] == '+' || ctx.json[p] == '-') {
			p++
		}
		if !pValid() || !isDigit(ctx.json[p]) {
			return 0.0, errors.New("syntax error")
		}
		for pValid() && isDigit(ctx.json[p]) {
			p++
		}
	}

	number, err := strconv.ParseFloat(ctx.json[:p], 64)
	if err != nil {
		return 0.0, errors.New("syntax error")
	}
	ctx.json = ctx.json[p:]
	return number, nil
}

// http://www.json.org/array.gif
func (ctx *Context) ParseArray() ([]interface{}, error) {
	if err := ctx.RemoveACharacter('['); err != nil {
		return nil, err
	}

	a := make([]interface{}, 0)

	for {
		// dispatch
		ele, err := ctx.ParseValue()
		if err != nil {
			return nil, err
		}

		a = append(a, ele)

		// read , represent have another ele
		ctx.RemoveWhite()
		if err = ctx.RemoveACharacter(','); err != nil {
			break
		}
	}

	// read ]
	ctx.RemoveWhite()
	if err := ctx.RemoveACharacter(']'); err != nil {
		return nil, err
	}

	return a, nil
}

// http://www.json.org/object.gif
func (ctx *Context) ParseObject() (map[string]interface{}, error) {
	if err := ctx.RemoveACharacter('{'); err != nil {
		return nil, err
	}

	o := make(map[string]interface{})
	for {
		// read key
		ctx.RemoveWhite()
		key, err := ctx.GetString()
		if err != nil {
			return nil, err
		}

		// read colon
		ctx.RemoveWhite()
		if err = ctx.RemoveACharacter(':'); err != nil {
			return nil, err
		}

		// dispatch
		attribute, err := ctx.ParseValue()
		if err != nil {
			return nil, err
		}

		// save to map
		o[key] = attribute

		// read , represent have another ele
		ctx.RemoveWhite()
		if err = ctx.RemoveACharacter(','); err != nil {
			break
		}
	}

	// read }
	ctx.RemoveWhite()
	if err := ctx.RemoveACharacter('}'); err != nil {
		return nil, err
	}

	return o, nil
}

// http://www.json.org/value.gif
func (ctx *Context) ParseValue() (interface{}, error) {
	ctx.RemoveWhite()

	c, err := ctx.PeekACharacter()
	if err != nil {
		return nil, err
	}

	switch c {
	case 'n':
		return ctx.ParseWord("null", nil)
	case 't':
		return ctx.ParseWord("true", true)
	case 'f':
		return ctx.ParseWord("false", false)
	case '"':
		return ctx.ParseString()
	case '[':
		return ctx.ParseArray()
	case '{':
		return ctx.ParseObject()
	default:
		return ctx.ParseNumber()
	}
}
