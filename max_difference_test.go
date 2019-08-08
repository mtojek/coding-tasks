package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDifference_OneElement(t *testing.T) {
	arr := []int{7}

	k := maxDifference(arr)

	assert.Equal(t, 0, k)
}

func TestMaxDifference_TwoElements(t *testing.T) {
	arr := []int{7, -3}

	k := maxDifference(arr)

	assert.Equal(t, 10, k)
}

func TestMaxDifference_SixElements(t *testing.T) {
	arr := []int{7, -3, 23, -4, 42, -37}

	k := maxDifference(arr)

	assert.Equal(t, 79, k)
}