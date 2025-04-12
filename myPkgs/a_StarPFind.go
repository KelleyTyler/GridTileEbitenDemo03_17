package mypkgs

import (
	"image/color"
	"math"
)

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
func (igd *IntegerGridManager) AStarPrep(mode, FailMax int, WallValues []int) {
	if igd.PFinder.IsFullyInitialized {
		if !igd.PFinder.Cursor.Position.IsEqualTo(igd.PFinder.StartPos) {
			igd.PFinder.StartPos = igd.PFinder.Cursor.Position
			//igd.UpdateCursor(WallValues)
		}
		var MarginValues [4]int = [4]int{1, 2, 2, 1}
		startNode := InitNode(igd.PFinder.StartPos, igd.PFinder.StartPos, igd.PFinder.EndPos)
		// igd.PFinder.OpenList = make(CoordList, 0)
		// igd.PFinder.ClosedList = make(CoordList, 0)
		// igd.PFinder.BlockedList = make(CoordList, 0)
		if igd.PFinder.pathComplete {
			igd.PFinder.n_OpenList = make([]*Node, 0)
			igd.PFinder.n_ClosedList = make([]*Node, 0)
			igd.PFinder.n_BlockedList = make([]*Node, 0)
			igd.PFinder.pathComplete = false
		}
		if !startNode.Postion.IsEqualTo(igd.PFinder.EndPos) {
			temp := igd.Imat.NodeAr_GetNeighbors4FILTERED(startNode, MarginValues, WallValues, igd.PFinder.StartPos, igd.PFinder.EndPos)
			// temp := igd.Imat.NodeAr_GetNeighbors8_Filtered_MD(startNode, MarginValues, WallValues, igd.PFinder.EndPos)

			igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList, temp...)
			NodesAr_Sort_ByF_Value(igd.PFinder.n_OpenList)

			igd.PFinder.showNodes = true
			igd.BoardOverlayChange = true

			for len(igd.PFinder.n_OpenList) > 0 && !igd.PFinder.pathComplete { //
				if len(igd.PFinder.n_OpenList) < 512 && len(igd.PFinder.n_ClosedList) < 1024 {
					igd.AStarTICK(MarginValues, WallValues)
					NodesAr_RemoveDuplicates(igd.PFinder.n_OpenList)
					NodesAr_RemoveDuplicates(igd.PFinder.n_ClosedList)
				} else if !igd.PFinder.pathComplete {
					// fmt.Printf("\n\nFAILURE\n\n")
					igd.PFinder.pathComplete = true
				}

			}
		}

		//TODO: MAKE THIS LEANER AND MEANER AND OVERALL BETTER!
		//------Ideally it should resemble many of the other similar projects already on github;
		//---------HOWEVER INTEGRATED WITH INTMATRIX!!!
		//---- SEARCHES FOR MULTIPLE ROUTES
		//-------- SEARCHES FOR THE "FASTEST" ROUTE
		//---- SEARCHES TAKING ADVANTAGE OF DIFFERENT MOVEMENT SPEED CONDITIONS OVER DIFFERENT TILE VALUES (instead of having them all be "wall"
		//---- Searches that follow different patterns;
		//---- "Give Up" limits/conditions;
		//---- Multi-tile creature calculations? IE: I want a Giant who's 2x2 tiles as far as their horizontal footprint is concerned;
		//--------- REQUIRES MAZE GENERATOR THAT CAN HANDLE 2x2  WIDE HALLWAYS
		//--------VERTICAL PATHFINDING;
	}

}

func (igd *IntegerGridManager) AStarTICK(MarginValues [4]int, WallValues []int) {
	if len(igd.PFinder.n_OpenList) > 0 {

		q := igd.PFinder.n_OpenList[0]
		if len(igd.PFinder.n_OpenList) > 1 {
			igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList[:0], igd.PFinder.n_OpenList[0+1:]...)

		}

		if q != nil {
			q.MCost_Sum = q.MCost_toEnd + q.MCost_toStart
			igd.Imat.DrawAGridTile(igd.BoardOverlayLayer, q.Postion, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 200, 0, 255}, color.RGBA{255, 0, 0, 255}, 1.0, true, true)
			//fmt.Printf("For Q is %s  MCOST %d\n", q.ToString(), q.MCost_Sum)
			if q.Postion.IsEqualTo(igd.PFinder.EndPos) {
				// fmt.Printf("\nEND FOUND!\n")
				igd.BoardOverlayChange = true
				igd.PFinder.pathComplete = true

			} else {
				if igd.Imat.IsValid(q.Postion) {
					if igd.Imat.GetCoordVal(q.Postion) == 1 {
						// temp_successors := igd.Imat.NodeAr_GetNeighbors4FILTERED(q, MarginValues, WallValues, igd.PFinder.EndPos)
						// temp_successors := igd.Imat.NodeAr_GetNeighbors8_Filtered_MD(q, MarginValues, WallValues, igd.PFinder.EndPos)
						temp_successors := igd.Imat.NodeAr_GetNeighbors4_Filtered_MD(q, MarginValues, WallValues, igd.PFinder.StartPos, igd.PFinder.EndPos)

						for _, suc := range temp_successors {
							suc.MCost_Sum = suc.MCost_toEnd + suc.MCost_toStart
							suc.MCost_Sum = suc.MCost_toEnd + suc.MCost_toStart
							if !NodeAr_Contains(igd.PFinder.n_BlockedList, suc) && !NodeAr_Contains(igd.PFinder.n_ClosedList, suc) && !q.Postion.IsEqualTo(suc.Postion) && !suc.Postion.IsEqualTo(igd.PFinder.StartPos) {
								igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList, suc)
							}

						} //---Temp Successors
						override_CList := true
						for _, prev := range igd.PFinder.n_ClosedList {
							if prev.Postion.IsEqualTo(q.Postion) {
								if q.MCost_Sum < prev.MCost_Sum {
									override_CList = true
								} else {
									override_CList = false
								}
							}
						}
						if override_CList {
							igd.PFinder.n_ClosedList = append(igd.PFinder.n_ClosedList, q)
						}
					} else {
						igd.Imat.DrawAGridTile(igd.BoardOverlayLayer, q.Postion, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{255, 0, 0, 255}, color.RGBA{255, 255, 0, 255}, 1.0, true, true)
					}
				}

			}
			NodesAr_Sort_ByF_Value(igd.PFinder.n_OpenList)
		}
		igd.AStarAddToBlocked(MarginValues, WallValues)
	}

	igd.PFinder.showNodes = true
	igd.BoardOverlayChange = true
}

func (igd *IntegerGridManager) AStarAddToBlocked(MarginValues [4]int, WallValues []int) {
	if len(igd.PFinder.n_OpenList) > 0 {
		for _, a := range igd.PFinder.n_OpenList {
			if !a.Postion.IsEqualTo(igd.PFinder.EndPos) && !a.Postion.IsEqualTo(igd.PFinder.StartPos) {
				temp := igd.Imat.NodeAr_GetNeighbors4(a, MarginValues, igd.PFinder.StartPos, igd.PFinder.EndPos)
				// temp := igd.Imat.NodeAr_GetNeighbors8_Filtered_MD(a, MarginValues, WallValues, igd.PFinder.EndPos)

				num := 0
				for _, c := range temp {
					if !IntArrayContains(WallValues, igd.Imat.GetCoordVal(c.Postion)) && !NodeAr_Contains(igd.PFinder.n_BlockedList, c) {
						num++
					}
				}
				if num < 2 && !NodeAr_Contains(igd.PFinder.n_BlockedList, a) {

					igd.PFinder.n_BlockedList = append(igd.PFinder.n_BlockedList, a)
					NodesAr_RemoveByNode(igd.PFinder.n_OpenList, a)
					//igd.PFinder.n_OpenList = NodesAr_RemoveByNode(igd.PFinder.n_OpenList, a)

				}

			}

		}
		temp0 := make([]*Node, 0)
		for _, b := range igd.PFinder.n_OpenList {
			if !NodeAr_Contains(igd.PFinder.n_BlockedList, b) {
				temp0 = append(temp0, b)
			}
		}
		igd.PFinder.n_OpenList = temp0
		if len(igd.PFinder.n_ClosedList) > 0 {
			//looks for a a corner or dead end;
			//--->DEAD END: a dead end is when the 3 points are all walls or otherwise invalid, and neither it, nor it's neighbors are start or end;
			for _, a := range igd.PFinder.n_ClosedList {
				if !a.Postion.IsEqualTo(igd.PFinder.EndPos) && !a.Postion.IsEqualTo(igd.PFinder.StartPos) {
					temp := igd.Imat.NodeAr_GetNeighbors4(a, MarginValues, igd.PFinder.StartPos, igd.PFinder.EndPos)
					// temp := igd.Imat.NodeAr_GetNeighbors8_Filtered_MD(a, MarginValues, WallValues, igd.PFinder.EndPos)

					num := 0
					for _, c := range temp {
						if !IntArrayContains(WallValues, igd.Imat.GetCoordVal(c.Postion)) && !NodeAr_Contains(igd.PFinder.n_BlockedList, c) {
							d, dBol := NodeAr_Contains_what(igd.PFinder.n_ClosedList, c)
							if dBol {
								if d.ParentPTR != nil {
									if d.ParentPTR == a || d == a.ParentPTR {
										num++
									}
								}
							} else {
								num++
							}
						}

					}
					// b0 := IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[0].Postion)) && IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[1].Postion))
					// b1 := IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[1].Postion)) && IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[2].Postion))
					// b2 := IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[2].Postion)) && IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[3].Postion))
					// b3 := IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[3].Postion)) && IntArrayContains(WallValues, igd.Imat.GetCoordVal(temp[0].Postion))
					// c0 := b0 && b1
					// c1 := b1 && b2
					// c2 := b3 && b0
					// c3 := b2 && b0
					// c5 := b3 && b1
					if num < 2 && !NodeAr_Contains(igd.PFinder.n_BlockedList, a) {
						igd.PFinder.n_BlockedList = append(igd.PFinder.n_BlockedList, a)
						igd.PFinder.n_ClosedList = NodesAr_RemoveByNode(igd.PFinder.n_ClosedList, a)
						//igd.PFinder.n_ClosedList = NodesAr_RemoveByNode(igd.PFinder.n_ClosedList, a)

						// igd.PFinder.n_ClosedList = NodesAr_RemoveByNode(igd.PFinder.n_ClosedList, a)
						// if (c1 || c0 || c2 || c3 || c5) && !NodeAr_Contains(igd.PFinder.n_BlockedList, a) {

						// }
					}

				}

			}

			///-------->RMEOVE ORPHANED POINTS FROM CLOSEDLIST!
			temp := make([]*Node, 0)
			for _, b := range igd.PFinder.n_ClosedList {
				if !NodeAr_Contains(igd.PFinder.n_BlockedList, b) {
					temp = append(temp, b)
				}
			}
			igd.PFinder.n_ClosedList = temp
			// igd.PFinder.n_BlockedList

			///remove children from
		}
		NodesAr_RemoveDuplicates(igd.PFinder.n_BlockedList)

		///remove children from Blocklist

	}
	// if len(igd.PFinder.n_BlockedList) > 0 {
	// 	// igd.PFinder.n_BlockedList[len(igd.PFinder.n_BlockedList)-1]
	// 	for i, a := range igd.PFinder.n_BlockedList {

	// 	}
	// }

}

/*
H cost = how far away this node is from the End node
(manhattan distance or diagonal distance?)
*/
func (cord *CoordInts) GetH_Cost(target CoordInts) int {
	x1, y1 := cord.GetDifferenceInInts(target)
	return int(math.Abs(float64(x1)) + math.Abs(float64(y1)))
}

/*
G cost = how far away this node is from the starting node
(manhattan distance or diagonal distance?)
*/
func (cord *CoordInts) GetG_Cost(start CoordInts) int {
	x0, y0 := cord.GetDifferenceInInts(start)
	return int(math.Abs(float64(x0)) + math.Abs(float64(y0)))
}

/*
F cost = G_Cost + H_Cost
*/
func (cord *CoordInts) GetF_Cost(start, target CoordInts) int {
	h := cord.GetH_Cost(target)
	g := cord.GetG_Cost(start)
	return g + h
}

/*
 */
// func (imat *IntMatrix) Sort_On_F_Cost(start, cord, target CoordInts, cList CoordList) {
// 	go down through the list removing those invalids;

// }
func (cList CoordList) Sort_On_F_Cost(start, cord, target CoordInts) CoordList { //(), imat *IntMatrix
	//sorts descending
	temp := make(CoordList, len(cList))
	copy(temp, cList)

	for range temp {
		for i := 1; i < len(temp)-1; i++ {
			if temp[i].GetF_Cost(start, target) < temp[i-1].GetF_Cost(start, target) {
				tempCord := temp[i]
				temp[i] = temp[i-1]
				temp[i-1] = tempCord
			}
		}
	}
	return temp
}

func (cList CoordList) SortAndRemoveProblems(start, cord, target CoordInts, WallValues []int, imat *IntMatrix) CoordList {
	temp := cList.Sort_On_F_Cost(start, cord, target)
	// for i := 1; i < len(temp)-1; i++ {
	for i, c := range temp {
		if !imat.IsValid(c) {
			temp[i] = CoordInts{X: -1, Y: -1}
		} else {
			val := imat.GetCoordVal(c)
			if IntArrayContains(WallValues, val) {
				temp[i] = CoordInts{X: -1, Y: -1}
			}
		}
	}
	temp, _ = temp.RemoveCoordFromList(CoordInts{X: -1, Y: -1})

	return temp
}
