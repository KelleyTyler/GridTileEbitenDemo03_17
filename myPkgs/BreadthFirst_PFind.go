package mypkgs

import "fmt"

func (igd *IntegerGridManager) FindPather_BreadthFirst(start, end CoordInts, num int, direct int, pathway CoordList) (CoordList, bool) {
	var directions = make(CoordList, 4)
	// directions[0] = CoordInts{-1, 0}
	// directions[1] = CoordInts{1, 0}
	// directions[2] = CoordInts{0, -1}
	// directions[3] = CoordInts{0, 1}
	xx, yy := start.GetDifferenceInInts(end)
	dire := 0
	// if xx >= 0 && yy >= 0 {
	// 	dire = 6
	// } else if xx <= 0 && yy >= 0 {
	// 	dire = 2
	// } else if xx <= 0 && yy <= 0 {
	// 	dire = 4
	// } else if xx >= 0 && yy <= 0 {
	// 	dire = 0
	// }
	if xx > 0 {
		if yy > 0 {
			if xx < yy {
				dire = 6
			} else {
				dire = 1
			}

		} else if yy == 0 {
			dire = 7
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
		} else if yy < 0 {
			dire = 0
		}
	} else if xx < 0 {
		if yy > 0 {
			if (xx * -1) < yy {
				dire = 2
			} else {
				dire = 5
			}
		} else if yy == 0 {
			dire = 5
		} else if yy < 0 { //north and west?
			if xx < yy {
				dire = 3
			} else {
				dire = 4
			}
		}
	}
	// if dire == direct {
	// 	if dire != 0 {
	// 		dire = 0
	// 	} else {
	// 		dire = 1
	// 	}
	// }
	switch dire {
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
	default:
		directions[0] = CoordInts{0, -1}
		directions[1] = CoordInts{0, 1}
		directions[2] = CoordInts{1, 0}
		directions[3] = CoordInts{-1, 0}
	}
	tempPath := make(CoordList, len(pathway))
	copy(tempPath, pathway)
	if !igd.Imat.IsValid(start) || num > 256 {
		//fmt.Printf("IS Invalid!-----> straight away!! %d\n", num)
		return tempPath, false
	}
	if igd.Imat.GetCoordVal(start) == 0 || igd.Imat.GetCoordVal(start) == 3 || igd.Imat.GetCoordVal(start) == 4 {
		//fmt.Printf("IS ALSO Invalid, off the bat\n")
		return tempPath, false
	}
	if start.IsEqualTo(end) {
		//fmt.Printf("Found it!\n")
		tempPath = append(tempPath, start)
		return tempPath, true
	}
	igd.Imat[start.Y][start.X] = 3
	tempPath = append(tempPath, start)
	for _, d := range directions {
		newNum := num + 1
		newPoint := start.AddCoords(d)
		tempPath2, isValid := igd.FindPather_BreadthFirst(newPoint, end, newNum, direct, tempPath)
		if isValid {
			//fmt.Printf("IS VALID! %d %d %d\n", i, newPoint.X, newPoint.Y)
			igd.Imat[start.Y][start.X] = 1
			return tempPath2, true
		}
		// else {
		// 	fmt.Printf("IS InValid! %d %d %d\n", i, newPoint.X, newPoint.Y)
		// 	return tempPath2, false
		// }
	}
	igd.Imat[start.Y][start.X] = 1
	tempPath = (tempPath)[:len(tempPath)-1]
	return tempPath, false
}

func (igd *IntegerGridManager) FindPath(n int) {
	if igd.PFinder.HasFalsePos {
		igd.PFinder.FalsePos = make(CoordList, 0)
		igd.PFinder.HasFalsePos = false
	}
	temp, isDone := make(CoordList, 0), false
	// temp = append(temp, igd.PFinder.StartPos)
	//fmt.Printf("STARTING UP FINDPATH\n")
	igd.PFinder.FalsePos, isDone = igd.FindPather_BreadthFirst(igd.PFinder.Cursor.Position, igd.PFinder.EndPos, 0, n, temp)
	if isDone {
		//fmt.Printf("IS Truely DONE\n")
		igd.PFinder.HasFalsePos = true
	} else {
		//fmt.Printf("FAILURE\n")
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
			fmt.Printf("FAILURE\n")
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
			dire = 7
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
		} else if yy < 0 {
			dire = 0
		}
	} else if xx < 0 {
		if yy > 0 {
			if (xx * -1) < yy {
				dire = 2
			} else {
				dire = 5
			}
		} else if yy == 0 {
			dire = 5
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
