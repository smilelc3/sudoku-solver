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
	Sudoku.DropSolver(sudoku)
	endTime := time.Now().UnixNano()

	sudoku.Show()
	fmt.Println("所用时间", float64((endTime-startTime)/1e6), "ms")
}

func exampleLoadDataFromFile() {
	sudoku := Sudoku.NewSudoku()
	sudoku.LoadDataFromFile("test/dream", ",")
	sudoku.Show()

	startTime := time.Now().UnixNano()
	Sudoku.DropSolver(sudoku) // 解数独
	endTime := time.Now().UnixNano()

	sudoku.Show()
	fmt.Println("所用时间", float64((endTime-startTime)/1e6), "ms")
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
	Sudoku.DropSolver(sudoku)
	endTime := time.Now().UnixNano()
	sudoku.Show()
	fmt.Println("所用时间", float64((endTime-startTime)/1e6), "ms")
}
func main() {
	//exampleLoadDataFromFile()

	//exampleLoadDataFromArray()

	//exampleLoadDataFromInput()
	//data := [][]int{
	////H  C1 C2 C3 C4 C5 C6 C7  //0
	//	{0, 0, 1, 0, 1, 1, 0}, //1
	//	{1, 0, 0, 1, 0, 0, 1}, //2
	//	{0, 1, 1, 0, 0, 1, 0}, //3
	//	{1, 0, 0, 1, 0, 0, 0}, //4
	//	{0, 1, 0, 0, 0, 0, 1}, //5
	//	{0, 0, 0, 1, 1, 0, 1}} //6

	//data := [][]int{
	////H  C1 C2 C3 C4 C5 C6  //0
	//	{1, 0, 1, 0, 0, 1}, //1
	//	{0, 1, 0, 1, 0, 0}, //2
	//	{0, 0, 0, 0, 1, 0}, //3
	//	{0, 1, 0, 1, 1, 1}, //4
	//}

	data := [][]int{
		//H  C1 C2 C3 C4 C5 C6 C7  //0
		{1, 0, 0, 1, 0, 0, 1}, //1
		{1, 0, 0, 1, 0, 0, 0}, //2
		{0, 0, 0, 1, 1, 0, 1}, //3
		{0, 0, 1, 0, 1, 1, 0}, //4
		{0, 1, 1, 0, 0, 1, 1}, //5
		{0, 1, 0, 0, 0, 0, 1}} //6

	fmt.Println(Sudoku.BaseDanceLinkXSolver(data))

	sudoku := Sudoku.NewSudoku()
	sudoku.LoadDataFromFile("test/easy", ",")
	sudoku.Show()

	startTime := time.Now().UnixNano()
	Sudoku.DanceLinkSolver(sudoku) // 解数独
	endTime := time.Now().UnixNano()

	sudoku.Show()
	fmt.Println("所用时间", float64((endTime-startTime)/1e3), "us")
}
