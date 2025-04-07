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

// assumes fully intialized, assumes it's solveable;
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

/*
H cost = how far away this node is from the End node
(manhattan distance or diagonal distance?)
*/
func (imat *IntMatrix) GetH_Cost(start, cord, target CoordInts) {

}

/*
G cost = how far away this node is from the starting node
(manhattan distance or diagonal distance?)
*/
func (imat *IntMatrix) GetG_Cost(start, cord, target CoordInts) {

}

/*
F cost = G_Cost + H_Cost
*/
func (imat *IntMatrix) GetF_Cost(start, cord, target CoordInts) {

}

/*
 */
func (imat *IntMatrix) Sort_On_F_Cost(start, cord, target CoordInts, cList CoordList) {

}
