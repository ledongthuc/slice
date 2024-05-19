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
