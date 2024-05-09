package service

import (
	"fmt"

	"github.com/manish-neemnarayan/assignment/types"
)

func HomeService(role string) (*types.HomeResponse, error) {
	switch role {
	case "user":
		//get user csv data
		userCSV, err := readCSV("./resource/regularUser.csv")
		if err != nil {
			return &types.HomeResponse{}, err
		}

		bookNames := make([]string, 0)
		for i, row := range userCSV {
			if i == 0 {
				continue
			}
			bookNames = append(bookNames, row[0])
		}

		return &types.HomeResponse{
			BookName: bookNames,
		}, nil

	case "admin":
		//get the data from admin csv
		adminCSV, err := readCSV("./resource/adminUser.csv")
		if err != nil {
			fmt.Printf("err in user csv %v\n", err)
			return &types.HomeResponse{}, err
		}

		//get the data from user csv
		userCSV, err := readCSV("./resource/regularUser.csv")
		if err != nil {

			fmt.Printf("err in admin csv %v\n", err)
			return &types.HomeResponse{}, err
		}

		bookNames := make([]string, 0)

		//combine admin and user csv data
		for i, row := range adminCSV {
			if i == 0 {
				continue
			}
			bookNames = append(bookNames, row[0])
		}

		for i, row := range userCSV {
			if i == 0 {
				continue
			}
			bookNames = append(bookNames, row[0])
		}

		return &types.HomeResponse{
			BookName: bookNames,
		}, nil

	default:
		fmt.Println("invalid role")
		return &types.HomeResponse{}, fmt.Errorf("unauthorized: invalid access")
	}

}

// // utility function
// // check the utility.service.go
