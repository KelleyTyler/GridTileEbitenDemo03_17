package mypkgs

import "fmt"

/*
type Pathfinding struct {
	StartPos           CoordInts
	EndPos             CoordInts
	IsActive           bool
	IsStartInit        bool
	IsEndInit          bool
	IsFullyInitialized bool
	Nodes              CoordList
	Color              color.Color
	SpriteDim          CoordInts
	Cursor             Cell
	FalsePos           CoordList
	Moves              CoordList
	HasFalsePos        bool
}

*/
type Node struct {
	Postion       CoordInts
	ParentPTR     *Node
	ChildPTR      *Node
	MCost_Sum     int //equal to the sum of G and H
	MCost_toStart int //Movement Cost from StartNode to Positon
	MCost_toEnd   int //Movement Cost from Position to End// will be -1 while not working
}

//assumes fully intialized, assumes it's solveable;
func (igd *IntegerGridManager) AStarPrep() {
	if igd.PFinder.IsFullyInitialized {
		if igd.PFinder.Cursor.Position.IsEqualTo(igd.PFinder.StartPos) {

		} else {
			igd.PFinder.Cursor.Position = igd.PFinder.StartPos
			igd.UpdateCursor()
		}
	}
}

func (igd *IntegerGridManager) AStarLeap() {

}

func (igd *IntegerGridManager) FindPather(start, end CoordInts, num int, direct int, pathway CoordList) (CoordList, bool) {
	var directions = make(CoordList, 4)
	// directions[0] = CoordInts{-1, 0}
	// directions[1] = CoordInts{1, 0}
	// directions[2] = CoordInts{0, -1}
	// directions[3] = CoordInts{0, 1}
	switch direct {
	case 0:
		directions[0] = CoordInts{0, -1}
		directions[1] = CoordInts{0, 1}
		directions[2] = CoordInts{1, 0}
		directions[3] = CoordInts{-1, 0}
	case 1:
		directions[0] = CoordInts{0, 1}
		directions[1] = CoordInts{0, -1}
		directions[2] = CoordInts{1, 0}
		directions[3] = CoordInts{-1, 0}
	case 2:
		directions[0] = CoordInts{0, 1}
		directions[1] = CoordInts{1, 0}
		directions[2] = CoordInts{-1, 0}
		directions[3] = CoordInts{0, -1}
	case 3:
		directions[0] = CoordInts{1, 0}
		directions[1] = CoordInts{-1, 0}
		directions[2] = CoordInts{0, 1}
		directions[3] = CoordInts{0, -1}
	default:
		directions[0] = CoordInts{0, -1}
		directions[1] = CoordInts{0, 1}
		directions[2] = CoordInts{1, 0}
		directions[3] = CoordInts{-1, 0}
	}
	tempPath := make(CoordList, len(pathway))
	copy(tempPath, pathway)
	if !igd.Imat.IsValid(start) || num > 256 {
		fmt.Printf("IS Invalid!-----> straight away!! %d\n", num)
		return tempPath, false
	}
	if igd.Imat.GetCoordVal(start) == 0 || igd.Imat.GetCoordVal(start) == 3 || igd.Imat.GetCoordVal(start) == 4 {
		fmt.Printf("IS ALSO Invalid, off the bat\n")
		return tempPath, false
	}
	if start.IsEqualTo(end) {
		fmt.Printf("Found it!\n")
		tempPath = append(tempPath, start)
		return tempPath, true
	}
	igd.Imat[start.Y][start.X] = 3
	tempPath = append(tempPath, start)
	for i, d := range directions {
		newNum := num + 1
		newPoint := start.AddCoords(d)
		tempPath2, isValid := igd.FindPather(newPoint, end, newNum, direct, tempPath)
		if isValid {
			fmt.Printf("IS VALID! %d %d %d\n", i, newPoint.X, newPoint.Y)
			igd.Imat[start.Y][start.X] = 1
			return tempPath2, true
		} else {
			//fmt.Printf("IS InValid! %d %d %d\n", i, newPoint.X, newPoint.Y)
			//return tempPath2, true
		}
	}
	igd.Imat[start.Y][start.X] = 1
	tempPath = (tempPath)[:len(tempPath)-1]
	return tempPath, false
}

func (igd *IntegerGridManager) FindPath(n int) {
	if !igd.PFinder.HasFalsePos {
		temp, isDone := make(CoordList, 0), false
		// temp = append(temp, igd.PFinder.StartPos)
		fmt.Printf("STARTING UP FINDPATH\n")
		igd.PFinder.FalsePos, isDone = igd.FindPather(igd.PFinder.StartPos, igd.PFinder.EndPos, 0, n, temp)
		if isDone {
			fmt.Printf("IS Truely DONE\n")
			igd.PFinder.HasFalsePos = true
		}
	}
}
