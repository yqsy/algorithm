package json

import (
	"github.com/golang-collections/collections/stack"
	"errors"
	"fmt"
	"strconv"
)

type Kind int

const (
	Null   Kind = 0
	True
	False
	String
	Number
	Array
	Object
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

type Node struct {
	// 类型,Null,True,False
	kind Kind

	// 字符串
	value *string

	// 数字
	number *float64

	// 数组
	array *[]*Node

	// 对象
	object *map[string]*Node
}

func (node *Node) GetString() string {
	if node.kind == String || node.value != nil {
		return *node.value
	} else {
		return ""
	}
}

func (node *Node) GetBool() bool {
	if node.kind == True {
		return true
	} else {
		return false
	}
}

func (node *Node) GetNull() bool {
	if node.kind == Null {
		return true
	} else {
		return false
	}
}

func (node *Node) GetNumber() float64 {
	if node.kind == Number {
		return *node.number
	} else {
		return 0.0
	}
}

type JsonParse struct {
	// 首节点
	node *Node
}

func (jp *JsonParse) Get(key string) *Node {
	return jp.node
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

// 去一个字符
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
				second = i
				break
			}
		}
	}

	if first == -1 || second == -1 {
		return "", errors.New("get str error")
	}
	key := ctx.json[first+1 : second]

	// 第二个超过范围就返回
	if second+1 >= len(ctx.json) {
		return "", errors.New("too short")
	}

	ctx.json = ctx.json[second+1:]
	return key, nil
}

func (ctx *Context) GetSpecifyString(s string) (string, error) {
	if len(ctx.json) <= len(s) {
		return "", errors.New("get str error")
	}
	rtn := ctx.json[:len(s)]

	if rtn != s {
		return "", errors.New("get str error")
	}

	ctx.json = ctx.json[len(s):]
	return rtn, nil
}

func (ctx *Context) ParseValue() (*Node, error) {
	value, err := ctx.GetString()
	if err != nil {
		return nil, err
	}

	node := &Node{}
	node.kind = String
	node.value = &value
	return node, nil
}

func (ctx *Context) ParseStr(specifyStr string, kind Kind) (*Node, error) {
	_, err := ctx.GetSpecifyString(specifyStr)

	if err != nil {
		return nil, err
	}

	node := &Node{}
	node.kind = kind
	return node, nil
}

// http://www.json.org/number.gif
func (ctx *Context) ParseNumber() (*Node, error) {
	p := 0

	if p < len(ctx.json) && ctx.json[p] == '-' {
		p++
	}

	if p < len(ctx.json) && ctx.json[p] == '0' {
		p++
	} else {
		if p >= len(ctx.json) || !isDigit1To9(ctx.json[p]) {
			return nil, errors.New("parse number error")
		}

		p++

		for ; p < len(ctx.json) && isDigit(ctx.json[p]); p++ {

		}
	}

	if p < len(ctx.json) && ctx.json[p] == '.' {
		p++

		if p >= len(ctx.json) || !isDigit(ctx.json[p]) {
			return nil, errors.New("parse number error")
		}

		p++

		for ; p < len(ctx.json) && isDigit(ctx.json[p]); p++ {

		}
	}

	if p < len(ctx.json) && (ctx.json[p] == 'e' || ctx.json[p] == 'E') {
		p++

		if p < len(ctx.json) && (ctx.json[p] == '+' || ctx.json[p] == '-') {
			p++
		}

		if p >= len(ctx.json) || !isDigit(ctx.json[p]) {
			return nil, errors.New("parse number error")
		}

		for ; p < len(ctx.json) && isDigit(ctx.json[p]); p++ {

		}
	}

	value, err := strconv.ParseFloat(ctx.json[:p], 64)
	if err != nil {
		return nil, errors.New("parse number error")
	}

	if p >= len(ctx.json) {
		return nil, errors.New("parse number error")
	}

	ctx.json = ctx.json[p:]

	node := &Node{}
	node.kind = Number
	node.number = &value
	return node, nil
}

// http://www.json.org/object.gif
func (ctx *Context) ParseObject() (*Node, error) {
	// read {
	ctx.RemoveWhite()
	err := ctx.RemoveACharacter('{')
	if err != nil {
		return nil, err
	}
	ctx.stack.Push(int32('}'))

	node := &Node{kind: Object}

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
		ctx.RemoveWhite()
		node, err := ctx.ParseNode()
		if err != nil {
			return nil, err
		}

		// save to map
		if node.object == nil {
			m := make(map[string]*Node)
			node.object = &m
		}
		(*node.object)[key] = node

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

	return node, nil
}

// http://www.json.org/array.gif
func (ctx *Context) ParseArray() (*Node, error) {
	// read [
	ctx.RemoveWhite()
	err := ctx.RemoveACharacter('[')
	if err != nil {
		return nil, err
	}
	ctx.stack.Push(int32(']'))

	node := &Node{kind: Array}

	for {
		// dispatch
		ctx.RemoveWhite()
		node, err := ctx.ParseNode()
		if err != nil {
			return nil, err
		}

		// save to array
		if node.array == nil {
			a := make([]*Node, 0)
			node.array = &a
		}

		*node.array = append(*node.array, node)

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

	return node, nil
}

// 在value位置处
func (ctx *Context) ParseNode() (*Node, error) {
	c, err := ctx.PeekACharacter()
	if err != nil {
		return nil, err
	}

	switch c {
	case '"':
		return ctx.ParseValue()
	case 'n':
		return ctx.ParseStr("null", Null)
	case 't':
		return ctx.ParseStr("true", True)
	case 'f':
		return ctx.ParseStr("false", False)
	case '[':
		return ctx.ParseArray()
	case '{':
		return ctx.ParseObject()
	default:
		return ctx.ParseNumber()
	}

	return nil, errors.New("unknown value")
}

func ParseJson(json string) (*JsonParse, error) {
	ctx := &Context{json: json}
	node, err := ctx.ParseObject()
	if err != nil {
		return nil, err
	}
	jp := &JsonParse{node: node}
	return jp, nil
}
