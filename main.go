package main

import (
	"fmt"
	"strings"
)

func isMutant(dna []string) bool {
	matrix := make([][]string, len(dna))
	for index, v := range dna {
		row := strings.Split(v, "")
		matrix[index] = row
	}
	fmt.Println(matrix)
	var i, j int
	isMutant := false
	for i = 0; i < len(matrix); i++ {
		for j = 0; j < len(matrix[0])-3; j++ {
			if isMutant {
				break
			}
			if i < len(matrix)-3 {
				if matrix[i][j] == matrix[i+1][j+1] && matrix[i][j] == matrix[i+2][j+2] && matrix[i][j] == matrix[i+3][j+3] {
					fmt.Printf("Value %v %v %v %v\n", matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3], matrix[i][j])
					fmt.Println("Is Mutant")
					isMutant = true
					break
				}

			}

			if matrix[i][j] == matrix[i][j+1] && matrix[i][j] == matrix[i][j+2] && matrix[i][j] == matrix[i][j+3] {
				fmt.Printf("Value %v %v %v %v\n", matrix[i][j+1], matrix[i][j+2], matrix[i][j+3], matrix[i][j])
				fmt.Println("Is Mutant")
				isMutant = true
				break
			}

			if matrix[j][i] == matrix[j+1][i] && matrix[j][i] == matrix[j+2][i] && matrix[j][i] == matrix[j+3][i] {
				fmt.Printf("Value %v %v %v %v\n", matrix[j+1][i], matrix[j+2][i], matrix[j+3][i], matrix[j][i])
				fmt.Println("Is Mutant")
				isMutant = true
				break
			}

		}

	}

	return isMutant
}

func main() {
	dna := []string{"TTGAGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	fmt.Println(isMutant(dna))

}
