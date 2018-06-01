package json

import "testing"

func TestSimple(t *testing.T) {
	json := ` { "name" :   "123" }`

	jp, err := ParseJson(json)

	if err != nil {
		t.Fatal("err")
	}

	if jp == nil {
		t.Fatal("err")
	}

	value := jp.Get("name")

	if value == nil {
		t.Fatal("err")
	}

	if value.GetString() != "123" {
		t.Fatal("err")
	}
}
