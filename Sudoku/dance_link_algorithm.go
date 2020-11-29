package Sudoku

import (
	"fmt"
)

// 精确覆盖问题的定义：给定一个由0-1组成的矩阵，是否能找到一个行的集合，使得集合中每一列都恰好包含一个1

type DanceLink struct {
	Headers []*node
	Nodes   []*node // 存所有Node

	RowFirstNode []*node //记录每行第一个节点，辅助空间，便于从行查找

	AnsStack []int //答案栈
}

var EXIST = struct{}{}

//元素类
type node struct {
	Status int
	Left   *node
	Right  *node
	Up     *node
	Down   *node
	RowNum int
	ColNum int
}

// 标记列首元素C，返回被标记的元素集合
func (danceLink *DanceLink) markOneHeaderNode(header *node) (map[*node]struct{}, []int, []map[*node]struct{}) {
	markedNodesSet := make(map[*node]struct{}) // 当前列首元素所标记的结果

	var rowNumSet []int                           // 涉及到的行号集合
	var rowNumLinkHeadersSet []map[*node]struct{} // 行号对应的所关联的其他列首元素

	markedNodesSet[header] = EXIST
	cNode := header.Down
	for cNode != header {
		markedNodesSet[cNode] = EXIST
		rowNumSet = append(rowNumSet, cNode.RowNum)
		rNode := cNode.Right
		rowNumLinkHeadersSet = append(rowNumLinkHeadersSet, make(map[*node]struct{}))
		for rNode != cNode {
			markedNodesSet[rNode] = EXIST
			rowNumLinkHeadersSet[len(rowNumSet)-1][danceLink.Headers[rNode.ColNum]] = EXIST
			rNode = rNode.Right
		}
		cNode = cNode.Down
	}
	return markedNodesSet, rowNumSet, rowNumLinkHeadersSet
}

// 删除某一个元素
func (danceLink *DanceLink) removeOneNode(node *node) {
	// 不修改node本身的标记，仅修改相应
	node.Up.Down = node.Down
	node.Down.Up = node.Up
	node.Left.Right = node.Right
	node.Right.Left = node.Left
}

//恢复某一个元素
func (danceLink *DanceLink) resumeOneNode(node *node) {
	//依照本身标记恢复
	node.Up.Down = node
	node.Down.Up = node
	node.Left.Right = node
	node.Right.Left = node
}

// 总入口函数
func DanceLinkSolver(data [][]int) {
	danceLink := new(DanceLink)
	//根据列数创建
	dataRowLength := len(data)
	dataColLength := len(data[0])

	// 构建RowFirstNode
	for rowIdx := 0; rowIdx <= dataRowLength+1; rowIdx++ {
		danceLink.RowFirstNode = append(danceLink.RowFirstNode, nil)
	}

	head := new(node) //创建辅助节点
	danceLink.Headers = append(danceLink.Headers, head)
	danceLink.RowFirstNode[0] = head
	for colNum := 1; colNum <= dataColLength; colNum++ {
		hNode := new(node) //创建头部节点

		hNode.Up = hNode
		hNode.Down = hNode

		danceLink.Headers[colNum-1].Right = hNode
		hNode.Left = danceLink.Headers[colNum-1]

		hNode.ColNum = colNum
		danceLink.Headers = append(danceLink.Headers, hNode)
	}
	//首尾相连
	danceLink.Headers[dataColLength].Right = head
	head.Left = danceLink.Headers[dataColLength]

	//	构建交叉十字循环双向链
	for dataRowNum, _rows := range data {
		for dataColNum, val := range _rows {
			//fmt.Printf("data[%d][%d] = %d \n", dataRowNum, dataColNum, val)
			if val != 0 {
				// 创建body节点
				bNode := new(node)
				bNode.RowNum = dataRowNum + 1
				bNode.ColNum = dataColNum + 1
				bNode.Left = bNode
				bNode.Right = bNode

				// 创建四个方向的link
				preNode := danceLink.Headers[dataColNum+1] //  创建临时指针node
				// 纵向插入
				for {
					nextRowNum := preNode.Down.RowNum
					if nextRowNum == 0 {
						nextRowNum = dataRowLength + 1
					}
					if preNode.RowNum < dataRowNum+1 && dataRowNum+1 < nextRowNum { //找到插入位置
						// 在pNode后面插入bNode
						bNode.Up = preNode
						bNode.Down = preNode.Down
						bNode.Down.Up = bNode
						preNode.Down = bNode
						break
					}
					preNode = preNode.Down
				}
				//横向插入
				preNode = danceLink.RowFirstNode[bNode.RowNum]
				if preNode == nil {
					danceLink.RowFirstNode[bNode.RowNum] = bNode // 因为从左至右遍历，先遇到的为firstNode
				} else {
					for {
						nextColNum := preNode.Right.ColNum
						if nextColNum <= preNode.ColNum {
							nextColNum = dataColLength + 1
						}
						if preNode.ColNum < bNode.ColNum && nextColNum > bNode.ColNum {
							bNode.Left = preNode
							bNode.Right = preNode.Right
							bNode.Right.Left = bNode
							preNode.Right = bNode
							break
						}
						preNode = preNode.Right
					}
				}

				danceLink.Nodes = append(danceLink.Nodes, bNode)
			}
		}
	}

	// 递归求解
	if dancing(danceLink) == true {
		fmt.Println("Ans: ", danceLink.AnsStack)
	} else {
		fmt.Println("无解")
	}

}

// 递归函数
func dancing(danceLink *DanceLink) bool {
	/*
		1、dancing函数的入口
		2、判断Head.Right=Head？，若是，输出答案，返回True，退出函数。
		3、获得Head.Right的列首元素C
		4、标示列首元素C（标示元素C，指的是标示C、和C所在列的所有元素、以及该元素所在行的元素，并从双向链中移除这些元素）
		5、获得元素C所在列的一个body元素
		6、标示该body元素同行的其余body元素所在的列首元素
		7、获得一个简化的问题，递归调用dancing函数，若返回的True，则返回True，退出函数。
		8、若返回的是False，则回标该元素同行的其余元素所在的列首元素，回标的顺序和之前标示的顺序相反
		9、获得元素C所在列的下一个元素，若有，跳转到步骤6
		10、若没有，回标元素C，返回False，退出函数。
	*/
	if danceLink.Headers[0].Right == danceLink.Headers[0] {
		return true
	}
	hNode := danceLink.Headers[0].Right
	markedNodesSet, rowNumSet, rowNumLinkHeadersSet := danceLink.markOneHeaderNode(hNode)

	if rowNumSet == nil { // 当前为空集
		return false
	}
	for idx, rowNum := range rowNumSet {
		// 深copy markedNodesSet到allMarkedNodesSet  大坑
		allMarkedNodesSet := make(map[*node]struct{})
		for node := range markedNodesSet {
			allMarkedNodesSet[node] = EXIST
		}
		// 标示该同行的其余body元素所在的列首元素
		for header := range rowNumLinkHeadersSet[idx] {
			linkMarkedNodesSet, _, _ := danceLink.markOneHeaderNode(header)
			for node := range linkMarkedNodesSet {
				allMarkedNodesSet[node] = EXIST
			}
		}
		// 删除这一行关联的其他列首元素
		for markNode := range allMarkedNodesSet {
			danceLink.removeOneNode(markNode)
		}

		//行号加入答案栈
		danceLink.AnsStack = append(danceLink.AnsStack, rowNum)

		// 递归
		ok := dancing(danceLink)
		if ok {
			return true
		}

		// 回滚
		for markNode := range allMarkedNodesSet {
			danceLink.resumeOneNode(markNode)
		}
		danceLink.AnsStack = danceLink.AnsStack[0 : len(danceLink.AnsStack)-1]

	}
	return false

}
