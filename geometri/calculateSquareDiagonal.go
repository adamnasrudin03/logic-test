package geometri

import (
	"fmt"
	"math/rand"
)

func CalculateSquareDiagonal(total int) {
	var square = map[int][]int{}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// create total array and contents
	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			randomIndex := rand.Intn(len(numbers))
			pick := numbers[randomIndex]
			square[i] = append(square[i], pick)
		}
	}

	for z := 0; z < total; z++ {
		fmt.Printf("Array #%v : %v \n", z+1, square[z])
	}

	diagonalXText := ""
	diagonalX := 0
	diagonalYText := ""
	diagonalY := 0
	y := total - 1

	// calculating the diagonals of the array
	for x := 0; x < total; x++ {
		diagonalX = diagonalX + square[x][x]
		if x == 0 {
			diagonalXText = "("
		}

		diagonalXText += fmt.Sprintf("%d", square[x][x])

		if x != len(square)-1 || x == 0 {
			diagonalXText += " + "
		}

		if x == len(square)-1 {
			diagonalXText += ")"
		}

		for y >= 0 {
			diagonalY = diagonalY + square[x][y]
			if y == len(square)-1 {
				diagonalYText += "("
			}

			diagonalYText += fmt.Sprintf("%d", square[x][y])

			if y == len(square)-1 || y != 0 {
				diagonalYText += " + "
			}
			if y == 0 {
				diagonalYText += ")"
			}

			y--
			break
		}
	}

	fmt.Println("==============================")
	fmt.Printf("Hasil : %v - %v = %v", diagonalXText, diagonalYText, (diagonalX - diagonalY))
}
