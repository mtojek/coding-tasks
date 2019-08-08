package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalancedBrackets_EmptyString(t *testing.T) {
	input := ""
	b := IsBalancedBrackets(input)
	assert.True(t, b)
}

func TestIsBalancedBrackets_AnyString(t *testing.T) {
	input := "gre"
	b := IsBalancedBrackets(input)
	assert.True(t, b)
}

func TestIsBalancedBrackets_BalancedBrackets(t *testing.T) {
	input := "((4+8)*3-4)*5"
	b := IsBalancedBrackets(input)
	assert.True(t, b)
}

func TestIsBalancedBrackets_NotBalancedBrackets_TooMuchOpened(t *testing.T) {
	input := "((4+8)*3-4)(*5"
	b := IsBalancedBrackets(input)
	assert.False(t, b)
}

func TestIsBalancedBrackets_NotBalancedBrackets_TooMuchClosed(t *testing.T) {
	input := "((4+8)*3-4)*)5"
	b := IsBalancedBrackets(input)
	assert.False(t, b)
}