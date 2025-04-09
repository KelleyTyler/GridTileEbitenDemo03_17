package mypkgs

import (
	"fmt"
	"log"
	"math"
)

type Node struct {
	Postion        CoordInts
	ParentPTR      *Node
	ChildPTR       *Node
	Index          int
	MCost_Sum      int //equal to the sum of G and H
	MCost_toParent int
	MCost_toStart  int //Movement Cost from StartNode to Positon
	MCost_toEnd    int //Movement Cost from Position to End// will be -1 while not working
	ValueOnCoord   int
}

/*
This assumes the point curr, Start, End are all valid for whatever "IntMatrix" you might be using;
*/
func InitNode(Start, Curr, End CoordInts) *Node {
	x0, y0 := Start.GetDifferenceInInts(Curr)
	x1, y1 := Curr.GetDifferenceInInts(End)

	a0 := int(math.Abs(float64(x0)) + math.Abs(float64(y0)))
	a1 := int(math.Abs(float64(x1)) + math.Abs(float64(y1)))
	node := Node{Index: 0, Postion: Curr, MCost_toStart: a0, MCost_toEnd: a1, ValueOnCoord: -1}
	return &node
}
func NodeTest(imat IntMatrix) {
	fmt.Printf("TEST TEST TEST TEST------------------------->\n")
	nodes := InitNode(CoordInts{X: 0, Y: 0}, CoordInts{X: 0, Y: 1}, CoordInts{10, 10})

	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 7, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 6, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 1, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 2, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 2, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 5, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 3, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 2, Y: 1}, CoordInts{10, 10})

	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 7, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 7, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 6, Y: 1}, CoordInts{10, 10})
	nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 6, Y: 1}, CoordInts{10, 10})
	// nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 7, Y: 1}, CoordInts{10, 10})
	// nodes.PushToBack(CoordInts{X: 0, Y: 0}, CoordInts{X: 7, Y: 1}, CoordInts{10, 10})
	// fmt.Printf("------------------------------------- %d \n", nodes.GetLength())
	// nodes.PrintNodesFrontToBack()
	// // nodes = nodes.RemoveConsecutivePositionDoubles()
	nodes.RemoveConsecutivePositionDoubles()
	// fmt.Printf("------------------------------------- %d \n", nodes.GetLength())
	// nodes.PrintNodesFrontToBack()
	quant := nodes.GetStart()
	// nodes.RemoveChild()
	// fmt.Printf("------------------------------------- %d \n", nodes.GetLength())
	// nodes.PrintNodesFrontToBack()
	// fmt.Printf("-------------------------------------\n")

	// fmt.Printf("-------------------------------------\n")
	// quant.MoveToFront()
	// quant.PrintNodesBackToFront()
	// fmt.Printf("------------------MOVE TO BACK-------------------\n")
	// quant.MoveToBack()
	// quant.PrintNodesBackToFront()

	// fmt.Printf("------------------TICK-------------------\n")

	// // nodes = nodes.GetStart()
	// // nodes = nodes.GetStart()
	// fmt.Printf("------------------------------------- %d \n", nodes.GetLength())
	// nodes.UpdateIndexAndValueOnCoord(imat)
	// nodes.PrintNodesFrontToBack()

	// nodes.MoveToBack()

	// nodes.UpdateIndexAndValueOnCoord(imat)

	// // quant = quant.GetStart()
	// fmt.Printf("------------------------------------- %d \n", quant.GetLength())
	// nodes.PrintNodesBackToFront()

	// fmt.Printf("------------------------------------- %d \n", quant.GetLength())
	// quant.PrintNodesFrontToBack()
	quant.UpdateIndexAndValueOnCoord(imat)
	quant = quant.GetStart()
	fmt.Printf("------------------------------------- %d %d\n", quant.GetLength(), quant.Index)

	quant.PrintNodesFrontToBack()

	quant.SortOnDistanceToEndDesc()
	quant = quant.GetStart()
	fmt.Printf("------------------------------------- %d %d\n", quant.GetLength(), quant.Index)

	quant.PrintNodesFrontToBack()

	quant.SortOnDistanceToEndDesc()
	quant = quant.GetStart()
	fmt.Printf("----------222----------------------- %d %d\n", quant.GetLength(), quant.Index)
	quant.PrintNodesFrontToBack()
	var mint *Node = nil
	mint, quant, _ = quant.PopFromLocationGetHead(5)
	fmt.Printf("----------Removal----------------------- %d %d\n", quant.GetLength(), quant.Index)
	quant.PrintNodesFrontToBack()
	fmt.Printf("----------MINT----------------------- %d %d\n", mint.GetLength(), mint.Index)
	mint.PrintNodesBackToFront()
	_, quant, _ = quant.PopFromHead()

	fmt.Printf("----------POP----------------------- %d %d\n", quant.GetLength(), quant.Index)
	quant.PrintNodesFrontToBack()

	// nodes.PrintNodesFrontToBack()
	// nodes2.PrintNodesFrontToBack()
}

func (node *Node) Init(point CoordInts) {

}
func (node *Node) InsertInto(Start, Curr, End CoordInts, index int) {
	temp00, _ := node.GetFromIndex(index)
	temp01 := InitNode(Start, Curr, End)
	if temp00.ChildPTR != nil {
		temp02 := temp00.ChildPTR
		temp00.ChildPTR = temp01
		temp01.ParentPTR = temp00
		temp01.ChildPTR = temp02
		temp02.ParentPTR = temp01
	}

	temp00 = temp00.GetStart()
	temp00.UpdateIndex()
}
func (node *Node) PushToFront(Start, Curr, End CoordInts) *Node {
	temp := InitNode(Start, Curr, End)
	node.ParentPTR = temp
	temp.ChildPTR = node
	temp.UpdateIndex()
	return temp
}
func (node *Node) PushToBack(Start, Curr, End CoordInts) {
	if node.ChildPTR != nil {
		node.ChildPTR.PushToBack(Start, Curr, End)
	} else {
		temp := InitNode(Start, Curr, End)
		node.ChildPTR = temp
		temp.ParentPTR = node
	}

}

/*Wrapper for a process that starts with updateIndexInternally*/
func (node *Node) UpdateIndex() {
	node.updateIndexInternally(0, false)
}
func (node *Node) updateIndexInternally(num int, active bool) {
	if !active {
		if node.ParentPTR != nil {
			node.ParentPTR.updateIndexInternally(num, false)
		} else {
			node.Index = num
			node.ChildPTR.updateIndexInternally(num+1, true)
			xx, yy := node.ChildPTR.Postion.GetDifferenceInInts(node.Postion)
			node.ChildPTR.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
		}
	} else {
		node.Index = num
		if node.ChildPTR != nil {
			node.ChildPTR.updateIndexInternally(num+1, true)
			xx, yy := node.ChildPTR.Postion.GetDifferenceInInts(node.Postion)
			node.ChildPTR.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
		}
	}
}

func (node *Node) UpdateIndexAndValueOnCoord(imat IntMatrix) {
	node.updateIndexAndValueOnCoord(imat, false)
}
func (node *Node) updateIndexAndValueOnCoord(imat IntMatrix, active bool) {
	if !active {
		if node.ParentPTR != nil {
			node.ParentPTR.updateIndexAndValueOnCoord(imat, false)
		} else {
			node.Index = 0
			if imat.IsValid(node.Postion) {
				node.ValueOnCoord = imat.GetCoordVal(node.Postion)
				if node.ChildPTR != nil {
					xx, yy := node.ChildPTR.Postion.GetDifferenceInInts(node.Postion)
					node.ChildPTR.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
					node.ChildPTR.updateIndexAndValueOnCoord(imat, true)
				}
			}
		}
	} else {
		node.Index = node.ParentPTR.Index + 1
		if imat.IsValid(node.Postion) {
			node.ValueOnCoord = imat.GetCoordVal(node.Postion)
			if node.ChildPTR != nil {
				xx, yy := node.ChildPTR.Postion.GetDifferenceInInts(node.Postion)
				node.ChildPTR.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
				node.ChildPTR.updateIndexAndValueOnCoord(imat, true)
			}
		}
	}
}

func (node *Node) SortBy() {

}

func (node *Node) SetCostToParent() {
	if node.ParentPTR != nil {
		xx, yy := node.ParentPTR.Postion.GetDifferenceInInts(node.Postion)
		node.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
	} else {
		node.MCost_toParent = -1
	}
}
func (node *Node) SetCostToParent_Cascading() {
	if node.ParentPTR != nil {
		xx, yy := node.ParentPTR.Postion.GetDifferenceInInts(node.Postion)
		node.MCost_toParent = int(math.Abs(float64(xx)) + math.Abs(float64(yy)))
		node.ParentPTR.SetCostToParent_Cascading()
	} else {
		node.MCost_toParent = -1
	}
}

// func (node *Node) Pop() *Node {

// }

func (node *Node) GetLength() int {
	if node.ChildPTR != nil {
		return node.ChildPTR.GetLength() + 1
	} else {
		return 1
	}
}

func (node *Node) GetFromIndex(desiredIndex int) (*Node, bool) {

	if desiredIndex == node.Index {
		return node, true
	} else {
		if desiredIndex > node.Index {
			if node.ChildPTR == nil {
				return nil, false
			} else {
				return node.ChildPTR.GetFromIndex(desiredIndex)
			}
		} else if desiredIndex < node.Index {
			if node.Index == 0 {
				return node, true
			} else {
				return node.ChildPTR.GetFromIndex(desiredIndex)
			}
		} else {
			return node, false
		}
	}
}

func (node *Node) GetFromPosition(desiredIndex int) (*Node, bool) {
	temp := node.GetStart()
	return temp.getFromPositionHelper(desiredIndex, 0)
}

func (node *Node) getFromPositionHelper(desiredIndex, currIndex int) (*Node, bool) {
	if desiredIndex == currIndex {
		return node, true
	} else {
		if node.ChildPTR != nil {
			return node.ChildPTR.getFromPositionHelper(desiredIndex, currIndex+1)
		} else {
			return node, false
		}
	}
}

func (node *Node) GetStart() *Node {
	if node.ParentPTR != nil {
		return node.ParentPTR.GetStart()
	} else {
		return node
	}
}
func (node *Node) GetEnd() *Node {
	if node.ChildPTR != nil {
		return node.ChildPTR.GetEnd()
	} else {
		return node
	}
}

// func (node *Node) PopFromFront() *Node {

func (node *Node) RemoveFromPlace(place int) {
	if place < node.GetLength() {
		if place != 0 {
			temp, tBool := node.GetFromPosition(place - 1)
			if tBool {
				if temp.ChildPTR.ChildPTR != nil {
					tempChildChild := temp.ChildPTR.ChildPTR
					temp.ChildPTR.ParentPTR = nil
					temp.ChildPTR.ChildPTR = nil
					tempChildChild.ParentPTR = temp
					temp.ChildPTR = tempChildChild
				} else {
					temp.ChildPTR.ParentPTR = nil
					temp.ChildPTR.ChildPTR = nil
					temp.ChildPTR = nil
				}
			} else {

			}
		} else {
			if node.ParentPTR == nil {
				if node.ChildPTR != nil {
					temp := node.ChildPTR
					temp.ParentPTR = nil
					node.ChildPTR = nil
					node = temp
				}
			}
		}
	}
}
func (node *Node) RemoveFromPlace_GetHead(place int) (*Node, bool) {
	tempHead := node.GetStart()
	leng := tempHead.GetLength()
	if place == 0 {
		if tempHead.ChildPTR != nil {
			temp02 := tempHead.ChildPTR
			tempHead.ChildPTR = nil
			tempHead = temp02
			tempHead.ParentPTR = nil //
			return tempHead, true
		} else {
			return nil, true
		}
	} else {
		if leng > place {
			temp02, b1 := tempHead.GetFromPosition(place - 1)
			if b1 {
				if temp02.ChildPTR.ChildPTR != nil {
					tempChildChild := temp02.ChildPTR.ChildPTR
					temp02.ChildPTR.ParentPTR = nil
					temp02.ChildPTR.ChildPTR = nil
					tempChildChild.ParentPTR = temp02
					temp02.ChildPTR = tempChildChild
				} else {
					temp02.ChildPTR.ParentPTR = nil
					temp02.ChildPTR.ChildPTR = nil
					temp02.ChildPTR = nil
				}
				return tempHead, true
			} else {
				return tempHead, false
			}
		} else {
			return tempHead, false
		}
	}

}
func (node *Node) PopFromLocationGetHead(place int) (Removed_node *Node, NewHead *Node, Success bool) {
	tempHead := node.GetStart()
	leng := tempHead.GetLength()
	if place == 0 {
		if tempHead.ChildPTR != nil {
			temp02 := tempHead.ChildPTR
			tempHead.ChildPTR = nil
			temp02.ParentPTR = nil
			// tempHead.ParentPTR = nil //
			return tempHead, temp02, true
		} else {
			return nil, nil, true
		}
	} else {
		if leng > place {
			temp02, b1 := tempHead.GetFromPosition(place - 1)
			if b1 {
				tempChild := temp02.ChildPTR
				if temp02.ChildPTR.ChildPTR != nil {
					tempChildChild := temp02.ChildPTR.ChildPTR
					temp02.ChildPTR.ParentPTR = nil
					temp02.ChildPTR.ChildPTR = nil
					tempChildChild.ParentPTR = temp02
					temp02.ChildPTR = tempChildChild
				} else {
					temp02.ChildPTR.ParentPTR = nil
					temp02.ChildPTR.ChildPTR = nil
					temp02.ChildPTR = nil
				}
				return tempChild, tempHead, true
			} else {
				return nil, tempHead, false
			}
		} else {
			return nil, tempHead, false
		}
	}

}

func (node *Node) PopFromHead() (*Node, *Node, bool) {
	tempHead := node.GetStart()
	oldHead := node.GetStart()
	if tempHead.ChildPTR != nil {
		tempHead = tempHead.ChildPTR
		oldHead.ChildPTR = nil
		tempHead.ParentPTR = nil
		return oldHead, tempHead, true
	} else {
		tempHead = nil
		if oldHead != nil {
			return oldHead, tempHead, true
		} else {
			return nil, nil, false
		}
	}
}

// }
func (node *Node) RemoveFromIndex(index int) {
	tNodeA := node.GetStart()
	tNode0, tempBool := node.GetFromIndex(index)
	if tempBool && tNode0.Index != node.Index {
		if tNode0.ChildPTR != nil && tNode0.ParentPTR != nil {
			tNode0.ChildPTR.ParentPTR = tNode0.ParentPTR
			tNode0.ParentPTR.ChildPTR = tNode0.ChildPTR
			tNode0.ParentPTR = nil
			tNode0.ChildPTR = nil
		} else {
			if tNode0.ParentPTR != nil {
				tNode0.ParentPTR.ChildPTR = nil
				//tNode1 := tNode0.ParentPTR
				tNode0.ParentPTR = nil

			} else if tNode0.ChildPTR != nil {

				tNode0.ChildPTR.ParentPTR = nil
			} else {
				//node = nil
			}
		}
	} else {
		fmt.Printf("ERROR YOU CAN'T DO THAT!")
		log.Fatal("ERROR")
	}
	tNodeA.UpdateIndex()
}

func (n00 *Node) IsOnSamePosition(n01 *Node) bool {
	return n00.Postion.IsEqualTo(n01.Postion)
}

// func (n00 *Node) Is_On_SamePointOrCloserToParent(n01 *Node) (bool, int) {

// }
func (node *Node) ToString() string {
	a := (node.ParentPTR != nil)
	b := (node.ChildPTR != nil)
	strng := fmt.Sprintf("NODE: %3d \tCoord: %3d %3d|to Start:%3d toParent:%3d toEnd: %3d\t has Parent:%5t hasChild:%5t \t COORDVAL: %d", node.Index, node.Postion.X, node.Postion.Y, node.MCost_toStart, node.MCost_toParent, node.MCost_toEnd, a, b, node.ValueOnCoord)
	if b {
		xx, yy := node.Postion.GetDifferenceInInts(node.ChildPTR.Postion)
		if xx > 0 {
			strng += fmt.Sprintf("\t %d MOVES: EAST", node.ChildPTR.Index)
		} else if xx < 0 {
			strng += fmt.Sprintf("\t %d MOVES: West", node.ChildPTR.Index)
		}

		if yy > 0 {
			strng += fmt.Sprintf("\t %d MOVES: South", node.ChildPTR.Index)
		} else if yy < 0 {
			strng += fmt.Sprintf("\t %d MOVES: North", node.ChildPTR.Index)
		}
	}
	return strng
}

func (node *Node) PrintNodesFrontToBack() {
	fmt.Printf("%s\n", node.ToString())

	if node.ChildPTR != nil {
		node.ChildPTR.PrintNodesFrontToBack()
		// fmt.Printf("\n")
	} else {
		fmt.Printf("------->\n")
	}

}
func (node *Node) PrintNodesBackToFront() {
	fmt.Printf("%s\n", node.ToString())
	if node.ParentPTR != nil {
		node.ParentPTR.PrintNodesBackToFront()
		//fmt.Printf("---------\n")
	} else {
		fmt.Printf("------->\n")
	}

}

func (node *Node) RemoveConsecutivePositionDoubles() {
	// temp := node.GetStart()
	node.remove_Consecutive_Position_Doubles(false)
	// return temp
}

func (node *Node) remove_Consecutive_Position_Doubles(isActive bool) {
	if !isActive {
		if node.ParentPTR != nil {
			node.remove_Consecutive_Position_Doubles(false)
		} else {
			if node.Postion.IsEqualTo(node.ChildPTR.Postion) {
				node.RemoveChild()
				if node.ChildPTR != nil {
					node.ChildPTR.remove_Consecutive_Position_Doubles(true)
				}
			}
			if node.ChildPTR != nil {
				node.ChildPTR.remove_Consecutive_Position_Doubles(true)
			}
			if node.ChildPTR != nil && node.Postion.IsEqualTo(node.ChildPTR.Postion) {
				node.RemoveChild()
			}
		}
	} else {
		if node.Postion.IsEqualTo(node.ParentPTR.Postion) {
			if node.Postion.IsEqualTo(node.ChildPTR.Postion) {
				node.RemoveChild()
				if node.ChildPTR != nil {
					node.ChildPTR.remove_Consecutive_Position_Doubles(true)
				}
			} else if node.ParentPTR != nil {
				fmt.Printf("HAS PARENT %d %d \n", node.Postion.X, node.Postion.Y)
			}
			if node.ChildPTR != nil {
				node.ChildPTR.remove_Consecutive_Position_Doubles(true)
			}
		} else {
			if node.ChildPTR != nil {
				fmt.Printf("TEST000!\n")
				if node.Postion.IsEqualTo(node.ChildPTR.Postion) {
					node.RemoveChild()
				}
				if node.ChildPTR != nil {
					node.ChildPTR.remove_Consecutive_Position_Doubles(true)
				}
			}
			// if node.ChildPTR != nil && node.Postion.IsEqualTo(node.ChildPTR.Postion) {
			// 	fmt.Printf("TEST!\n")

			// }
		}

	}
}
func (node *Node) MoveToFront() {
	for node.ParentPTR != nil {
		node.SwapWithParent()
	}
}
func (node *Node) MoveToBack() {
	for node.ChildPTR != nil {
		node.SwapWithChild()
	}
}

func (node *Node) SwapWithParent() {
	if node.ParentPTR != nil {
		parent := node.ParentPTR
		if node.ChildPTR != nil {
			parent.ChildPTR = node.ChildPTR
		} else {
			parent.ChildPTR = nil
		}
		node.ChildPTR = parent
		if parent.ParentPTR != nil {
			parent.ParentPTR.ChildPTR = node
			node.ParentPTR = parent.ParentPTR
		} else {
			node.ParentPTR = nil
		}
		parent.ParentPTR = node
	}
}
func (node *Node) SwapWithChild() {
	if node.ChildPTR != nil {
		child := node.ChildPTR
		if node.ParentPTR != nil {
			node.ParentPTR.ChildPTR = child
			child.ParentPTR = node.ParentPTR
			node.ParentPTR = child
		} else {
			child.ParentPTR = nil
		}
		node.ParentPTR = child
		if child.ChildPTR != nil {
			node.ChildPTR = child.ChildPTR
		} else {
			node.ChildPTR = nil
		}
		child.ChildPTR = node
	}
}
func (node *Node) RemoveChild() {
	if node.ChildPTR != nil {
		if node.ChildPTR.ChildPTR != nil {
			// tempChild := node.ChildPTR
			tempChildChild := node.ChildPTR.ChildPTR
			node.ChildPTR.ChildPTR.ParentPTR = node
			node.ChildPTR.ParentPTR = nil
			node.ChildPTR.ChildPTR = nil
			node.ChildPTR = tempChildChild
			tempChildChild = nil
		} else {
			//tempChild:=node.ChildPTR
			node.ChildPTR.ParentPTR = nil
			node.ChildPTR = nil
		}
	}
}

func (node *Node) SortOnDistanceToEndDesc() {

}
func (node *Node) Cycler() {
	temp := node.GetStart()
	lenger := temp.GetLength()
	for range 2 {
		temp := temp.GetStart()
		for range lenger {
			temp.SwapWithChild()
		}

	}
}
func (node *Node) Sort_D_2End_Tick() {
	if node.ChildPTR != nil {
		if node.ChildPTR.MCost_toEnd > node.MCost_toEnd {
			node.SwapWithChild()
			if node.ChildPTR != nil {
				node.ChildPTR.Sort_D_2End_Tick()
			}
		} else {
			if node.ChildPTR != nil {
				//node.ChildPTR.Sort_D_2End_Tick()
			}
		}
	}
}
func (node *Node) Sort_D_3End_Tick() {
	if node.ParentPTR != nil {
		if node.ParentPTR.MCost_toEnd < node.MCost_toEnd {
			node.SwapWithParent()
			if node.ParentPTR != nil {
				node.ParentPTR.Sort_D_3End_Tick()
			}
		} else {
			if node.ParentPTR != nil {
				node.ParentPTR.Sort_D_3End_Tick()
			}
		}
	}
}

func (node *Node) SnakeMove(newPos CoordInts, addToTail bool) {
	node.snakeMoveHelper(newPos, addToTail, false)
}

func (node *Node) snakeMoveHelper(newPos CoordInts, addToTail, active bool) {
	if active {
		temp := node.Postion
		node.Postion = newPos
		if node.ChildPTR != nil {
			node.ChildPTR.snakeMoveHelper(temp, addToTail, true)
		}
	} else {
		if node.ParentPTR != nil {
			node.ParentPTR.snakeMoveHelper(newPos, addToTail, false)
		} else {
			if addToTail {
				node.PushToBack(CoordInts{2, 2}, CoordInts{-1, -1}, CoordInts{6, 6})
			}
			temp := node.Postion
			num := node.GetNumberOfOccurances_Position(newPos)
			// fmt.Printf("%d\n", num)
			if num < 1 {
				node.Postion = newPos
				if node.ChildPTR != nil {
					node.ChildPTR.snakeMoveHelper(temp, addToTail, true)
				}
			} else {
				//fmt.Printf("WHAT?\n")
			}
		}
	}
}

func (node *Node) GetNumberOfOccurances_Position(pos CoordInts) int {
	return node.get_num_occurances_position_helper(pos, false, 0)
}
func (node *Node) get_num_occurances_position_helper(pos CoordInts, active bool, count int) int {
	if active {

		if node.Postion.IsEqualTo(pos) {
			if node.ChildPTR != nil {
				return node.ChildPTR.get_num_occurances_position_helper(pos, true, count+1)
			} else {
				return count
			}
		} else {
			if node.ChildPTR != nil {
				return node.ChildPTR.get_num_occurances_position_helper(pos, true, count)
			} else {
				return count
			}
		}
	} else {
		if node.ParentPTR != nil {
			return node.get_num_occurances_position_helper(pos, false, count)
		} else {
			if node.Postion.IsEqualTo(pos) {
				if node.ChildPTR != nil {
					return node.ChildPTR.get_num_occurances_position_helper(pos, true, count+1)
				} else {
					return 1
				}
			} else {
				if node.ChildPTR != nil {
					return node.ChildPTR.get_num_occurances_position_helper(pos, true, count)
				} else {
					return 0
				}
			}
		}
	}
}

// // pick a point that is one off from the tail but which does not interact with the others
// func (node *Node) AddSnakeTail() {
// 	temp := node.GetEnd()
// 	if()
// }
