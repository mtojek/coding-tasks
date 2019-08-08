package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsChainable_EmptyNodes(t *testing.T) {
	var nodes []int
	c := IsChainable(nodes, 0, 9)
	assert.False(t, c)
}

func TestIsChainable_Sample(t *testing.T) {
	nodes := []int{8363, 6388, 8183, 5364, 8353, 8365, 9380}
	c := IsChainable(nodes, 8183, 6388)
	assert.True(t, c)
}

func TestIsChainable_FewHops(t *testing.T) {
	nodes := []int{1122, 2233, 3344, 4455, 5566}
	c := IsChainable(nodes, 1122, 5566)
	assert.True(t, c)
}

func TestIsChainable_FewHops_PotentialLoop(t *testing.T) {
	nodes := []int{1122, 2233, 3322, 3344, 4455, 5566}
	c := IsChainable(nodes, 1122, 5566)
	assert.True(t, c)
}

func TestIsChainable_AnotherCase(t *testing.T) {
	nodes := []int{1122, 2233, 2243, 2245, 4322, 4355, 6655, 5566, 6677}
	c := IsChainable(nodes, 1122, 6677)
	assert.True(t, c)
}