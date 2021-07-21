package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

func WriteToCsv(array *[]Issue, fileName string) error {
	clientsFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer clientsFile.Close()
	err = gocsv.MarshalFile(array, clientsFile) // Use this to save the CSV back to the file
	return err
}
