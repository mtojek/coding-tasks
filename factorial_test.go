package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorial_Zero(t *testing.T) {
	k := factorial(0)
	assert.Equal(t, 1, k)
}

func TestFactorial_One(t *testing.T) {
	k := factorial(1)
	assert.Equal(t, 1, k)
}

func TestFactorial_Two(t *testing.T) {
	k := factorial(2)
	assert.Equal(t, 2, k)
}

func TestFactorial_Three(t *testing.T) {
	k := factorial(3)
	assert.Equal(t, 6, k)
}

func TestFactorial_Four(t *testing.T) {
	k := factorial(4)
	assert.Equal(t, 24, k)
}