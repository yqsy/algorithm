package main

import (
	"testing"
	"github.com/yqsy/algorithm/repeat_number/repeat_sort"
	"github.com/yqsy/algorithm/repeat_number/repeat_hash"
	"github.com/yqsy/algorithm/repeat_number/repeat_exchange"
	"github.com/yqsy/algorithm/repeat_number/repeat_exchange_space"
	"github.com/yqsy/algorithm/repeat_number/repeat_binary"
)

func createOneRepeatArrayBegin() []int {
	return []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func createOneRepeatArrayMid() []int {
	return []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10}
}

func createOneRepeatArrayEnd() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10}
}

func createMultiRepeat() []int {
	return []int{0, 1, 2, 3, 4, 5, 5, 6, 5, 8, 9, 10}
}

func createNoRepeat() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func createOneNum() []int {
	return []int{0}
}

func createNotInRangeNums() []int {
	return []int{0, 1, 2, 3, 4, 100, 6, 7, 8, 9, 10, 10}
}

func TestOneRepeat(t *testing.T) {

	if !repeat_sort.IsArrayRepeat(createOneRepeatArrayBegin()) {
		t.Fatal("err")
	}

	if !repeat_hash.IsArrayRepeat(createOneRepeatArrayBegin()) {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange.IsArrayRepeat(createOneRepeatArrayBegin()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange_space.IsArrayRepeat(createOneRepeatArrayBegin()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_binary.IsArrayRepeat(createOneRepeatArrayBegin()); !result {
		t.Fatal("err")
	}

	// ===

	if !repeat_sort.IsArrayRepeat(createOneRepeatArrayMid()) {
		t.Fatal("err")
	}

	if !repeat_hash.IsArrayRepeat(createOneRepeatArrayMid()) {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange.IsArrayRepeat(createOneRepeatArrayMid()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange_space.IsArrayRepeat(createOneRepeatArrayMid()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_binary.IsArrayRepeat(createOneRepeatArrayMid()); !result {
		t.Fatal("err")
	}

	// ===

	if !repeat_sort.IsArrayRepeat(createOneRepeatArrayEnd()) {
		t.Fatal("err")
	}

	if !repeat_hash.IsArrayRepeat(createOneRepeatArrayEnd()) {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange.IsArrayRepeat(createOneRepeatArrayEnd()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange_space.IsArrayRepeat(createOneRepeatArrayEnd()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_binary.IsArrayRepeat(createOneRepeatArrayEnd()); !result {
		t.Fatal("err")
	}
}

func TestMultiRepeat(t *testing.T) {
	if !repeat_sort.IsArrayRepeat(createMultiRepeat()) {
		t.Fatal("err")
	}

	if !repeat_hash.IsArrayRepeat(createMultiRepeat()) {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange.IsArrayRepeat(createMultiRepeat()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange_space.IsArrayRepeat(createMultiRepeat()); !result {
		t.Fatal("err")
	}

	if result, _ := repeat_binary.IsArrayRepeat(createMultiRepeat()); !result {
		t.Fatal("err")
	}
}

func TestNoRepeat(t *testing.T) {
	if repeat_sort.IsArrayRepeat(createNoRepeat()) {
		t.Fatal("err")
	}

	if repeat_hash.IsArrayRepeat(createNoRepeat()) {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange.IsArrayRepeat(createNoRepeat()); result {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange_space.IsArrayRepeat(createNoRepeat()); result {
		t.Fatal("err")
	}

	if result, _ := repeat_binary.IsArrayRepeat(createNoRepeat()); result {
		t.Fatal("err")
	}
}

func TestOneNum(t *testing.T) {
	if repeat_sort.IsArrayRepeat(createOneNum()) {
		t.Fatal("err")
	}

	if repeat_hash.IsArrayRepeat(createOneNum()) {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange.IsArrayRepeat(createOneNum()); result {
		t.Fatal("err")
	}

	if result, _ := repeat_exchange_space.IsArrayRepeat(createOneNum()); result {
		t.Fatal("err")
	}

	if result, _ := repeat_binary.IsArrayRepeat(createOneNum()); result {
		t.Fatal("err")
	}
}

func TestNotInRangeNums(t *testing.T) {

	if result, err := repeat_exchange.IsArrayRepeat(createNotInRangeNums()); result || err == nil {
		t.Fatal("err")
	}

	if result, err := repeat_exchange_space.IsArrayRepeat(createNotInRangeNums()); result || err == nil {
		t.Fatal("err")
	}

	if result, err := repeat_binary.IsArrayRepeat(createNotInRangeNums()); result || err == nil {
		t.Fatal("err")
	}
}
