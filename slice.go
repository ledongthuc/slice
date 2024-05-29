package slice

type Slice[M any] struct {
	internal []M
}

func (s Slice[M]) GetLength() int {
	return len(s.internal)
}

func (s Slice[M]) GetRaw() []M {
	return s.internal
}

func (s Slice[M]) Copy() Slice[M] {
	newS := make([]M, len(s.internal))
	copy(newS, s.internal)
	return Slice[M]{newS}
}

func (s *Slice[M]) Delete(index int) {
	length := len(s.internal)

	if index == 0 {
		s.internal = s.internal[1:]
		return
	}

	if index == length-1 {
		s.internal = s.internal[:length-1]
		return
	}

	if index > 0 && index < length {
		s.internal = append(s.internal[:index], s.internal[index+1:]...)
		return
	}
}

func (s *Slice[M]) SoftDelete(index int) {
	length := len(s.internal)

	if index < 0 || index >= length {
		return
	}

	if index >= 0 && index < length-1 {
		copy(s.internal[index:], s.internal[index+1:])
	}

	var defaultValue M
	s.internal[length-1] = defaultValue
}

func (s *Slice[M]) Cut(start int, end int) {
	start_index := max(0, start)
	end_index := min(len(s.internal)-1, max(start_index, end))

	s.internal = append(s.internal[:start_index], s.internal[end_index:]...)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
