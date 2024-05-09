package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/manish-neemnarayan/assignment/types"
)

func AddBookService(params *types.AddBook) (string, error) {
	data := [][]string{
		{"" + params.BookName, params.Author, params.PublicationYear}, // Single row inside
	}

	if err := writeCSV("./resource/regularUser.csv", data); err != nil {
		return "", err
	}

	return "success", nil
}

func DeleteBookService(bookName string) error {
	records, err := readCSV("./resource/regularUser.csv")
	if err != nil {
		return fmt.Errorf("error while reading file")
	}

	//filter record based on book name
	filteredRecords := [][]string{}
	for _, row := range records {
		if !strings.EqualFold(row[0], bookName) {
			filteredRecords = append(filteredRecords, row)
		}
	}

	//overwrite
	csvFile, err := os.Create("./resource/regularUser.csv")
	if err != nil {
		return err
	}

	defer csvFile.Close()

	writer := csv.NewWriter((csvFile))
	if err := writer.WriteAll(filteredRecords); err != nil {
		return err
	}

	writer.Flush()
	return nil
}

// // utility function
// // check the utility.service.go
