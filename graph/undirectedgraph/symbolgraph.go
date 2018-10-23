package undirectedgraph

import (
	"bufio"
	"os"
	"strings"
)

type SymbolGraph struct {
	st   map[string]int // 符号名 -> 索引
	keys []string       // 索引 -> 符号名
	g    *Graph
}

func NewSymbolGraph(stream, sp string) *SymbolGraph {
	s := &SymbolGraph{}
	s.st = make(map[string]int)

	f, err := os.Open(stream)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		lineBytes, _, err := r.ReadLine()
		if err != nil {
			break
		}

		line := string(lineBytes)
		a := strings.Split(line, sp)

		for i := 0; i < len(a); i++ {
			if _, ok := s.st[a[i]]; !ok {
				s.st[a[i]] = len(s.st)
			}
		}
	}

	s.keys = make([]string, len(s.st))

	for k := range s.st {
		s.keys[s.st[k]] = k
	}

	s.g = NewGraph(len(s.st))

	f2, err := os.Open(stream)
	if err != nil {
		panic(err)
	}

	r = bufio.NewReader(f2)

	for {
		lineBytes, _, err := r.ReadLine()
		if err != nil {
			break
		}

		line := string(lineBytes)
		a := strings.Split(line, sp)

		v := s.st[a[0]]

		for i := 1; i < len(a); i++ {
			s.g.AddEdge(v, s.st[a[i]])
		}
	}

	return s
}

func (s *SymbolGraph) Contains(symbol string) bool{
	if _, ok := s.st[symbol]; ok {
		return true
	} else {
		return false
	}
}

func (s *SymbolGraph) Index(symbol string) int {
	return s.st[symbol]
}

func (s *SymbolGraph) Name(v int) string {
	return s.keys[v]
}

func (s *SymbolGraph) G() *Graph {
	return s.g
}

