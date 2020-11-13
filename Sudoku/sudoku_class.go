package Sudoku

// Sudoku Class
type Sudoku struct {
	Cells		[81] *Cell
	Boxes       [9] Box
	isLoadData  bool
	blankInShow byte
	Rows        [9] Row
	Cols        [9] Col

}

// 每一个9小格为一个Box
type Box struct {
	Cells  	[9] *Cell
	BoxIdx 	uint8
}

//每一个数字为一个Cell
type Cell struct {
	//绑定行列
	Row *Row
	Col *Col

	//绑定Box
	Box *Box

	// 受Cell修改影响的其他的cells集合
	AffectedCellsSet map[*Cell]struct{}

	Val uint8
	Idx uint8

}


// 定义横行
type Row struct {
	line
	RowIdx uint8
}

// 定义纵行
type Col struct {
	line
	ColIdx uint8
}

// 定义行
type line struct {
	Cells [9] *Cell
}
