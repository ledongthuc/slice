package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	empty := Slice[int]{make([]int, 0)}
	copiedEmpty := empty.Copy()
	require.Equalf(t, empty, copiedEmpty, "empty")

	padding := Slice[int]{make([]int, 10)}
	copiedPadding := padding.Copy()
	require.Equalf(t, padding, copiedPadding, "padding")

	list := Slice[int]{[]int{1, 2, 3, 4, 5}}
	copiedList := list.Copy()
	require.Equalf(t, list, copiedList, "list")
}

func TestDelete(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	list.Delete(-1)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "delete out of range from left")

	list.Delete(5)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "delete out of range from right")

	list.Delete(0)
	require.Equalf(t, list.internal, []int{2, 3, 4, 5}, "delete at index 0")

	list.Delete(len(list.internal) - 1)
	require.Equalf(t, list.internal, []int{2, 3, 4}, "delete at last element index")

	list.Delete(1)
	require.Equalf(t, list.internal, []int{2, 4}, "delete at index")
}

func TestDeleteUnordered(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5, 6}}

	list.DeleteUnordered(-1)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5, 6}, "delete out of range from left")

	list.DeleteUnordered(6)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5, 6}, "delete out of range from right")

	last_val := &list.internal[len(list.internal)-1]
	require.Equalf(t, *last_val, 6, "last element before deleting")
	list.DeleteUnordered(0)
	require.Equalf(t, list.internal, []int{6, 2, 3, 4, 5}, "delete at index 0")
	require.Equalf(t, *last_val, 6, "last element after deleting")

	last_val = &list.internal[len(list.internal)-1]
	require.Equalf(t, *last_val, 5, "last element before deleting")
	list.DeleteUnordered(len(list.internal) - 1)
	require.Equalf(t, list.internal, []int{6, 2, 3, 4}, "delete at last element index")
	require.Equalf(t, *last_val, 5, "last element after deleting")

	last_val = &list.internal[len(list.internal)-1]
	require.Equalf(t, *last_val, 4, "last element before deleting")
	list.DeleteUnordered(1)
	require.Equalf(t, list.internal, []int{6, 4, 3}, "delete at index")
	require.Equalf(t, *last_val, 4, "last element after deleting")
}

func TestGetLength(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}
	require.Equalf(t, list.GetLength(), 5, "get length of slice")
}

func TestGetRaw(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}
	require.Equalf(t, list.GetRaw(), []int{1, 2, 3, 4, 5}, "get raw array of slice")
}

func TestDeleteClean(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	list.DeleteClean(-1)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "delete out of range from left")

	list.DeleteClean(5)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "delete out of range from right")

	val := &list.internal[len(list.internal)-1]
	require.Equalf(t, *val, 5, "last element before deleting")
	list.DeleteClean(len(list.internal) - 1)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4}, "delete at last element index")
	require.Equalf(t, *val, 0, "last element after deleting")

	val = &list.internal[len(list.internal)-1]
	require.Equalf(t, *val, 4, "last element before deleting")
	list.DeleteClean(0)
	require.Equalf(t, list.internal, []int{2, 3, 4}, "delete at index 0")
	require.Equalf(t, *val, 0, "last element after deleting")

	val = &list.internal[len(list.internal)-1]
	require.Equalf(t, *val, 4, "last element before deleting")
	list.DeleteClean(1)
	require.Equalf(t, list.internal, []int{2, 4}, "delete at index")
	require.Equalf(t, *val, 0, "last element after deleting")
}

func TestDeleteUnorderedClean(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5, 6}}

	list.DeleteUnorderedClean(-1)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5, 6}, "delete out of range from left")

	list.DeleteUnorderedClean(6)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5, 6}, "delete out of range from right")

	last_val := &list.internal[len(list.internal)-1]
	require.Equalf(t, *last_val, 6, "last element before deleting")
	list.DeleteUnorderedClean(0)
	require.Equalf(t, list.internal, []int{6, 2, 3, 4, 5}, "delete at index 0")
	require.Equalf(t, *last_val, 0, "last element after deleting")

	last_val = &list.internal[len(list.internal)-1]
	require.Equalf(t, *last_val, 5, "last element before deleting")
	list.DeleteUnorderedClean(len(list.internal) - 1)
	require.Equalf(t, list.internal, []int{6, 2, 3, 4}, "delete at last element index")
	require.Equalf(t, *last_val, 0, "last element after deleting")

	last_val = &list.internal[len(list.internal)-1]
	require.Equalf(t, *last_val, 4, "last element before deleting")
	list.DeleteUnorderedClean(1)
	require.Equalf(t, list.internal, []int{6, 4, 3}, "delete at index")
	require.Equalf(t, *last_val, 0, "last element after deleting")
}

func TestCut(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	list.Cut(-1, 0)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "cut out of range from left")

	list.Cut(4, 6)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "cut out of range from right")

	list.Cut(1, 4)
	require.Equalf(t, list.internal, []int{1, 5}, "cut from valid range")
}

func TestCutClean(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	list.Cut(-1, 0)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "cut out of range from left")

	list.Cut(4, 6)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "cut out of range from right")

	list.Cut(1, 4)
	require.Equalf(t, list.internal, []int{1, 5}, "cut from valid range")
}

func TestAppend(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}
	dest_list := Slice[int]{[]int{6, 7, 8}}

	list.Append(dest_list)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5, 6, 7, 8}, "append dest_list")
}

func TestPop(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	val := list.Pop()
	require.Equalf(t, list.internal, []int{1, 2, 3, 4}, "pop list")
	require.Equalf(t, val, 5, "pop value")
}

func TestShift(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	val := list.Shift()
	require.Equalf(t, list.internal, []int{2, 3, 4, 5}, "shift list")
	require.Equalf(t, val, 1, "shift value")
}

func TestPush(t *testing.T) {
	list := Slice[int]{[]int{1}}

	list.Push(2)
	require.Equalf(t, list.internal, []int{1, 2}, "push 2")

	list.Push(3)
	require.Equalf(t, list.internal, []int{1, 2, 3}, "push 3")
}

func TestExpand(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	list.Expand(2, 4)
	require.Equalf(t, list.internal, []int{1, 2, 0, 0, 0, 0, 3, 4, 5}, "expand from 2 with 4 elements")
}
