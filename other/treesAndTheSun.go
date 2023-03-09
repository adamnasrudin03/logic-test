package other

func TreesAndTheSun(trees []int, sun int) (count int) {
	maxTree := trees[0]
	// Initialize result (first tree always sees sunlight)
	count = 1

	for i := 1; i < sun; i++ {
		if trees[i] > maxTree || trees[i] == maxTree {
			maxTree = trees[i]
			count++
		}

	}
	return count
}
