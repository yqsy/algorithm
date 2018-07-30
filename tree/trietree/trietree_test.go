package trie

import (
	"testing"
	"bufio"
	"strings"
	"reflect"
)

func TestSimple(t *testing.T) {

	trie := NewTrie()

	test := `Alaska
Arizona
Arkansas
California
Colorado
Connecticu
Hawaii
Idaho
Illinois
Indiana
Maine
Maryland
Massachusetts
Michigan
Minnesota
Mississippi
Missouri
Montana
`
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())

		trie.Insert(word)
	}

	out := trie.FindPrefix("ar")

	if !reflect.DeepEqual(out, []string{"arizona", "arkansas"}) {
		t.Fatal("err")
	}

}
