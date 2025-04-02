package mypkgs

import (
	"fmt"
)

func (imat *IntMatrix) SaveIntMatrixToFile(filename string) error {

	fmt.Printf("Saving Matrix\n")
	return nil
}

func (imat *IntMatrix) LoadIntMatrixFromFile(filename string) (IntMatrix, error) {
	temp := make(IntMatrix, 0)
	fmt.Printf("Saving Matrix\n")
	return temp, nil
}
