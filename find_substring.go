package coding_tasks

func FindSubstring(stack, needle []byte) int {
	if len(stack) == 0 || len(needle) == 0 {
		return -1
	}

	j := 0
	for i := 0; i < len(stack); i++ {
		for ii := i; ii < len(stack); ii++ {
			if stack[ii] == needle[j] {
				j++
				if j == len(needle) {
					return ii - len(needle) + 1
				}
			} else {
				j = 0
				break
			}
		}
	}
	return -1
}