package Sudoku

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// 构造函数
func NewSudoku() *Sudoku {
	sudoku := new(Sudoku)

	sudoku.isLoadData = false
	sudoku.blankInShow = ' '

	var row, col uint8
	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			targetCell := new(Cell)
			// 为sudoku构建Rows Cols
			sudoku.Rows[row].Cells[col] = targetCell
			sudoku.Cols[col].Cells[row] = targetCell
			sudoku.Rows[row].RowIdx = row
			sudoku.Cols[col].ColIdx = col

			//为sudoku构建Boxes
			boxIdx := sudoku.getBoxIdxByCoords(row, col)
			cellIdx := sudoku.GetCellIdByCoords(row, col)
			sudoku.Boxes[boxIdx].Cells[cellIdx] = targetCell
			sudoku.Boxes[boxIdx].BoxIdx = boxIdx

			// 构建Cell(值，索引，父链)
			dataIdx := row*9 + col
			targetCell.Idx = dataIdx
			targetCell.Row = &sudoku.Rows[row]
			targetCell.Col = &sudoku.Cols[col]
			targetCell.Box = &sudoku.Boxes[boxIdx]

			sudoku.Cells[dataIdx] = targetCell

		}
	}

	// 构造cell对象的
	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			thisCell := sudoku.Rows[row].Cells[col]
			thisCell.AffectedCellsSet = thisCell.GetCellAffectedCellsSet()
		}
	}
	return sudoku
}

//从List加载数据
func (Sudoku *Sudoku) LoadDataFromArray(data [81]uint8) {
	for idx, cell := range Sudoku.Cells {
		cell.Val = data[idx]
	}
	Sudoku.isLoadData = true
}

//从控制框中输入数据
func (Sudoku *Sudoku) LoadDataFromInput() {
	fmt.Println("请输入81个 [1-9] 的数字， 待填写数字用 0 代替, 数字间用空格间隔")
	var data [9 * 9]uint8
	for idx := 0; idx < (9 * 9); idx++ {
		_, err := fmt.Scanf("%d", &data[idx])
		if err != nil {
			panic(err)
		}
		if data[idx] > 9 {
			panic("数字不合法")
		}
	}
	Sudoku.LoadDataFromArray(data)
}

//从文本文件中输入数据
func (Sudoku *Sudoku) LoadDataFromFile(filePath string, splitMark string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("文件读取错误")
		panic(err)

	}
	idx := 0
	stringLines := strings.Split(string(data), "\n")
	for _, Line := range stringLines {
		stringValues := strings.Split(Line, splitMark)
		for _, stringValue := range stringValues {
			stringValue = strings.Trim(stringValue, " \t\r")
			if stringValue == "" {
				continue
			}
			val, err := strconv.Atoi(stringValue)
			if err != nil {
				fmt.Println("文件内容错误")
				panic(err)
			}
			if val > 9 || val < 0 {
				panic("文件数据错误")
			}
			Sudoku.Cells[idx].Val = uint8(val)
			idx++
		}
	}
	if idx != 81 {
		panic("数据内容错误")
	} else {
		Sudoku.isLoadData = true
	}

}

// 通过坐标获取值
func (Sudoku *Sudoku) GetValByCoords(row uint8, col uint8) uint8 {
	return Sudoku.GetCellByCoords(row, col).Val
}

//通过坐标获取Box
func (Sudoku *Sudoku) GetBoxByCoords(row uint8, col uint8) *Box {
	return &Sudoku.Boxes[Sudoku.getBoxIdxByCoords(row, col)]
}

// 通过坐标返回Box idx
func (Sudoku *Sudoku) getBoxIdxByCoords(row uint8, col uint8) uint8 {
	boxIdx := row/3*(9/3) + col/3
	return boxIdx
}

//通过坐标返回Cell位于Box的内部 cell id
func (Sudoku *Sudoku) GetCellIdByCoords(row uint8, col uint8) uint8 {
	cellIdx := row%3*3 + col%3
	return cellIdx
}

//通过坐标获取Cell对象
func (Sudoku *Sudoku) GetCellByCoords(row uint8, col uint8) *Cell {
	return Sudoku.GetBoxByCoords(row, col).Cells[Sudoku.GetCellIdByCoords(row, col)]
}

// 格式化数独输出
func (Sudoku *Sudoku) GetFormString() string {
	if Sudoku.isLoadData == false {
		panic("请先加载数据")
	}
	var row, col uint8
	formString := ""
	formString += "+-----------------------------+" + "\n"
	for row = 0; row < 9; row++ {
		formString += "| "
		for col = 0; col < 9; col++ {
			if Sudoku.GetValByCoords(row, col) == 0 {
				formString += string(Sudoku.blankInShow)
			} else {
				formString += fmt.Sprintf("%d", Sudoku.GetValByCoords(row, col))
			}

			if col == 2 || col == 5 {
				formString += " | "
			} else if col == 8 {
				formString += " |"
			} else {
				formString += "  "
			}

		}
		formString += "\n"
		if row == 2 || row == 5 {
			formString += "|---------+---------+---------|" + "\n"
		}
	}
	formString += "+-----------------------------+" + "\n"
	return formString
}

// 以一维数组返回所有数据
func (Sudoku *Sudoku) GetDataByArray() [81]uint8 {
	data := [81]uint8{}
	for idx, cell := range Sudoku.Cells {
		data[idx] = cell.Val
	}
	return data
}

// 设置blank
func (Sudoku *Sudoku) SetBlank(blank byte) {
	Sudoku.blankInShow = blank
}

// 打印数独
func (Sudoku *Sudoku) Show() {
	fmt.Print(Sudoku.GetFormString())
}

// 获取某一个cell所受影响的周边cells的集合（借助map[*Cell]struct{}实现）
func (cell *Cell) GetCellAffectedCellsSet() map[*Cell]struct{} {
	if cell.AffectedCellsSet != nil {
		return cell.AffectedCellsSet
	}
	var EXIST = struct{}{}
	cellsSet := make(map[*Cell]struct{})
	for _, nearCell := range cell.Box.Cells {
		if nearCell != cell {
			cellsSet[nearCell] = EXIST
		}
	}
	for _, nearCell := range cell.Row.Cells {
		if nearCell != cell {
			cellsSet[nearCell] = EXIST
		}
	}

	for _, nearCell := range cell.Col.Cells {
		if nearCell != cell {
			cellsSet[nearCell] = EXIST
		}
	}
	return cellsSet
}
