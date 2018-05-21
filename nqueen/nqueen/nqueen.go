package nqueen

type Context struct {
	Board [][]byte

	SolvedBoards [][]string

	// 竖向候选
	Cols []bool

	// 正斜着候选(/)  左上角idx=0
	SkewsForward []bool

	// 反斜着候选(\)  左下角idx=0
	SkewsBack []bool
}

func (ctx *Context) SetCell(row, col, n int, b byte, flag bool) {
	ctx.Board[row][col] = b
	ctx.Cols[col] = flag
	ctx.SkewsForward[row+col] = flag
	ctx.SkewsBack[(n-1-row)+col] = flag
}

func (ctx *Context) Check(row, col, n int) bool {
	if !ctx.Cols[col] &&
		!ctx.SkewsForward[row+col] &&
		!ctx.SkewsBack[(n-1-row)+col] {
		return true
	} else {
		return false
	}
}

func NewContext(n int) *Context {
	ctx := &Context{}
	ctx.Board = make([][]byte, n)
	for i := 0; i < n; i++ {
		ctx.Board[i] = make([]byte, n)

		for j := 0; j < n; j++ {
			ctx.Board[i][j] = '.'
		}
	}
	ctx.Cols = make([]bool, n)
	ctx.SkewsForward = make([]bool, 2*n-1)
	ctx.SkewsBack = make([]bool, 2*n-1)
	return ctx
}

func Solve(row, n int, ctx *Context) {
	if row == n {
		table := make([]string, n)
		for i, chessRow := range ctx.Board {
			table[i] = string(chessRow[:])
		}
		ctx.SolvedBoards = append(ctx.SolvedBoards, table)
		return
	}

	// 横向每一格进行尝试
	for i := 0; i < n; i++ {
		if !ctx.Check(row, i, n) {
			continue
		}
		ctx.SetCell(row, i, n, 'Q', true)
		Solve(row+1, n, ctx)
		ctx.SetCell(row, i, n, '.', false)
	}
}

func solveNQueens(n int) [][]string {
	ctx := NewContext(n)
	Solve(0, n, ctx)
	return ctx.SolvedBoards
}
