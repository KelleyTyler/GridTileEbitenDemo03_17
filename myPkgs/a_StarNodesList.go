package mypkgs

import (
	"fmt"
	"math"
)

/*

	here I'm implementing some helper functions for node lists;
	which are dumber than linked lists but still good;

*/

// func (nodel []Node) SortByFValue() {

// }
func NodesAr_Sort_ByF_Value(nodel []*Node) {
	if len(nodel) > 1 {
		for range nodel {
			for i := 1; i < len(nodel); i++ {
				if nodel[i] != nil {
					if nodel[i-1] != nil {
						if nodel[i].MCost_Sum < nodel[i-1].MCost_Sum {
							tNode := nodel[i]
							nodel[i] = nodel[i-1]
							nodel[i-1] = tNode
						}
					} else {
						fmt.Printf("AT %d-1 WE'VE GOT NILS\n", i)
					}
				} else {
					fmt.Printf("AT %d WE'VE GOT NILS\n", i)
				}
			}
		}
	} else {
		//fmt.Printf("SMALLER THAN 1\n")
	}

	//fmt.Printf("END SORT _______\n")
}
func NodesAr_Sort_ByF_Value_OLD(nodel []*Node) []*Node {
	temp := make([]*Node, len(nodel))
	copy(temp, nodel)
	if len(temp) > 4 {
		for range temp {
			for i := 1; i < len(temp)-1; i++ {
				if temp[i].MCost_Sum > temp[i-1].MCost_Sum {
					tNode := temp[i]
					temp[i] = temp[i-1]
					temp[i-1] = tNode
				}
			}
		}
	} else {
		fmt.Printf("SMALLER THAN 4\n")
	}
	return temp
}
func NodesAr_PopFromFront(nodel []*Node) (*Node, []*Node) {
	tempAr := make([]*Node, 0)
	// copy(temp, nodel)
	var temp *Node
	if len(nodel) > 0 {
		temp = nodel[0]
		if len(nodel) > 1 {
			for i := 1; i < len(nodel)-1; i++ {
				tempAr = append(tempAr, nodel[i])
			}
		}

	} else {
		fmt.Printf("NODEL IS LESS THAN ZERO\n")

	}
	fmt.Printf("TEMPAR: %d NODEL %d ", len(tempAr), len(nodel))
	if temp != nil {
		fmt.Printf("TEMP EXISTS\n")
	} else {
		fmt.Printf("TEMP DOESN'T EXIST\n")
	}
	nodel = append(nodel[:0], nodel[0+1:]...)
	//nodel = tempAr
	return temp, tempAr
}

func NodesAr_UpdateOnDistanceToParent(nodel []*Node) {
	for _, n := range nodel {
		if n.ParentPTR != nil {
			n.SetCostToParent()
		}
	}
}

func NodesAr_PrintArray(nodel []*Node) {
	for _, n := range nodel {
		fmt.Printf("%s\n", n.ToString())
	}
}

func NodesAr_RemoveDuplicates(nodel []*Node) []*Node {
	temp := make([]*Node, 0)
	// copy(temp, nodel)
	for _, a := range nodel {
		test := true
		for _, b := range nodel {
			if b.Postion.IsEqualTo(a.Postion) && a.ParentPTR.Postion.IsEqualTo(b.Postion) {
				test = false
			}
		}
		if test {
			temp = append(temp, a)
		}
	}
	nodel = temp
	return temp
}
func NodesAr_RemoveByNode(nodel []*Node, node *Node) []*Node {
	// temp := make([]*Node, len(nodel))
	// copy(temp, nodel)
	tee := 0
	for i, a := range nodel {
		if a.Postion.IsEqualTo(node.Postion) {
			// if node.ParentPTR != a.ParentPTR {
			// 	temp = append(temp, a)
			// }
			tee = i

		}
	}
	nodel = append(nodel[:tee], nodel[tee+1:]...)
	//nodel = temp
	return nodel
}

// func

func (imat *IntMatrix) NodeAr_GetNeighbors4(Parent *Node, buffer [4]int, startpoint, endPoint CoordInts) []*Node {
	retList := make([]*Node, 0)
	templist, _, _ := imat.GetNeighbors4(Parent.Postion, buffer)
	for _, c := range templist {
		if !c.IsEqualTo(Parent.Postion) {
			temp := InitNode(Parent.Postion, c, endPoint)
			temp.ParentPTR = Parent

			xx, yy := temp.Postion.GetDifferenceInInts(endPoint)
			temp.MCost_toEnd = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
			xx, yy = temp.Postion.GetDifferenceInInts(Parent.Postion)
			temp.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
			temp.MCost_toStart = temp.Postion.GetOverallManhattanDistance(startpoint)

			temp.MCost_Sum = temp.MCost_toEnd + temp.MCost_toStart
			if imat.IsValid(temp.Postion) {
				temp.ValueOnCoord = imat.GetCoordVal(temp.Postion)
			}
			retList = append(retList, temp)
		}
	}

	// retList = append(retList)
	return retList
}

func NodeAr_Contains(nodel []*Node, nod *Node) bool {
	for _, c := range nodel {
		if c.Postion.IsEqualTo(nod.Postion) {
			return true
		}
	}
	return false
}
func NodeAr_Contains_what(nodel []*Node, nod *Node) (*Node, bool) {
	for _, c := range nodel {
		if c.Postion.IsEqualTo(nod.Postion) {
			return c, true
		}
	}
	return nil, false
}
func (imat *IntMatrix) NodeAr_GetNeighbors4FILTERED(Parent *Node, buffer [4]int, WallValues []int, startpoint, endPoint CoordInts) []*Node {
	retList := make([]*Node, 0)
	templist, _, _ := imat.GetNeighbors4(Parent.Postion, buffer)
	for _, c := range templist {
		if !c.IsEqualTo(Parent.Postion) {
			if !imat.IsCoordValueInArrayOfValues(c, WallValues) {
				temp := InitNode(Parent.Postion, c, endPoint)
				temp.ParentPTR = Parent

				xx, yy := temp.Postion.GetDifferenceInInts(endPoint)
				temp.MCost_toEnd = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
				xx, yy = temp.Postion.GetDifferenceInInts(Parent.Postion)
				temp.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
				temp.MCost_Sum = temp.MCost_toEnd + temp.MCost_toParent
				if imat.IsValid(temp.Postion) {
					temp.ValueOnCoord = imat.GetCoordVal(temp.Postion)
				}
				retList = append(retList, temp)
			}
		}
	}

	// retList = append(retList)
	return retList
}
func (imat *IntMatrix) NodeAr_GetNeighbors4_Filtered_MD(Parent *Node, buffer [4]int, WallValues []int, startpoint, endPoint CoordInts) []*Node {
	retList := make([]*Node, 0)
	templist, _, _ := imat.GetNeighbors4(Parent.Postion, buffer)
	for _, c := range templist {
		if !c.IsEqualTo(Parent.Postion) {
			if !imat.IsCoordValueInArrayOfValues(c, WallValues) {
				temp := InitNode(Parent.Postion, c, endPoint)
				temp.ParentPTR = Parent

				// xx, yy := temp.Postion.GetDifferenceInInts(endPoint)
				// temp.MCost_toEnd = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
				// xx, yy = temp.Postion.GetDifferenceInInts(Parent.Postion)
				// temp.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
				temp.MCost_toStart = temp.Postion.GetDistanceIntTimesTen(startpoint)
				temp.MCost_toEnd = temp.Postion.GetDistanceIntTimesTen(endPoint)
				temp.MCost_toParent = temp.Postion.GetDistanceIntTimesTen(Parent.Postion)
				temp.MCost_Sum = temp.MCost_toEnd + temp.MCost_toStart
				if imat.IsValid(temp.Postion) {
					temp.ValueOnCoord = imat.GetCoordVal(temp.Postion)
				}
				retList = append(retList, temp)
			}
		}
	}

	// retList = append(retList)
	return retList
}
func (imat *IntMatrix) NodeAr_GetNeighbors8_Filtered_Manhattan(Parent *Node, buffer [4]int, WallValues []int, startpoint, endPoint CoordInts) []*Node {
	retList := make([]*Node, 0)
	// templist, _, _ := imat.GetNeighbors4(Parent.Postion, buffer)
	templist, _, _ := imat.GetNeighbors8(Parent.Postion, buffer)
	for _, c := range templist {
		if !c.IsEqualTo(Parent.Postion) {
			if !imat.IsCoordValueInArrayOfValues(c, WallValues) {
				temp := InitNode(Parent.Postion, c, endPoint)
				temp.ParentPTR = Parent

				xx, yy := temp.Postion.GetDifferenceInInts(endPoint)
				temp.MCost_toEnd = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
				xx, yy = temp.Postion.GetDifferenceInInts(Parent.Postion)
				temp.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
				// temp.MCost_toStart = temp.Postion.GetDistanceIntTimesTen(startpoint)
				temp.MCost_toStart = temp.Postion.GetOverallManhattanDistance(startpoint)
				temp.MCost_Sum = temp.MCost_toEnd + temp.MCost_toParent
				if imat.IsValid(temp.Postion) {
					temp.ValueOnCoord = imat.GetCoordVal(temp.Postion)
				}
				retList = append(retList, temp)
			}
		}
	}

	// retList = append(retList)
	return retList
}
func (imat *IntMatrix) NodeAr_GetNeighbors8_Filtered_MD(Parent *Node, buffer [4]int, WallValues []int, endPoint CoordInts) []*Node {
	retList := make([]*Node, 0)
	// templist, _, _ := imat.GetNeighbors4(Parent.Postion, buffer)
	templist, _, _ := imat.GetNeighbors8(Parent.Postion, buffer)
	for _, c := range templist {
		if !c.IsEqualTo(Parent.Postion) {
			if !imat.IsCoordValueInArrayOfValues(c, WallValues) {
				temp := InitNode(Parent.Postion, c, endPoint)
				temp.ParentPTR = Parent

				temp.MCost_toEnd = temp.Postion.GetDistanceIntTimesTen(endPoint)
				temp.MCost_toParent = temp.Postion.GetDistanceIntTimesTen(Parent.Postion)

				// temp.MCost_toEnd = temp.Postion.GetOverallManhattanDistance(endPoint)
				// temp.MCost_toParent = temp.Postion.GetOverallManhattanDistance(Parent.Postion)
				temp.MCost_Sum = temp.MCost_toEnd + temp.MCost_toParent
				if imat.IsValid(temp.Postion) {
					temp.ValueOnCoord = imat.GetCoordVal(temp.Postion)
				}
				retList = append(retList, temp)
			}
		}
	}

	// retList = append(retList)
	return retList
}

// //=====> Have a value

// type Node_PriorityQueue []*Node

// func (nodel Node_PriorityQueue) Sort_ByF_Value() {
// 	// temp := make(Node_PriorityQueue, len(*nodel))
// 	// copy(temp, *nodel)
// 	for range nodel {
// 		for i := 1; i < len(nodel)+1; i++ {
// 			if nodel[i].MCost_Sum > nodel[i-1].MCost_Sum {
// 				tNode := nodel[i]
// 				nodel[i] = nodel[i-1]
// 				nodel[i-1] = tNode
// 			}
// 		}
// 	}
// }
// func (nodel *Node_PriorityQueue) Sort_ByF_Value_return2() Node_PriorityQueue {
// 	temp := make(Node_PriorityQueue, len(*nodel))
// 	copy(temp, *nodel)
// 	for range temp {
// 		for i := 1; i < len(temp)+1; i++ {
// 			if temp[i].MCost_Sum > temp[i-1].MCost_Sum {
// 				tNode := temp[i]
// 				temp[i] = temp[i-1]
// 				temp[i-1] = tNode
// 			}
// 		}
// 	}
// 	return temp
// }

// // func (nodel *Node_PriorityQueue) PopFromFront() *Node {
// // 	tempAr := make(Node_PriorityQueue, len(*nodel))
// // 	// copy(temp, nodel)
// // 	var temp *Node
// // 	if len(*nodel) > 0 {
// // 		temp = nodel.GetValueAt(0)
// // 		if len(*nodel) > 1 {
// // 			for i := 1; i < len(*nodel)-1; i++ {
// // 				tempAr = append(tempAr, nodel.GetValueAt(i))
// // 			}
// // 		}
// // 	}
// // 	nodel = &tempAr
// // 	return temp
// // }

// func (nodel *Node_PriorityQueue) PopFromFront_return2() (*Node, Node_PriorityQueue) {
// 	tempAr := make(Node_PriorityQueue, len(*nodel))
// 	// copy(temp, nodel)
// 	var temp *Node
// 	if len(*nodel) > 0 {
// 		temp = nodel.GetValueAt(0)
// 		if len(*nodel) > 1 {
// 			for i := 1; i < len(*nodel)-1; i++ {
// 				tempAr = append(tempAr, nodel.GetValueAt(i))
// 			}
// 		}
// 	}
// 	return temp, tempAr
// }

// func (nodel *Node_PriorityQueue) UpdateOnDistanceToParent() {
// 	for _, n := range *nodel {
// 		if n.ParentPTR != nil {
// 			n.SetCostToParent()
// 		}
// 	}
// }
// func (nodel Node_PriorityQueue) GetValueAt(index int) *Node {
// 	return nodel[index]
// }
