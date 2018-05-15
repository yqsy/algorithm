package main

import (
	"github.com/yqsy/algorithm/sudoku/common"
	"github.com/yqsy/algorithm/sudoku/sudoku1"
	"github.com/yqsy/algorithm/sudoku/sudoku2"
	"github.com/yqsy/algorithm/sudoku/sudoku_extra"
	"time"
	"fmt"
	"github.com/yqsy/algorithm/sudoku/sudoku3"
)

func TestFoo(foo func(), hint string) {
	start := time.Now()

	foo()

	elapsed := time.Since(start).Nanoseconds()

	fmt.Printf("%v cost: %.2f us\n", hint, float64(elapsed/1000))
}

func main() {
	problem := "080001030500804706000270000920400003103958402400002089000029000305106008040300010\r\n"

	// problem := "000000010400000000020000000000050407008000300001090000300400200050100000000806000\r\n"

	TestFoo(func() {
		table1Extra, _ := sudoku_extra.ConvertLineToTable(problem)
		var pos int
		sudoku_extra.Solve(table1Extra, pos)
		fmt.Println(table1Extra.Tb)
	}, "sudoku_extra")

	TestFoo(func() {
		table1, _ := common.ConvertLineToTable(problem)
		var pos int
		sudoku1.Solve(table1, pos)
		fmt.Println(*table1)
	}, "sudoku1")

	TestFoo(func() {
		table2, _ := common.ConvertLineToTable(problem)
		var pos int
		candidate := table2.GetCandidate()
		sudoku2.Solve(table2, candidate, pos)
		fmt.Println(*table2)
	}, "sudoku2")

	TestFoo(func() {
		table3, _ := common.ConvertLineToTable(problem)
		var pos int
		candidate := table3.GetCandidate()
		sudoku3.Solve(table3, candidate, pos)
		fmt.Println(*table3)
	}, "sudoku3")

}
