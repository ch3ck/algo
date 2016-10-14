package insertsort

func Swap(arraysrt []int, i, j int) {
	tmp := arraysrt[j]
	arraysrt[j] = arraysrt[i]
	arraysrt[i] = tmp
}

func InsertSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			Swap(array, j, j-1)
		}
	}
	return array
}
