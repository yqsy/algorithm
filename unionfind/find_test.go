package unionfind

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestQuickFind(t *testing.T) {
	f, err := os.Open("tinyUF.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	// print
	NewQuickFindFromBufio(r)

}

func TestQuickFind2(t *testing.T) {
	tmpbuf := `10
3 4
4 9
`
	r := strings.NewReader(tmpbuf)

	// print
	NewQuickFindFromBufio(r)
}

func TestQuickFindMediumUF(t *testing.T) {
	f, err := os.Open("mediumUF.txt")

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	// print
	NewQuickFindFromBufio(r)
}

func TestUnionFind(t *testing.T) {
	f, err := os.Open("tinyUF.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	// print
	NewQuickUnionFromBufio(r)
}