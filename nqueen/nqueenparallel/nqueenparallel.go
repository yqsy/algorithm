package nqueenparallel

import (
	"github.com/yqsy/algorithm/nqueen/nqueen"
	"log"
)

type Task struct {
}

type Result struct {
	SolvedBoards [][]string
}

type ComputingUnit struct {
	TaskChan chan *Task

	ResultChan chan *Result

	// 问题规模
	N int

	// 所属子集(下标)
	Id int

	ctx *nqueen.Context
}

func (cu *ComputingUnit) FlushCtx() {
	cu.ctx = nqueen.NewContext(cu.N)
	cu.ctx.SetCell(0, cu.Id, cu.N, 'Q', true)
}

// 开启计算单元
func (cu *ComputingUnit) Run() {
	for {
		task := <-cu.TaskChan

		if task == nil {
			break
		}

		nqueen.Solve(1, cu.N, cu.ctx)
		result := &Result{}
		result.SolvedBoards = cu.ctx.SolvedBoards
		cu.ResultChan <- result

		cu.FlushCtx()
	}

	log.Printf("computing unit stop %v", cu.Id)
}

// 在初始化时就已经打好桩,分配好了计算的单元
func NewComputingUnit(n int, id int) *ComputingUnit {
	cu := &ComputingUnit{}
	cu.TaskChan = make(chan *Task)
	cu.ResultChan = make(chan *Result)
	cu.N = n
	cu.Id = id
	cu.FlushCtx()
	return cu
}

type ComputingModule struct {
	Unints map[int][]*ComputingUnit
}

func NewComputingModule() *ComputingModule {
	cm := &ComputingModule{}
	cm.Unints = make(map[int][]*ComputingUnit)
	return cm
}

func (cm *ComputingModule) Prepare(n int) {
	if _, ok := cm.Unints[n]; !ok {
		var units []*ComputingUnit

		for i := 0; i < n; i++ {
			unit := NewComputingUnit(n, i)
			go unit.Run()
			units = append(units, unit)
		}

		cm.Unints[n] = units
	}
}

func (cm *ComputingModule) Solve(n int) [][]string {
	for i := 0; i < n; i++ {
		cm.Unints[n][i].TaskChan <- &Task{}
	}

	var allSolvedBoards [][]string

	for i := 0; i < n; i++ {
		result := <-cm.Unints[n][i].ResultChan

		allSolvedBoards = append(allSolvedBoards, result.SolvedBoards...)
	}

	return allSolvedBoards
}

func (cm *ComputingModule) GetResult(n int) {
	for i := 0; i < n; i ++ {
		cm.Unints[n][i].TaskChan <- &Task{}
	}
}

