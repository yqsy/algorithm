package main

import (
	"os"
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/yqsy/algorithm/repeat_number/repeat_sort"
	"github.com/yqsy/algorithm/repeat_number/repeat_hash"
	"github.com/yqsy/algorithm/repeat_number/repeat_exchange"
	"github.com/yqsy/algorithm/repeat_number/repeat_exchange_space"
	"github.com/yqsy/algorithm/repeat_number/repeat_binary"
)

var usage = `Usage:
%v number [--rand]
`

func TestAndPrintTime(method string, foo func()) {
	start := time.Now()
	foo()
	elapsed := time.Since(start)
	fmt.Printf("method:%v take:%.2f us\n", method, float64(elapsed.Nanoseconds())/1000)
}

func main() {
	arg := os.Args

	usage = fmt.Sprintf(usage, arg[0])

	if len(arg) < 2 {
		fmt.Printf(usage)
		return
	}

	number, err := strconv.Atoi(arg[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v numbers\n", number)

	isRand := false
	if len(arg) > 2 {
		if arg[2] == "--rand" {
			isRand = true
			fmt.Printf("rand\n")
		}
	}

	var buf []int
	if isRand {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		buf = make([]int, number)
		for i := 0; i < len(buf); i++ {
			buf[i] = rnd.Intn(len(buf))
		}
	} else {
		buf = make([]int, number)
		for i := 0; i < len(buf); i++ {
			buf[i] = i
		}

		// 最后一个数字是len(buf)-1
		// 修改成len(buf) - 2,产生重复
		buf[len(buf)-1] = len(buf) - 2
	}

	bufSort := make([]int, number)
	copy(bufSort, buf)

	bufExchange := make([]int, number)
	copy(bufExchange, buf)

	TestAndPrintTime("repeat_sort", func() {
		repeat_sort.IsArrayRepeat(bufSort)
	})

	TestAndPrintTime("repeat_hash", func() {
		repeat_hash.IsArrayRepeat(buf)
	})

	TestAndPrintTime("repeat_exchange", func() {
		if result, err := repeat_exchange.IsArrayRepeat(bufExchange); err != nil {
			panic(err)
			_ = result
		}

	})

	TestAndPrintTime("repeat_exchange_space", func() {
		if result, err := repeat_exchange_space.IsArrayRepeat(buf); err != nil {
			panic(err)
			_ = result
		}
	})

	TestAndPrintTime("repeat_binary", func() {
		if result, err := repeat_binary.IsArrayRepeat(buf); err != nil {
			panic(err)
			_ = result
		}
	})

	fmt.Printf("done\n")
}
