package other

import "fmt"

func MinSwapArray(arr []int) int {
	// Initialise count variable
	count := 0
	i := 0
	fmt.Printf("arr %v, index saat ini %v \n", arr, i)
	for i < len(arr) {
		fmt.Printf("CEK PERTAMA 1 array: %v, array i: %v, i+1 = %v, result if: %v \n", arr, arr[i], i+1, arr[i] != i+1)
		// If current element is not at the right position
		// Jika elemen saat ini tidak pada posisi yang tepat
		if arr[i] != i+1 {

			fmt.Printf("FOR KEDUA 2 array: %v, array i: %v, i+1 = %v, result if: %v \n", arr, arr[i], i+1, arr[i] != i+1)
			for arr[i] != i+1 {
				temp := 0
				fmt.Printf(" arr: %v  , index: %v \n", arr, i)

				// Swap current element with correct position of that element
				// Tukar elemen saat ini dengan posisi yang benar dari elemen itu
				temp = arr[arr[i]-1]

				arr[arr[i]-1] = arr[i]
				arr[i] = temp
				fmt.Printf(" arr: %v  , index: %v \n", arr, i)
				fmt.Println()
				count++
			}
			fmt.Printf("FOR KEDUA 2 SELESAI array: %v, array i: %v, i+1 = %v, result if: %v \n", arr, arr[i], i+1, arr[i] != i+1)
		}

		// Increment for next index when current element is at correct position
		// Inkremen untuk indeks berikutnya saat elemen saat ini berada di posisi yang benar
		i++
	}
	fmt.Printf("FINAL arr %v, index saat ini %v \n", arr, i)
	return count
}
