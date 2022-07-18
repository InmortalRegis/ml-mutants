package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isMutant(dna []string) (bool, error) {
	base := []string{"A", "T", "C", "G"}
	matrix := make([][]string, len(dna))
	for index, v := range dna {
		row := strings.Split(v, "")
		for _, e := range row {
			if !contains(base, e) {
				return false, errors.New("Invalid DNA")
			}
		}
		matrix[index] = row
	}
	var i, j int
	isMutant := false
	for i = 0; i < len(matrix); i++ {
		for j = 0; j < len(matrix[0])-3; j++ {
			if isMutant {
				break
			}
			if i < len(matrix)-3 {
				// Diagonal
				if matrix[i][j] == matrix[i+1][j+1] && matrix[i][j] == matrix[i+2][j+2] && matrix[i][j] == matrix[i+3][j+3] {
					fmt.Println("Diagonal")
					fmt.Printf("m[%d][%d]: %s %s %s %s \n", i, j, matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3])
					isMutant = true
					break
				}

			}
			// Horizontal
			if matrix[i][j] == matrix[i][j+1] && matrix[i][j] == matrix[i][j+2] && matrix[i][j] == matrix[i][j+3] {
				fmt.Println("Horizontal")
				fmt.Printf("m[%d][%d]: %s %s %s %s \n", i, j, matrix[i][j], matrix[i][j+1], matrix[i][j+2], matrix[i][j+3])
				isMutant = true
				break
			}
			// Vertical
			if matrix[j][i] == matrix[j+1][i] && matrix[j][i] == matrix[j+2][i] && matrix[j][i] == matrix[j+3][i] {
				fmt.Println("Vertical")
				fmt.Printf("m[%d][%d]: %s %s %s %s \n", j, i, matrix[j][i], matrix[j+1][i], matrix[j+2][i], matrix[j+3][i])
				isMutant = true
				break
			}

		}

	}

	return isMutant, nil
}

func main() {
	start := time.Now()
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	isMutant, err := isMutant(dna)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Is Mutant: ", isMutant)
	fmt.Println(time.Since(start))
}
