package json

import (
	"github.com/golang-collections/collections/stack"
	"errors"
	"fmt"
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

type Node struct {
	// 类型,Null,True,False
	kind Kind

	// 字符串
	value *string

	// 数字
	number *float64

	// 数组
	array *[]Node

	// 对象
	object *map[string]Node

	// 键值
	key *string
}

func (node *Node) GetString() string {
	if node.kind == String || node.value != nil {
		return *node.value
	} else {
		return ""
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

// slice移动到下一个指针处
func (ctx *Context) NextNoWhite() {
	i := 0
	for ; i < len(ctx.json); i++ {
		b := ctx.json[i]
		if b != ' ' && b != '\t' && b != '\r' && b != '\n' {
			break
		}
	}
	ctx.json = ctx.json[i:]
}

// 吃掉一个字符
func (ctx *Context) EatACharacter(c byte) error {
	if len(ctx.json) < 1 || ctx.json[0] != c {
		return errors.New(fmt.Sprintf("err %v", c))
	}

	ctx.json = ctx.json[1:]
	return nil
}

// 获取字符串,并把slice指向到第二个"的后面的idx
func (ctx *Context) GetStrAndEat() (string, error) {
	if len(ctx.json) < 2 {
		return "", errors.New("get key error")
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
		return "", errors.New("get key error")
	}

	key := ctx.json[first+1 : second]

	// 第二个"是最后一个字符就返回错误
	if second == len(ctx.json)-1 {
		return "", errors.New("too short")
	}

	ctx.json = ctx.json[second+1:]
	return key, nil
}

func ParseJson(json string) (*JsonParse, error) {
	ctx := &Context{json: json}

	// read {
	ctx.NextNoWhite()
	err := ctx.EatACharacter('{')
	if err != nil {
		return nil, err
	}

	// read key
	ctx.NextNoWhite()
	key, err := ctx.GetStrAndEat()
	if err != nil {
		return nil, err
	}

	// read colon
	ctx.NextNoWhite()
	err = ctx.EatACharacter(':')
	if err != nil {
		return nil, err
	}

	// read value
	value, err := ctx.GetStrAndEat()
	if err != nil {
		return nil, err
	}

	// read }
	ctx.NextNoWhite()
	err = ctx.EatACharacter('}')
	if err != nil {
		return nil, err
	}

	// generate node
	node := &Node{}
	node.kind = String
	node.key = &key
	node.value = &value

	jp := &JsonParse{}
	jp.node = node

	return jp, nil
}
