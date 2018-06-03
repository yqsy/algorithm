package json

import (
	"github.com/golang-collections/collections/stack"
	"errors"
	"fmt"
	"strconv"
	"bytes"
	"unicode/utf8"
	"strings"
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

func isDigit1To9(c byte) bool {
	if c >= '1' && c <= '9' {
		return true
	}

	return false
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

type Value struct {
	// 类型,Null,True,False
	kind Kind

	// 字符串
	string_ *string

	// 数字
	number *float64

	// 数组
	array *[]*Value

	// 对象
	object *map[string]*Value
}

func (value *Value) GetNull() bool {
	if value.kind == Null {
		return true
	} else {
		return false
	}
}

func (value *Value) GetBool() bool {
	if value.kind == True {
		return true
	} else {
		return false
	}
}

func (value *Value) GetString() string {
	if value.kind == String || value.string_ != nil {
		return *value.string_
	} else {
		return ""
	}
}

func (value *Value) GetNumber() float64 {
	if value.kind == Number {
		return *value.number
	} else {
		return 0.0
	}
}

func (value *Value) GetArray() *[]*Value {
	if value.kind == Array {
		return value.array
	} else {
		return nil
	}
}

func (value *Value) GetObject() *map[string]*Value {
	if value.kind == Object {
		return value.object
	} else {
		return nil
	}
}

type Context struct {
	json string

	// 用来做符号对称匹配
	stack stack.Stack
}

// slice移动到下一个非空白处
func (ctx *Context) RemoveWhite() {
	i := 0
	for ; i < len(ctx.json); i++ {
		b := ctx.json[i]
		if b != ' ' && b != '\t' && b != '\r' && b != '\n' {
			break
		}
	}

	if i < len(ctx.json) {
		ctx.json = ctx.json[i:]
	}
}

// remove一个字符
func (ctx *Context) RemoveACharacter(c byte) error {
	if len(ctx.json) < 2 || ctx.json[0] != c {
		return errors.New(fmt.Sprintf("err %v", c))
	}

	ctx.json = ctx.json[1:]
	return nil
}

// peek一个字符
func (ctx *Context) PeekACharacter() (byte, error) {
	if len(ctx.json) < 1 {
		return '0', errors.New("too short")
	}
	return ctx.json[0], nil
}

// 获取""中的字符串,并把slice指向到第二个"的后面的idx
func (ctx *Context) GetString() (string, error) {
	if len(ctx.json) < 2 {
		return "", errors.New("get str error")
	}

	first := -1
	second := -1
	for i := 0; i < len(ctx.json); i++ {
		if ctx.json[i] == '"' {
			if first == -1 {
				first = i
			} else {
				if ctx.json[i-1] != '\\' {
					second = i
					break
				}
			}
		}
	}

	if first == -1 || second == -1 {
		return "", errors.New("get str error")
	}

	// decode utf8
	var buf bytes.Buffer

	for i := first + 1; i < second; {
		if ctx.json[i] == '\\' {
			i++
			if i >= second {
				return "", errors.New("too short")
			}

			switch ctx.json[i] {
			case '"':
				buf.Write([]byte{'"'})
				i++
			case '\\':
				buf.Write([]byte{'\\'})
				i++
			case '/':
				buf.Write([]byte{'/'})
				i++
			case 'b':
				buf.Write([]byte{'\b'})
				i++
			case 'f':
				buf.Write([]byte{'\f'})
				i++
			case 'n':
				buf.Write([]byte{'\n'})
				i++
			case 'r':
				buf.Write([]byte{'\r'})
				i++
			case 't':
				buf.Write([]byte{'\t'})
				i++
			case 'u':
				i++
				if i+4 > second {
					return "", errors.New("too short")
				}

				r, err := strconv.ParseUint(ctx.json[i:i+4], 16, 32)

				if err != nil {
					return "", err
				}

				utf8Buf := make([]byte, 8)
				n := utf8.EncodeRune(utf8Buf, rune(r))
				if n < 1 {
					return "", errors.New("utf8 parse error")
				}
				buf.Write(utf8Buf[:n])

				i += 4
			default:
				return "", errors.New("error transference symbol")
			}
		} else {
			buf.Write([]byte{ctx.json[i]})
			i++
		}
	}

	ctx.json = ctx.json[second+1:]
	return buf.String(), nil
}

// 获取指定的字符串,并把slice指向字符串后面的idx
func (ctx *Context) GetWord(s string) (string, error) {
	if len(ctx.json) < len(s) {
		return "", errors.New("get str error")
	}
	rtn := ctx.json[:len(s)]

	if rtn != s {
		return "", errors.New("get str error")
	}

	ctx.json = ctx.json[len(s):]
	return rtn, nil
}

func (ctx *Context) ParseString() (*Value, error) {
	string_, err := ctx.GetString()
	if err != nil {
		return nil, err
	}

	value := &Value{}
	value.kind = String
	value.string_ = &string_
	return value, nil
}

func (ctx *Context) ParseWord(specifyStr string, kind Kind) (*Value, error) {
	_, err := ctx.GetWord(specifyStr)

	if err != nil {
		return nil, err
	}

	value := &Value{}
	value.kind = kind
	return value, nil
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
			return nil, errors.New("parse number error")
		}

		p++

		for ; pValid() && isDigit(ctx.json[p]); p++ {

		}
	}

	if pValid() && ctx.json[p] == '.' {
		p++

		if !pValid() || !isDigit(ctx.json[p]) {
			return nil, errors.New("parse number error")
		}

		p++

		for ; pValid() && isDigit(ctx.json[p]); p++ {

		}
	}

	if pValid() && (ctx.json[p] == 'e' || ctx.json[p] == 'E') {
		p++

		if pValid() && (ctx.json[p] == '+' || ctx.json[p] == '-') {
			p++
		}

		if p >= len(ctx.json) || !isDigit(ctx.json[p]) {
			return nil, errors.New("parse number error")
		}

		for ; pValid() && isDigit(ctx.json[p]); p++ {

		}
	}

	number, err := strconv.ParseFloat(ctx.json[:p], 64)
	if err != nil {
		return nil, errors.New("parse number error")
	}

	ctx.json = ctx.json[p:]

	value := &Value{}
	value.kind = Number
	value.number = &number
	return value, nil
}

// http://www.json.org/object.gif
func (ctx *Context) ParseObject() (*Value, error) {
	err := ctx.RemoveACharacter('{')
	if err != nil {
		return nil, err
	}
	ctx.stack.Push(int32('}'))

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
		err = ctx.RemoveACharacter(':')
		if err != nil {
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
			value.object = &m
		}
		(*value.object)[key] = attribute

		// read ,
		ctx.RemoveWhite()
		err = ctx.RemoveACharacter(',')
		if err != nil {
			break
		}
	}

	// read }
	ctx.RemoveWhite()
	err = ctx.RemoveACharacter('}')
	if err != nil {
		return nil, err
	}

	if ctx.stack.Pop().(int32) != '}' {
		return nil, errors.New("not match")
	}

	return value, nil
}

// http://www.json.org/array.gif
func (ctx *Context) ParseArray() (*Value, error) {
	err := ctx.RemoveACharacter('[')
	if err != nil {
		return nil, err
	}
	ctx.stack.Push(int32(']'))

	value := &Value{kind: Array}

	for {
		// dispatch
		ele, err := ctx.ParseValue()
		if err != nil {
			return nil, err
		}

		// save to array
		if value.array == nil {
			a := make([]*Value, 0)
			value.array = &a
		}

		*value.array = append(*value.array, ele)

		// read ,
		ctx.RemoveWhite()
		err = ctx.RemoveACharacter(',')
		if err != nil {
			break
		}
	}

	// read ]
	ctx.RemoveWhite()
	err = ctx.RemoveACharacter(']')
	if err != nil {
		return nil, err
	}

	if ctx.stack.Pop().(int32) != ']' {
		return nil, errors.New("not match")
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
	case '"':
		return ctx.ParseString()
	case 'n':
		return ctx.ParseWord("null", Null)
	case 't':
		return ctx.ParseWord("true", True)
	case 'f':
		return ctx.ParseWord("false", False)
	case '[':
		return ctx.ParseArray()
	case '{':
		return ctx.ParseObject()
	default:
		return ctx.ParseNumber()
	}

	return nil, errors.New("unknown value")
}

type Parser struct {
	value *Value
}

//  a , a.b , a.b.c  . 是对象的属性
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

		nd, ok := (*value.object)[e]

		if !ok {
			return nil
		}

		value = nd
	}

	return value
}

func ParseJson(json string) (*Parser, error) {
	ctx := &Context{json: json}
	value, err := ctx.ParseValue()
	if err != nil {
		return nil, err
	}
	jp := &Parser{value: value}
	return jp, nil
}
