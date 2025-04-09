package mypkgs

import (
	"fmt"
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
func (igd *IntegerGridManager) AStarPrep(WallValues []int) {
	if igd.PFinder.IsFullyInitialized {
		if !igd.PFinder.Cursor.Position.IsEqualTo(igd.PFinder.StartPos) {
			igd.PFinder.Cursor.Position = igd.PFinder.StartPos
			igd.UpdateCursor()
		}
		var MarginValues [4]int = [4]int{1, 2, 2, 1}
		startNode := InitNode(igd.PFinder.StartPos, igd.PFinder.StartPos, igd.PFinder.EndPos)
		igd.PFinder.OpenList = make(CoordList, 0)
		igd.PFinder.ClosedList = make(CoordList, 0)
		igd.PFinder.BlockedList = make(CoordList, 0)
		// temp := igd.Imat.NodeAr_GetNeighbors4(startNode, MarginValues, igd.PFinder.EndPos)
		temp := igd.Imat.NodeAr_GetNeighbors4FILTERED(startNode, MarginValues, WallValues, igd.PFinder.EndPos)
		igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList, temp...)

		// for _, q := range igd.PFinder.n_OpenList {
		// 	q.MCost_Sum = q.MCost_toEnd + q.MCost_toParent
		// 	fmt.Printf("%s %d\n", q.ToString(), q.MCost_Sum)
		// }
		// fmt.Printf("--------------------------\n")
		NodesAr_Sort_ByF_Value(igd.PFinder.n_OpenList)
		// for _, r := range igd.PFinder.n_OpenList {
		// 	fmt.Printf("%s %d\n", r.ToString(), r.MCost_Sum)
		// }
		// fmt.Printf("--------------------------\n\n|||||||\n")
		// igd.PFinder.n_OpenList =

		// temp, _, _ := igd.Imat.GetNeighbors4(igd.PFinder.StartPos, [4]int{1, 2, 2, 1})
		// igd.PFinder.OpenList = append(igd.PFinder.OpenList, temp...)
		// igd.PFinder.OpenList.
		igd.PFinder.showNodes = true
		igd.BoardOverlayChange = true

		for len(igd.PFinder.n_OpenList) > 0 && !igd.PFinder.pathComplete {
			igd.AStarTICK(MarginValues, WallValues)

		}
	}

}

func (igd *IntegerGridManager) AStarTICK(MarginValues [4]int, WallValues []int) {
	if len(igd.PFinder.n_OpenList) > 0 {
		///fmt.Printf("SIZE OF N_OPENLIST: %d\n", len(igd.PFinder.n_OpenList))
		// for _, j := range igd.PFinder.n_OpenList {
		// 	fmt.Printf("%s\n", j.ToString())
		// }
		//igd.PFinder.n_OpenList
		// q, igd.PFinder.n_OpenList = NodesAr_PopFromFront(igd.PFinder.n_OpenList)

		// var tList []*Node
		// q, _ = NodesAr_PopFromFront(igd.PFinder.n_OpenList)
		// var q *Node

		q := igd.PFinder.n_OpenList[0]
		if len(igd.PFinder.n_OpenList) > 1 {
			igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList[:0], igd.PFinder.n_OpenList[0+1:]...)

		}
		// igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList, tList...)
		// igd.Pfinder.n_OpenList
		if q != nil {
			q.MCost_Sum = q.MCost_toEnd + q.MCost_toParent
			igd.Imat.DrawAGridTile(igd.BoardOverlayLayer, q.Postion, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{200, 200, 0, 255}, color.RGBA{255, 0, 0, 255}, 1.0, true, true)
			//fmt.Printf("For Q is %s  MCOST %d\n", q.ToString(), q.MCost_Sum)
			if q.Postion.IsEqualTo(igd.PFinder.EndPos) {
				//fmt.Printf("\n\n\n\nEND FOUND!\n\n\n\n\n\n")
				igd.BoardOverlayChange = true
				igd.PFinder.pathComplete = true

			} else {
				if igd.Imat.IsValid(q.Postion) {
					//fmt.Printf("IS VALID!!\n\n")
					// _, b2, b3 := igd.Imat.IsCoordValueInArrayOfValues_What_Exists(q.Postion, WallValues)//!IntArrayContains(WallValues, igd.Imat.GetCoordVal(q.Postion))
					if igd.Imat.GetCoordVal(q.Postion) == 1 {
						//fmt.Printf("NO WALLS!!\n\n")
						temp_successors := igd.Imat.NodeAr_GetNeighbors4FILTERED(q, MarginValues, WallValues, igd.PFinder.EndPos)
						//for loop(for each succesor){
						//fmt.Printf("SUCCESSORS :%d\n", len(temp_successors))
						for _, suc := range temp_successors {
							suc.MCost_Sum = suc.MCost_toEnd + suc.MCost_toParent
							//fmt.Printf("SUCCESOR: %d %d---\n %s\n", i, suc.MCost_Sum, suc.ToString())
							suc.MCost_Sum = suc.MCost_toEnd + suc.MCost_toParent
							// t, v := NodeAr_Contains_what(igd.PFinder.n_OpenList, suc)
							// if v && t != nil {

							// 	if t.MCost_toParent > suc.MCost_toParent {

							// 	}
							// }
							if !NodeAr_Contains(igd.PFinder.n_ClosedList, suc) && !q.Postion.IsEqualTo(suc.Postion) && !suc.Postion.IsEqualTo(igd.PFinder.StartPos) {
								igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList, suc)
								//fmt.Printf("ADDE DTO OPENLIST !\n")
							} else {
								//fmt.Printf("NOT ADDED TO OPENLIST !\n")

							}

							// if !NodeAr_Contains(igd.PFinder.n_ClosedList, suc) {
							// 	igd.PFinder.n_ClosedList = append(igd.PFinder.n_ClosedList, suc)
							// }
						} //---Temp Successors
						//remove any matching/duplicated positions from the openlist that have higher F-values;
						// igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList, temp_successors...)
						//end for loop

						override_CList := true
						for _, prev := range igd.PFinder.n_ClosedList {

							if prev.Postion.IsEqualTo(q.Postion) {
								//fmt.Printf("Q MATCH\nPREV:\t%s\n Q:\t%s\n", prev.ToString(), q.ToString())
								if q.MCost_Sum < prev.MCost_Sum {

									override_CList = true
								} else {
									override_CList = false
								}
							}
						}
						if override_CList {
							//fmt.Printf("Q ADDED!%d\n", q.ValueOnCoord)
							igd.PFinder.n_ClosedList = append(igd.PFinder.n_ClosedList, q)
						} else {
							// fmt.Printf("Q RETURNS TO THING%d\n", q.ValueOnCoord)
							// igd.PFinder.n_OpenList = append(igd.PFinder.n_OpenList, q)
						}

						//push q onto closed list;
						// igd.PFinder.n_ClosedList = append(igd.PFinder.n_ClosedList, q)
						//<---PROBLEM!!!! how do we send feedback to the main thing with a dead end?---->perhaps this is about finding the closed list and then using that to narrow down the potential path;
					} else {
						igd.Imat.DrawAGridTile(igd.BoardOverlayLayer, q.Postion, igd.BoardMargin.X, igd.BoardMargin.Y, igd.Tile_Size.X, igd.Tile_Size.Y, igd.Margin.X, igd.Margin.Y, color.RGBA{255, 0, 0, 255}, color.RGBA{255, 255, 0, 255}, 1.0, true, true)

						//fmt.Printf("Q (%d,%d) SLAMS INTO WALL  %d %d \n", q.Postion.X, q.Postion.Y, q.ValueOnCoord, igd.Imat.GetCoordVal(q.Postion))
					}
				}

			}
			// for _,openl := range igd.PFinder.n_OpenList{
			// 	for _,closel := range igd.PFinder.n_ClosedList{

			// 	}
			// }
			//fmt.Printf("----> END LEN:%3d Closed List:%3d\n\n", len(igd.PFinder.n_OpenList), len(igd.PFinder.n_ClosedList))
			NodesAr_Sort_ByF_Value(igd.PFinder.n_OpenList)
		} else {
			fmt.Printf("Q IS NILL!\n\n")
		}
	}
	// else {
	// 	// if !igd.PFinder.pathComplete {

	// 	// }
	// }
	// igd.PFinder.n_OpenList = NodesAr_Sort_ByF_Value(igd.PFinder.n_OpenList)
	// var q CoordInts
	// igd.PFinder.OpenList = igd.PFinder.OpenList.SortAndRemoveProblems(igd.PFinder.StartPos, igd.PFinder.Cursor.Position, igd.PFinder.EndPos, WallValues, &igd.Imat)
	// q, igd.PFinder.OpenList = igd.PFinder.OpenList.PopFromFront()

	//so I'm supposed to make 8 succesors to q and have said successors point to q as their parent; hence why I was thinking I needed a linked list;
	// however this isn't particularly useful;
	// so I'm going to probably need to create an array of like
	///fmt.Printf("-------------------\n")
	igd.PFinder.showNodes = true
	igd.BoardOverlayChange = true
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
