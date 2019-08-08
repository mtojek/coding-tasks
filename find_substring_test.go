package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubstring_EmptyStack(t *testing.T) {
	stack := []byte("")
	needle := []byte("needle")
	i := FindSubstring(stack, needle)
	assert.Equal(t, -1, i)
}

func TestFindSubstring_EmptyNeedle(t *testing.T) {
	stack := []byte("stack")
	needle := []byte("")
	i := FindSubstring(stack, needle)
	assert.Equal(t, -1, i)
}

func TestFindSubstring_NeedleFound(t *testing.T) {
	stack := []byte("pejfioneedlepejw")
	needle := []byte("needle")
	i := FindSubstring(stack, needle)
	assert.Equal(t, 6, i)
}

func TestFindSubstring_NeedleNotFound(t *testing.T) {
	stack := []byte("pejfioneedlpejw")
	needle := []byte("needle")
	i := FindSubstring(stack, needle)
	assert.Equal(t, -1, i)
}


func TestFindSubstring_NeedleFoundFurther(t *testing.T) {
	stack := []byte("pejfioneedlpejwneedlei")
	needle := []byte("needle")
	i := FindSubstring(stack, needle)
	assert.Equal(t, 15, i)
}

func TestFindSubstring_NeedleFound_Tricky(t *testing.T) {
	stack := []byte("bcbcbcd")
	needle := []byte("bcbcd")
	i := FindSubstring(stack, needle)
	assert.Equal(t, 2, i)
}