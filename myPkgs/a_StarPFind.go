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
