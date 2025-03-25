package mypkgs

type Cell struct {
	Position        CoordInts
	Neighbors       [8]CoordInts
	Neighbor_Values [8]int
	ticker          int
	ShowNeighbors   bool
	// MCost_Sum       int //equal to the sum of G and H
	// MCost_toStart   int //Movement Cost from StartNode to Positon
	// MCost_toEnd     int //Movement Cost from Position to End// will be -1 while not working
	// Parent        *Cell
	// Child            *Cell
	// Number          int
}

func (cell *Cell) InitP(StartPos CoordInts, EndPos CoordInts, imat IntMatrix) {
	// limit_X, limit_Y := imat.GetDimensions()

	cell.Position = StartPos
	temp, temp2, _ := imat.GetNeighbors8(StartPos)
	cell.Neighbor_Values = temp2

	cell.Neighbors = [8]CoordInts(temp)
	cell.ticker = 0
	cell.ShowNeighbors = false
	// cell.MCost_toStart = 0
	// cell.MCost_toEnd = -1
}

// func (cell *Cell) InitQ(StartPos, ThisPos, EndPos CoordInts, imat IntMatrix) {
// 	// limit_X, limit_Y := imat.GetDimensions()

// 	cell.Position = ThisPos
// 	temp, temp2, _ := imat.GetNeighbors8(ThisPos)
// 	cell.Neighbor_Values = temp2

//		cell.Neighbors = [8]CoordInts(temp)
//		cell.ticker = 0
//		a, b := ThisPos.GetDifferenceInInts(StartPos)
//		cell.ShowNeighbors = false
//		cell.MCost_toStart = a + b
//		cell.MCost_toEnd = -1
//	}
func (cell *Cell) UpdateCell(Imat IntMatrix) {
	temp, temp2, _ := Imat.GetNeighbors8(cell.Position)
	cell.Neighbor_Values = temp2
	cell.Neighbors = [8]CoordInts(temp)
}
func (cell *Cell) IsAt(cord CoordInts) bool {
	return cell.Position.IsEqualTo(cord)
}

// type Cells []Cell

// func (cells Cells) PushToReturn(c Cell) Cells {
// 	var temp Cells = make(Cells, len(cells))
// 	copy(cells, temp)
// 	temp = append(temp, c)
// 	return temp
// }
// func (cells Cells) RemoveCellFromList(cord CoordInts) (Cells, bool) {
// 	temp := make(Cells, 0)
// 	isThere := false
// 	for _, c := range cells {
// 		if !c.IsAt(cord) {
// 			temp = append(temp, c)
// 		}
// 	}
// 	return temp, isThere
// }

// func (cells Cells) SortDescDistanceFromMainPos(center CoordInts) Cells {
// 	temp := make(Cells, len(cells))
// 	copy(temp, cells)
// 	var tempPos CoordInts
// 	if(len(temp)>1){
// 		for range temp{
// 			for i:=1;i<(len(temp));i++{
// 				if(temp[i].Position.GetDistance(center)>temp[i-1].Position.GetDistance(center)){
// 					tempPos =
// 				}
// 			}
// 		}
// 	}
// }

// func (cells Cells) RemoveDuplicates() Cells {
// 	temp := make(Cells, len(cells))
// 	copy(cells, temp)

// 	return temp
// }

// func (cells Cells)PopFromFront()(Cell,Cells){
// 	temp:=cells[0]
// 	temp2:=
// }

type Node struct {
	Postion   CoordInts
	ParentPTR *Node
	ChildPTR  *Node
}
