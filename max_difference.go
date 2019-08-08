package coding_tasks

func maxDifference(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	min := arr[0]
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max - min

}