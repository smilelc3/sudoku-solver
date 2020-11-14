package Sudoku

import "fmt"

//摒弃法 https://www.cnblogs.com/grenet/p/3138654.html

// 空白单元格 blankCell
// 唯一数单元格  onlyCell
// 无解单元格 noSolutionCell
// 无唯一数单元格 notOnlyCell

// 栈缓存搜索(dfs)可解答案
func DropSolver(sudoku *Sudoku) {
	// 维护一个查找步骤栈
	var stepStack []stepRecord

findBlankCell:
	// 找空白单元格
	cell := findBlankCellFromIdx(sudoku, 0)
	//	没有空白单元格 -> 输出解
	if cell == nil { // 已补全所有
		fmt.Println("已找到解")
		return
	}
	// 有空白单元格 -> 找唯一单元格
	var hasOnlyCell = false
	var availableNums []uint8
	var startIdx uint8 = 0
	for { //遍历所有空白单元格，找唯一单元格
		cell = findBlankCellFromIdx(sudoku, startIdx)
		if cell == nil {
			break // 找完所有cell
		}
		hasOnlyCell, availableNums = isOnlyCell(cell)
		if hasOnlyCell {
			// 找到唯一数
			break
		}
		startIdx = cell.Idx + 1
	}
	if hasOnlyCell { // 找到唯一单元格
		//填数
		cell.Val = availableNums[0]

		if len(stepStack) != 0 { // 栈非空，说明已经剔除唯一解，记录后续修改步骤
			stepStack[len(stepStack)-1].changedCells = append(stepStack[len(stepStack)-1].changedCells, cell)
		}
		goto checkNoSolutionCell

	} else { // 没有找到唯一单元格
		// 从空白格中找一个最少可能单元格
		var minPossibleNums []uint8
		cell, minPossibleNums = findMinPossibleNumsCell(sudoku)

		// 步骤栈记录所选的单元格和数字
		step := stepRecord{cell, minPossibleNums, nil}
		stepStack = append(stepStack, step)

		// 给单元格填数
		step.cell.Val = step.nums[0]
		//并检测是否会产生无解单元格
		goto checkNoSolutionCell
	}

	// 检测是否会产生无解单元格
checkNoSolutionCell:
	for nearCell, _ := range cell.AffectedCellsSet {
		if isNoSolutionCell(nearCell) { // 存在无解单元格
			goto isExistNoSolutionCell
		}
	}
	// 没有产生无解单元格，继续找空白单元格
	goto findBlankCell

	// 产生无解单元格，解空间删除该元素
isExistNoSolutionCell:
	// 步骤栈为空，无解
	if len(stepStack) == 0 {
		fmt.Println("无解")
		return
	}
	// 步骤栈非空
	// 恢复步骤栈最后状态， 并找该单元格下一个可能解

	// 取出stepStack的上一次的记录
	step := stepStack[len(stepStack)-1]
	stepStack = stepStack[0 : len(stepStack)-1]

	// 删除首位已尝试的解,并恢复数独状态
	step.nums = step.nums[1:]
	for _, cell := range step.changedCells {
		cell.Val = 0
	}
	step.cell.Val = 0

	// 无可能解，继续判断是否步骤栈非空
	if len(step.nums) == 0 {
		goto isExistNoSolutionCell
	} else { // 有可能解
		stepStack = append(stepStack, step)
		// 给单元格填数
		cell = step.cell
		step.cell.Val = step.nums[0]
		goto checkNoSolutionCell
	}

}

// 记录步骤，便于回溯
type stepRecord struct {
	cell         *Cell   // 记录当前操作的单元格
	nums         []uint8 // 记录候选答案
	changedCells []*Cell // 记录在此步骤中被改变的其他cell
}

// 找最小可能解的单元格
func findMinPossibleNumsCell(sudoku *Sudoku) (*Cell, []uint8) {
	var startIdx uint8 = 0
	var minPossibleCell *Cell
	var minPossibleNums []uint8
	numsLenMin := 10
	for { //遍历所有空白单元格,记录可能解
		cell := findBlankCellFromIdx(sudoku, startIdx)
		if cell == nil {
			break // 找完所有cell
		}
		_, tempNums := isOnlyCell(cell)
		if numsLenMin > len(tempNums) {
			numsLenMin = len(tempNums)
			minPossibleNums = tempNums
			minPossibleCell = cell
		}
		startIdx = cell.Idx + 1
	}
	return minPossibleCell, minPossibleNums
}

// 判断cell是否是唯一单元格
func isOnlyCell(cell *Cell) (bool, []uint8) {
	cellCountResult := getCellUsedCount(cell)
	notFillCount := 0
	var availableNums []uint8
	for idx, val := range cellCountResult[1:] {
		if val == 0 {
			notFillCount++
			availableNums = append(availableNums, uint8(idx+1))
		}
	}
	return notFillCount == 1, availableNums
}

// 从 startIdx位置开始寻找第一个空白单元格
func findBlankCellFromIdx(sudoku *Sudoku, startIdx uint8) *Cell {
	if startIdx >= 81 {
		return nil
	}
	for _, cell := range sudoku.Cells[startIdx:] {
		if cell.Val == 0 { // 找到空白单元格
			return cell
		}
	}
	return nil
}

// 获取给定单元格所关联的其他单元格，所占用数字的信息(不重复)
func getCellUsedCount(cell *Cell) [10]uint8 {
	result := [10]uint8{}
	//标记已用数
	nearCells := cell.GetCellAffectedCellsSet()
	for nearCell, _ := range nearCells {
		if nearCell.Val != 0 {
			result[nearCell.Val]++
		}
	}
	return result
}

// 检测某一cell是否是noSolutionCell
func isNoSolutionCell(cell *Cell) bool {
	isNoSolutionCell := false
	cellUsedCount := getCellUsedCount(cell)
	zeroCount := 0
	for _, val := range cellUsedCount[1:] {
		if val > 3 {
			// 存在三个以上的相同值，认定无解
			isNoSolutionCell = true
			break
		}

		if val == 0 {
			// 统计可用数字
			zeroCount++
		}
	}
	if zeroCount == 0 { // 无可用数字，认定无解
		isNoSolutionCell = true
	}
	return isNoSolutionCell
}
