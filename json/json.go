package json

import (
	"errors"
	"strconv"
	"strings"
	"bytes"
	"unicode/utf8"
)

type Kind int

const (
	Null   Kind = 0
	True        = 1
	False       = 2
	String      = 3
	Number      = 4
	Array       = 5
	Object      = 6
)

var (
	Escapee = map[byte]byte{'"': '"', '\\': '\\', '/': '/', 'b': '\b', 'f': '\f', 'n': '\n', 'r': '\r', 't': '\t'}
)

func isDigit1To9(c byte) bool {
	return c >= '1' && c <= '9'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'

}

type Value struct {
	// Null,True,False
	kind Kind

	string_ string

	number float64

	array []*Value

	object map[string]*Value
}

func (value *Value) GetNull() bool {
	return value.kind == Null
}

func (value *Value) GetBool() bool {
	return value.kind == True
}

func (value *Value) GetString() string {
	return value.string_
}

func (value *Value) GetNumber() float64 {
	return value.number
}

func (value *Value) GetArray() []*Value {
	return value.array
}

func (value *Value) GetObject() map[string]*Value {
	return value.object
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
func (ctx *Context) ParseString() (*Value, error) {
	if string_, err := ctx.GetString(); err != nil {
		return nil, err
	} else {
		return &Value{kind: String, string_: string_}, nil
	}
}

// null true false
func (ctx *Context) ParseWord(specifyStr string, kind Kind) (*Value, error) {
	if _, err := ctx.GetWord(specifyStr); err != nil {
		return nil, err
	} else {
		return &Value{kind: kind}, nil
	}
}

// http://www.json.org/number.gif
func (ctx *Context) ParseNumber() (*Value, error) {
	p := 0
	pValid := func() bool { return p < len(ctx.json) }

	if pValid() && ctx.json[p] == '-' {
		p++
	}

	if pValid() && ctx.json[p] == '0' {
		p++
	} else {
		if !pValid() || !isDigit1To9(ctx.json[p]) {
			return nil, errors.New("syntax error")
		}
		p++
		for pValid() && isDigit(ctx.json[p]) {
			p++
		}
	}

	if pValid() && ctx.json[p] == '.' {
		p++
		if !pValid() || !isDigit(ctx.json[p]) {
			return nil, errors.New("syntax error")
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
			return nil, errors.New("syntax error")
		}
		for pValid() && isDigit(ctx.json[p]) {
			p++
		}
	}

	number, err := strconv.ParseFloat(ctx.json[:p], 64)
	if err != nil {
		return nil, errors.New("syntax error")
	}
	ctx.json = ctx.json[p:]
	return &Value{kind: Number, number: number}, nil
}

// http://www.json.org/object.gif
func (ctx *Context) ParseObject() (*Value, error) {
	if err := ctx.RemoveACharacter('{'); err != nil {
		return nil, err
	}

	value := &Value{kind: Object}
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
		if value.object == nil {
			m := make(map[string]*Value)
			value.object = m
		}
		value.object[key] = attribute

		// read ,
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

	return value, nil
}

// http://www.json.org/array.gif
func (ctx *Context) ParseArray() (*Value, error) {
	if err := ctx.RemoveACharacter('['); err != nil {
		return nil, err
	}

	value := &Value{kind: Array}

	for {
		// dispatch
		ele, err := ctx.ParseValue()
		if err != nil {
			return nil, err
		}

		value.array = append(value.array, ele)

		// read ,
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

	return value, nil
}

// http://www.json.org/value.gif
func (ctx *Context) ParseValue() (*Value, error) {
	ctx.RemoveWhite()

	c, err := ctx.PeekACharacter()
	if err != nil {
		return nil, err
	}

	switch c {
	case 'n':
		return ctx.ParseWord("null", Null)
	case 't':
		return ctx.ParseWord("true", True)
	case 'f':
		return ctx.ParseWord("false", False)
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

type Parser struct {
	value *Value
}

//  like: "a" , "a.b" , "a.b.c", ""
func (jp *Parser) Get(key string) *Value {
	if key == "" {
		return jp.value
	}
	// parse
	a := strings.Split(key, ".")
	if jp.value == nil {
		return nil
	}
	value := jp.value
	for _, e := range a {
		if value.object == nil {
			return nil
		}

		if nd, ok := value.object[e]; !ok {
			return nil
		} else {
			value = nd
		}
	}
	return value
}

func ParseJson(json string) (*Parser, error) {
	ctx := &Context{json: json}
	if value, err := ctx.ParseValue(); err != nil {
		return nil, err
	} else {
		return &Parser{value: value}, nil
	}
}
