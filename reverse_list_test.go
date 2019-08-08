package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse_NoElements(t *testing.T) {
	var elems ListOfElements
	elems.Reverse()
	assert.Nil(t, elems.Tail)
}

func TestReverse_SingleElement(t *testing.T) {
	var elems ListOfElements
	elems.Push(1)
	elems.Reverse()
	assert.Equal(t, elems.Tail.Value, 1)
}

func TestReverse_TwoElements(t *testing.T) {
	var elems ListOfElements
	elems.Push(1)
	elems.Push(2)
	elems.Reverse()
	assert.Equal(t, elems.Tail.Value, 2)
	assert.Equal(t, elems.Tail.Next.Value, 1)
}

func TestReverse_ThreeElements(t *testing.T) {
	var elems ListOfElements
	elems.Push(1)
	elems.Push(2)
	elems.Push(3)
	elems.Reverse()
	assert.Equal(t, elems.Tail.Value, 3)
	assert.Equal(t, elems.Tail.Next.Value, 2)
	assert.Equal(t, elems.Tail.Next.Next.Value, 1)
}

func TestReverse_FourElements(t *testing.T) {
	var elems ListOfElements
	elems.Push(1)
	elems.Push(2)
	elems.Push(3)
	elems.Push(4)
	elems.Reverse()
	assert.Equal(t, elems.Tail.Value, 4)
	assert.Equal(t, elems.Tail.Next.Value, 3)
	assert.Equal(t, elems.Tail.Next.Next.Value, 2)
	assert.Equal(t, elems.Tail.Next.Next.Next.Value, 1)
}