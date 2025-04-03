package mypkgs

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func (imat *IntMatrix) SaveIntMatrixToFile(filename string) error {
	//fmt.Printf("Saving Matrix\n")
	file, err := os.Create(fmt.Sprintf("%s.gob", filename))
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(imat)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (imat *IntMatrix) LoadIntMatrixFromFile(filename string) (IntMatrix, error) {
	temp := make(IntMatrix, 0)
	//fmt.Printf("Loading Matrix\n")
	file, err := os.Open(fmt.Sprintf("%s.gob", filename))
	if err != nil {
		return temp, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&temp)
	if err != nil {
		return nil, err
	}

	return temp, nil
}
