package mypkgs

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
