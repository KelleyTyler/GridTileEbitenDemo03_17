package mypkgs

import "fmt"

func (igd *IntegerGridManager) FindPather_BreadthFirst(start, end CoordInts, num int, direct int, pathway CoordList) (CoordList, bool) {
	var directions = make(CoordList, 4)
	dire := GetDiffer(start, end)
	directions = getDiffer2(dire)
	tempPath := make(CoordList, len(pathway))
	// rows := len(grid)
	// cols := len(grid[0])

	// Check if start and end points are valid
	if !igd.Imat.IsValid(start) || !igd.Imat.IsValid(end) || num > 512 {
		return tempPath, false // Invalid start or end point
	}
	if start.IsEqualTo(end) {
		if tempPath.CountInstances(start) < 1 {
			tempPath = append(tempPath, start)
		}
		return tempPath, true
	}
	for _, d := range directions {
		//newNum := num + 1
		newPoint := start.AddCoords(d)
		//fmt.Printf("HAH! %d\n\n", newNum)

		if igd.Imat.IsValid(newPoint) {

			return tempPath, false
		}
	}
	// return -1 // No path found
	tempPath = (tempPath)[:len(tempPath)-1]
	return tempPath, false
}

func (igd *IntegerGridManager) FindPather_BreadthFirst02(start, end CoordInts, num int, direct int, pathway CoordList) (CoordList, bool) {
	var directions = make(CoordList, 4)
	dire := GetDiffer(start, end)
	directions = getDiffer2(dire)
	tempPath := make(CoordList, len(pathway))
	// visited := make(IntMatrix, len(igd.Imat))
	// distance := make(IntMatrix, len(igd.Imat))
	// for i, a := range igd.Imat {
	// 	visited[i] = make([]int, len(a))
	// 	distance[i] = make([]int, len(a))
	// }
	copy(tempPath, pathway)
	if start.IsEqualTo(end) { //

		tempPath = append(tempPath, start)
		//fmt.Printf("Found it! %d\n", len(tempPath))
		return tempPath, true
	}
	if !igd.Imat.IsValid(start) || num > 1024 { //num > (2048/((num+2)/2))
		//fmt.Printf("IS Invalid!-----> straight away!! %d\n", num)
		return tempPath, false
	} //igd.PFinder.ClosedList.CountInstances(start) > 0
	//igd.Imat.GetCoordVal(start) == 3
	if igd.Imat.GetCoordVal(start) == 0 || igd.PFinder.BlockedList.CountInstances(start) > 0 || igd.Imat.GetCoordVal(start) == 4 {

		//fmt.Printf("IS ALSO Invalid, off the bat\n")
		return tempPath, false
	}
	// if igd.PFinder.FalsePos.CountInstances(start) > 1*(num+2) {
	// 	//fmt.Printf("IS ALSO Invalid, off the bat\n")
	// 	return tempPath, false
	// }
	// igd.Imat[start.Y][start.X] = 3
	//|| igd.PFinder.FalsePos.CountInstances(start) > 0
	tempPath = append(tempPath, start)
	var CLists []CoordList
	//var t int = 0
	igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, start)
	for _, d := range directions {
		newNum := num + 1
		newPoint := start.AddCoords(d)
		if !igd.Imat.IsValid(newPoint) || newPoint.IsEqualTo(start) {
			//igd.PFinder.ClosedList = append(igd.PFinder.ClosedList, newPoint)
		} else {
			tempPath2, isValid := igd.FindPather_BreadthFirst(newPoint, end, newNum, direct, tempPath)
			if isValid {

				//fmt.Printf("IS VALID! %d %d %d\n", i, newPoint.X, newPoint.Y)
				// igd.Imat[start.Y][start.X] = 1
				//igd.PFinder.FalsePos = append(igd.PFinder.FalsePos, newPoint)
				// return tempPath2, true

				if len(CLists) > 0 {
					if len(tempPath2) < len(CLists) {
						fmt.Printf("BIG C! %d\n", len(tempPath2))
						//
						for _, cc := range CLists[0] {
							igd.PFinder.FalsePos, _ = igd.PFinder.FalsePos.RemoveCoordFromList(cc)
						}
						igd.PFinder.FalsePos.RemoveDuplicates()
						CLists[0] = tempPath2
					}
					// if num > 16 {

					// }
					// else {
					// 	CLists = append(CLists, tempPath2)
					// }
				} else {
					CLists = append(CLists, tempPath2)
				}

			}
		}

		// else {
		// 	fmt.Printf("IS InValid! %d %d %d\n", i, newPoint.X, newPoint.Y)
		// 	return tempPath2, false
		// }
	}
	// igd.Imat[start.Y][start.X] = 1
	if len(CLists) > 0 {
		//temp := make(CoordList, 0)
		for i, _ := range CLists {
			tempPath = append(tempPath, CLists[i]...)
			tempPath.RemoveDuplicates()
		}

		return tempPath, true
		// if len(CLists) > 1 {
		// 	fmt.Printf("BIG C! %d\n", len(CLists))

		// 	t := CLists[0]
		// 	for _, cl := range CLists {
		// 		if len(cl) < len(t) {
		// 			t = cl
		// 		} else {
		// 			fmt.Printf("REMOVED! %d %d  %d\n", num, len(CLists), len(cl))
		// 			for _, cc := range cl {
		// 				igd.PFinder.FalsePos, _ = igd.PFinder.FalsePos.RemoveCoordFromList(cc)
		// 			}

		// 		}

		// 	}
		// 	fmt.Printf("SELECTED! %d\n", len(t))
		// 	return t, true
		// } else {
		// 	return CLists[0], true
		// }
	} else {
		// temp := true
		igd.PFinder.FalsePos, _ = igd.PFinder.FalsePos.RemoveCoordFromList(start)
		igd.PFinder.BlockedList = append(igd.PFinder.BlockedList, start)
		// if temp {
		// 	fmt.Printf("REMOVED!\n")
		// }
		tempPath = (tempPath)[:len(tempPath)-1]
		return tempPath, false
	}
}

func (igd *IntegerGridManager) FindPath(n int) {
	if igd.PFinder.HasFalsePos {
		igd.PFinder.FalsePos = make(CoordList, 0)
		igd.PFinder.HasFalsePos = false
		igd.PFinder.ClosedList = make(CoordList, 0)
		igd.PFinder.OpenList = make(CoordList, 0)
		igd.PFinder.BlockedList = make(CoordList, 0)
		//------
		// igd.PFinder.Distance_To_Start.IntmatrixDistanceMarking(igd.PFinder.EndPos)
		// igd.PFinder.Distance_To_Start.Intmatrix_Dist_With_Walls(igd.Imat, []int{0, 4})
		igd.PFinder.Distance.IntmatrixDistanceMarking(igd.PFinder.EndPos)
		igd.PFinder.Distance.Intmatrix_Dist_With_Walls(igd.Imat, []int{0, 4})
		// igd.PFinder.Distance.PrintMatrix()
		// igd.PFinder.OpenList =make
	}
	temp, isDone := make(CoordList, 0), false
	// temp = append(temp, igd.PFinder.StartPos)
	//fmt.Printf("STARTING UP FINDPATH\n")
	igd.PFinder.FalsePos, isDone = igd.FindPather_BreadthFirst(igd.PFinder.Cursor.Position, igd.PFinder.EndPos, 0, n, temp)
	if isDone {
		//fmt.Printf("IS Truely DONE\n")
		igd.PFinder.HasFalsePos = true
	} else {
		// fmt.Printf("FAILURE\n")
		igd.PFinder.HasFalsePos = true
	}
}
func (igd *IntegerGridManager) FindPath2(n int) {
	if !igd.PFinder.HasFalsePos {
		temp, isDone := make(CoordList, 0), false
		// temp = append(temp, igd.PFinder.StartPos)
		fmt.Printf("STARTING UP FINDPATH\n")
		igd.PFinder.FalsePos, isDone = igd.FindPather_BreadthFirst(igd.PFinder.StartPos, igd.PFinder.EndPos, 0, n, temp)
		if isDone {
			fmt.Printf("IS Truely DONE\n")
			igd.PFinder.HasFalsePos = true
		} else {
			//fmt.Printf("FAILURE\n")
			igd.PFinder.HasFalsePos = true
		}
	}
}

func GetDiffer(c1, c2 CoordInts) int {
	xx, yy := c1.GetDifferenceInInts(c2)
	var dire int = 0
	if xx > 0 {
		if yy > 0 {
			if xx < yy {
				dire = 6
			} else {
				dire = 1
			}

		} else if yy == 0 {
			dire = 12
		} else if yy < 0 {
			if (yy * -1) < xx {
				dire = 7
			} else {
				dire = 0
			}
		}
	} else if xx == 0 {
		if yy > 0 {
			dire = 6
		} else if yy == 0 {
			dire = 0 //??
			// fmt.Printf("Meow\n")
		} else if yy < 0 {
			// dire = 0
			dire = 8
		}
	} else if xx < 0 {
		if yy > 0 {
			if (xx * -1) < yy {
				dire = 2
			} else {
				dire = 5
			}
		} else if yy == 0 {
			dire = 14
		} else if yy < 0 { //north and west?
			if xx < yy {
				dire = 3
			} else {
				dire = 4
			}
		}
	}
	return dire
}
func getDiffer2(d int) CoordList {
	var directions = make(CoordList, 4)
	switch d {
	case 0:
		directions[0] = CoordInts{0, -1} //north
		directions[1] = CoordInts{1, 0}  //east
		directions[2] = CoordInts{0, 1}  //south
		directions[3] = CoordInts{-1, 0} //west
	case 1:
		directions[0] = CoordInts{1, 0}  //east
		directions[1] = CoordInts{0, 1}  //south
		directions[2] = CoordInts{-1, 0} //west
		directions[3] = CoordInts{0, -1} //north

	case 2:
		directions[0] = CoordInts{0, 1}  //south
		directions[1] = CoordInts{-1, 0} //west
		directions[2] = CoordInts{0, -1} //north
		directions[3] = CoordInts{1, 0}  //east

	case 3:
		directions[0] = CoordInts{-1, 0} //west
		directions[1] = CoordInts{0, -1} //north
		directions[2] = CoordInts{1, 0}  //east
		directions[3] = CoordInts{0, 1}  //south
	//counter clockwise
	case 4:
		directions[0] = CoordInts{0, -1} //north
		directions[1] = CoordInts{-1, 0} //west
		directions[2] = CoordInts{0, 1}  //south
		directions[3] = CoordInts{1, 0}  //east

	case 5:
		directions[0] = CoordInts{-1, 0} //west
		directions[1] = CoordInts{0, 1}  //south
		directions[2] = CoordInts{1, 0}  //east
		directions[3] = CoordInts{0, -1} //north
	case 6:
		directions[0] = CoordInts{0, 1}  //south
		directions[1] = CoordInts{1, 0}  //east
		directions[2] = CoordInts{0, -1} //north
		directions[3] = CoordInts{-1, 0} //west
	case 7:
		directions[0] = CoordInts{1, 0}  //east
		directions[1] = CoordInts{0, -1} //north
		directions[2] = CoordInts{-1, 0} //west
		directions[3] = CoordInts{0, 1}  //south

	//--- DIRECTIONAL
	case 8: //going North
		directions[0] = CoordInts{0, -1} //north
		directions[1] = CoordInts{1, 0}  //east
		directions[2] = CoordInts{-1, 0} //west
		directions[3] = CoordInts{0, 1}  //south
	case 9: //going North
		directions[0] = CoordInts{0, -1} //north
		directions[1] = CoordInts{-1, 0} //west
		directions[2] = CoordInts{1, 0}  //east
		directions[3] = CoordInts{0, 1}  //south
	case 10: //going South
		directions[0] = CoordInts{0, 1}  //south
		directions[1] = CoordInts{1, 0}  //east
		directions[2] = CoordInts{-1, 0} //west
		directions[3] = CoordInts{0, -1} //north
	case 11: //going South
		directions[0] = CoordInts{0, 1}  //south
		directions[1] = CoordInts{-1, 0} //west
		directions[2] = CoordInts{1, 0}  //east
		directions[3] = CoordInts{0, -1} //north
	case 12: //going East
		directions[0] = CoordInts{1, 0}  //east
		directions[1] = CoordInts{0, -1} //north
		directions[2] = CoordInts{0, 1}  //south
		directions[3] = CoordInts{-1, 0} //west
	case 13: //going East
		directions[0] = CoordInts{1, 0}  //east
		directions[1] = CoordInts{0, 1}  //south
		directions[2] = CoordInts{0, -1} //north
		directions[3] = CoordInts{-1, 0} //west
	case 14: //going West
		directions[0] = CoordInts{-1, 0} //west
		directions[1] = CoordInts{0, -1} //north
		directions[2] = CoordInts{0, 1}  //south
		directions[3] = CoordInts{1, 0}  //east
	case 15: //going West
		directions[0] = CoordInts{-1, 0} //west
		directions[1] = CoordInts{0, 1}  //south
		directions[2] = CoordInts{0, -1} //north
		directions[3] = CoordInts{1, 0}  //east
	default:
		directions[0] = CoordInts{0, -1}
		directions[1] = CoordInts{0, 1}
		directions[2] = CoordInts{1, 0}
		directions[3] = CoordInts{-1, 0}
	}
	return directions
}
