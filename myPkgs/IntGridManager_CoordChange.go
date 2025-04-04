package mypkgs

import (
	"fmt"
	"image/color"
)

/*
	the idea here is to have a means by which the integer grid manager is better able to recieve and handle changes,
	preferably without resorting to a total redraw each and every
	time a tile is changed or altered;
*/

type CoordIntChange struct {
	coord    CoordInts
	changeTo int
}

func (coorChange *CoordIntChange) ToString() string {
	strng := fmt.Sprintf("CoordInt Change: at %d, %d change value to %d\n", coorChange.coord.X, coorChange.coord.Y, coorChange.changeTo)
	return strng
}

type CoordIntChangeQueue struct {
	CoList  CoordList
	Changes []int
}

func (Queue *CoordIntChangeQueue) PopNext() (CoordInts, int) {
	var tempCoords CoordInts
	tempCoords, Queue.CoList = Queue.CoList.PopFromFront() //,ight need to make this go to something else;
	tempInt := Queue.Changes[0]
	if len(Queue.Changes) > 1 {
		tempList := make([]int, 0)
		for i := 1; i < len(Queue.Changes); i++ {
			tempList = append(tempList, Queue.Changes[i])
		}
		Queue.Changes = tempList
	}
	return tempCoords, tempInt
}
func (Queue CoordIntChangeQueue) PopNext_Backup() (CoordIntChangeQueue, CoordInts, int) {
	tempQueue := Queue
	var tempCoords CoordInts
	tempCoords, tempQueue.CoList = Queue.CoList.PopFromFront() //,ight need to make this go to something else;
	tempInt := Queue.Changes[0]
	if len(Queue.Changes) > 1 {
		tempList := make([]int, 0)
		for i := 1; i < len(Queue.Changes); i++ {
			tempList = append(tempList, Queue.Changes[i])
		}
		tempQueue.Changes = tempList
	}
	return tempQueue, tempCoords, tempInt
}

func (Queue *CoordIntChangeQueue) PushToBack(coords CoordInts, change int) {
	Queue.Changes = append(Queue.Changes, change)
	Queue.CoList = Queue.CoList.PushToReturn(coords)
}
func (Queue CoordIntChangeQueue) PushToBack_Backup(coords CoordInts, change int) CoordIntChangeQueue {
	tempQueue := Queue
	tempQueue.Changes = append(tempQueue.Changes, change)
	tempQueue.CoList = tempQueue.CoList.PushToReturn(coords)
	return tempQueue
}

func (igm *IntegerGridManager) ManageChangesToGameboard() int {
	if len(igm.BoardChangesCoords) != len(igm.BoardChangeValues) {
		fmt.Printf("\n\nERROR\n\n\n")
		return -1
	}
	for {
		if len(igm.BoardChangesCoords) < 1 {
			break
		}
		_, tempCoord, tempNum := igm.GetNextChangeToGameboard()
		igm.Imat.DrawAGridTile(igm.BoardBuffer, tempCoord, igm.BoardMargin.X, igm.BoardMargin.Y, igm.Tile_Size.X, igm.Tile_Size.Y, igm.Margin.X, igm.Margin.Y, igm.Colors[tempNum], color.Black, 1.0, false, true)
		igm.Imat[tempCoord.Y][tempCoord.X] = tempNum
	}

	return 0
}
func (igm *IntegerGridManager) GetNextChangeToGameboard() (int, CoordInts, int) {
	if len(igm.BoardChangesCoords) != len(igm.BoardChangeValues) {
		fmt.Printf("\n\nERROR\n\n\n")
		return -1, CoordInts{X: 0, Y: 0}, -1
	}
	if len(igm.BoardChangeValues) != 0 {
		var tempCoord CoordInts
		tempInt := 0
		tempIntAr := make([]int, 0)
		tempCoord, igm.BoardChangesCoords = igm.BoardChangesCoords.PopFromFront()
		tempInt = igm.BoardChangeValues[0]
		if len(igm.BoardChangeValues) != 1 {
			for i := 1; i < len(igm.BoardChangeValues); i++ {
				tempIntAr = append(tempIntAr, igm.BoardChangeValues[i])
			}
		}
		igm.BoardChangeValues = tempIntAr
		return -1, tempCoord, tempInt
	}

	return -1, CoordInts{X: 0, Y: 0}, -1

}
