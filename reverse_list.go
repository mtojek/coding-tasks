package coding_tasks

type ListOfElements struct {
	Tail *Element
}

type Element struct {
	Value int

	Next *Element
	Prev *Element
}

func (lof *ListOfElements) Push(value int) {
	if lof.Tail == nil {
		lof.Tail = &Element{
			Value: value,
		}
		return
	}

	lof.Tail.Next = &Element{
		Value: value,

		Prev: lof.Tail,
	}
	lof.Tail = lof.Tail.Next
}

func (lof *ListOfElements) Peek() int {
	if lof.Tail == nil {
		return -1
	}
	return lof.Tail.Value
}

func (lof *ListOfElements) Pop() int {
	k := lof.Tail.Value
	lof.Tail = lof.Tail.Prev
	if lof.Tail != nil {
		lof.Tail.Next = nil
	}
	return k
}

func (lof *ListOfElements) Reverse() {
	if lof.Tail == nil {
		return
	}

	ptr := lof.Tail
	for ptr.Prev != nil {
		k := ptr.Next
		ptr.Next = ptr.Prev
		ptr.Prev = k
		ptr = ptr.Next
	}
}
