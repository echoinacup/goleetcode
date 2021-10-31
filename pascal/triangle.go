package pascal

func getRow(rowIndex int) []int {
	return createRow([]int{1}, rowIndex)
}

func createRow(previousRow []int, rowIndex int) []int {
	if rowIndex == 0 {
		return previousRow
	}
	currentRow := []int{1}
	for i := 0; i < len(previousRow)-1; i++ {
		currentRow = append(currentRow, previousRow[i]+previousRow[i+1])
	}
	currentRow = append(currentRow, 1)

	return createRow(currentRow, rowIndex-1)
}
