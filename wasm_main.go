package main

import (
	"strconv"
	"sudoku-solver/Sudoku"
	"syscall/js"
	"time"
)

var sudoku = Sudoku.NewSudoku()

func LoadDataFromJavaScript(this js.Value, p []js.Value) interface{} {
	jsArray := p[0]
	var matrix [81]uint8

	for i := 0; i < 81; i++ {
		matrix[i] = uint8(jsArray.Index(i).Int())
	}
	//加载数据
	sudoku.LoadDataFromArray(matrix)
	startTime := time.Now().UnixNano()
	isSolved := Sudoku.DanceLinkSolver(sudoku)
	endTime := time.Now().UnixNano()
	jsResult := js.Global().Get("Array").New()
	if isSolved {
		for i := 0; i < 81; i++ {
			jsResult.Set(strconv.Itoa(i), js.ValueOf(sudoku.Cells[i].Val))
		}
		return js.ValueOf(map[string]interface{}{
			"isSolved": isSolved,
			"result":   jsResult,
			"timeMs":   (endTime - startTime) / 1e6,
		})
	} else {
		return js.ValueOf(map[string]interface{}{
			"isSolved": isSolved,
		})
	}
}

func main() {
	js.Global().Set("GO_sudoku", js.FuncOf(LoadDataFromJavaScript))
	<-make(chan struct{}) // 防止程序退出
}
