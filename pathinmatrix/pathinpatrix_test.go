package pathinmatrix

import "testing"

func TestSimple(t *testing.T) {

	a := [][]byte{
		{'a', 'b', 't', 'g'},
		{'c', 'f', 'c', 's'},
		{'j', 'd', 'e', 'h'}}

	if ok, err := hasPath(a, "bfce"); !ok || err != nil {
		t.Fatal("err")
	}

}

func TestSimple2(t *testing.T) {

	a := [][]byte{
		{'a'},
		{'c'},
		{'j'}}

	if ok, err := hasPath(a, "cj"); !ok || err != nil {
		t.Fatal("err")
	}
}

func TestSimple3(t *testing.T) {

	a := [][]byte{
		{'a', 'a', 'a'},
		{'a', 'a', 'a'},
		{'a', 'a', 'a'}}

	if ok, err := hasPath(a, "aaa"); !ok || err != nil {
		t.Fatal("err")
	}
}

func TestSimple4(t *testing.T) {

	a := [][]byte{
		{'a', 'a', 'a'},
		{'a', 'a', 'a'},
		{'a', 'a', 'a'}}

	if ok, err := hasPath(a, "bbb"); ok || err != nil {
		t.Fatal("err")
	}
}
