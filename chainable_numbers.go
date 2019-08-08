package coding_tasks

/**
 * Given a set of four digit numbers, and two numbers A and B which are in the set,
 * indicate if A and B are chainable. Any numbers X and Y are chainable
 * if the last two digits of X are the first two digits of Y, with any number of chainable numbers in between.
 *
 * For example, given the set {8363, 6388, 8183, 5364, 8353, 8365, 9380}, A=8183, B=6388, yes, A and B are chainable
 * (as [8183, 8363, 6388]).
 */

type Node struct {
	Value int

	OtherNodes map[int]bool
}

func IsChainable(nodes []int, a, b int) bool {
	visited := []int{a}
	return isChainableV(nodes, visited, a, b)
}

func isChainableV(nodes []int, visited []int, a, b int) bool {
	if len(nodes) == 0 {
		return false
	}

	if areTwoNumbersChainable(a, b) {
		return true
	}

	for _, node := range nodes {
		if node != a && areTwoNumbersChainable(a, node) {
			if !isVisitedNode(visited, node) {
				visited = append(visited, node)
				isc := isChainableV(nodes, visited, node, b)
				if isc {
					return true
				}
			}
		}
	}

	return false
}

func areTwoNumbersChainable(a, b int) bool {
	return (a % 100) == (b / 100)
}

func isVisitedNode(visited []int, n int) bool {
	for _, v := range visited {
		if v == n {
			return true
		}
	}
	return false
}