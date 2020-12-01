package main

import (
	"fmt"
	"sudoku-solver/Sudoku"
	"time"
)

func exampleLoadDataFromArray() {
	sudoku := Sudoku.NewSudoku()
	//EASY LEVEL
	data := [...]uint8{
		0, 1, 0, 0, 9, 5, 0, 7, 6,
		0, 9, 0, 7, 0, 0, 0, 0, 0,
		7, 0, 0, 0, 0, 0, 4, 0, 0,

		0, 6, 7, 0, 0, 4, 9, 0, 5,
		5, 0, 0, 0, 6, 0, 0, 0, 0,
		0, 0, 0, 5, 3, 0, 6, 8, 0,

		0, 4, 9, 2, 0, 0, 5, 0, 0,
		0, 0, 2, 8, 0, 3, 1, 9, 0,
		0, 0, 5, 0, 1, 0, 0, 6, 0,
	}
	sudoku.LoadDataFromArray(data)
	sudoku.SetBlank(' ')
	sudoku.Show()

	startTime := time.Now().UnixNano()
	isSolved := Sudoku.DropSolver(sudoku)
	endTime := time.Now().UnixNano()
	if isSolved {
		sudoku.Show()
		fmt.Println("所用时间", float64((endTime-startTime)/1e6), "ms")
	}

}

func exampleLoadDataFromFile() {
	sudoku := Sudoku.NewSudoku()
	sudoku.LoadDataFromFile("test/1dream", ",")
	sudoku.Show()

	startTime := time.Now().UnixNano()
	isSolved := Sudoku.DropSolver(sudoku) // 解数独
	endTime := time.Now().UnixNano()
	if isSolved {
		sudoku.Show()
		fmt.Println("所用时间", float64((endTime-startTime)/1e6), "ms")
	}

}

func exampleLoadDataFromInput() {
	sudoku := Sudoku.NewSudoku()
	//加载数据
	sudoku.LoadDataFromInput()
	/* 需手动输入数据,例如：
	0 0 1 2 0 0 3 0 0
	2 0 0 4 0 0 0 0 0
	0 0 5 6 1 0 0 7 0
	6 8 0 0 0 0 0 0 0
	0 5 2 0 4 0 9 0 0
	0 0 0 0 0 0 0 5 3
	0 3 0 0 5 2 1 0 0
	0 0 0 0 0 4 0 0 8
	0 0 9 0 0 1 5 0 0
	*/
	sudoku.SetBlank(' ')
	sudoku.Show()
	//
	startTime := time.Now().UnixNano()
	isSolved := Sudoku.DropSolver(sudoku)
	endTime := time.Now().UnixNano()
	if isSolved {
		sudoku.Show()
		fmt.Println("所用时间", float64((endTime-startTime)/1e6), "ms")
	}

}
func main() {
	//exampleLoadDataFromFile()

	//exampleLoadDataFromArray()

	//exampleLoadDataFromInput()

	var startTime, endTime int64
	var isSolved bool // 是否有解
	sudoku := Sudoku.NewSudoku()
	file := "test/7hardest"

	sudoku.LoadDataFromFile(file, ",")
	sudoku.Show()
	fmt.Println("已读取数独", file)

	//使用摈弃优化的深度优先搜索法
	startTime = time.Now().UnixNano()
	isSolved = Sudoku.DropSolver(sudoku) // 解数独
	endTime = time.Now().UnixNano()
	if isSolved {
		sudoku.Show()
		fmt.Println("深度优先搜索法所用时间：", float64(endTime-startTime)/1e6, "ms")
	} else {
		fmt.Println("无解")
	}

	//使用舞蹈链法
	sudoku.LoadDataFromFile(file, ",")
	startTime = time.Now().UnixNano()
	isSolved = Sudoku.DanceLinkSolver(sudoku) // 解数独
	endTime = time.Now().UnixNano()

	if isSolved {
		sudoku.Show()
		fmt.Println("舞蹈链法所用时间：", float64(endTime-startTime)/1e6, "ms")
	} else {
		fmt.Println("无解")
	}

}
