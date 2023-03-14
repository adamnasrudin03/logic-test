package other

import "fmt"

func MinSwapArray(arr []int) int {
	// Initialise count variable
	count := 0
	// // another solution by using asc loop
	// temp := 0
	// min := 0
	// for i := 0; i < len(arr); i++ {
	// 	min = i
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[j] < arr[min] {

	// 			// changing the index to show the min value
	// 			min = j
	// 			count++
	// 			break
	// 		}
	// 	}
	// 	temp = arr[i]
	// 	arr[i] = arr[min]
	// 	arr[min] = temp
	// }

	//  solution
	i := 0
	for i < len(arr) {
		// If current element is not at the right position
		// Jika elemen saat ini tidak pada posisi yang tepat
		if arr[i] != i+1 {

			for arr[i] != i+1 {
				temp := 0
				fmt.Printf(" arr: %v  , index: %v \n", arr, i)

				// Swap current element with correct position of that element
				// Tukar elemen saat ini dengan posisi yang benar dari elemen itu
				temp = arr[arr[i]-1]

				arr[arr[i]-1] = arr[i]
				arr[i] = temp
				count++
			}
		}

		// Increment for next index when current element is at correct position
		// Inkremen untuk indeks berikutnya saat elemen saat ini berada di posisi yang benar
		i++
	}
	return count
}
