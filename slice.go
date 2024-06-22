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

func (s *Slice[M]) DeleteClean(index int) {
	length := len(s.internal)

	if index < 0 || index >= length {
		return
	}

	if index >= 0 && index < length-1 {
		copy(s.internal[index:], s.internal[index+1:])
	}

	var defaultValue M
	s.internal[length-1] = defaultValue
	s.internal = s.internal[:length-1]
}

func (s *Slice[M]) DeleteUnordered(index int) {
	length := len(s.internal)

	if index < 0 || index >= length {
		return
	}

	s.internal[index] = s.internal[length-1]
	s.Delete(length - 1)
}

func (s *Slice[M]) DeleteUnorderedClean(index int) {
	length := len(s.internal)

	if index < 0 || index >= length {
		return
	}

	s.internal[index] = s.internal[length-1]
	s.DeleteClean(length - 1)
}

func (s *Slice[M]) CutClean(start int, end int) {
	length := len(s.internal)
	start_index := max(0, start)
	end_index := min(length-1, max(start_index, end))
	copy(s.internal[start_index:], s.internal[end_index:])

	var default_value M
	has_deleted := false
	new_length := length - end_index + start_index

	for i := new_length; i < length; i++ {
		has_deleted = true
		s.internal[i] = default_value
	}

	if has_deleted {
		s.internal = s.internal[:new_length]
	}
}

func (s *Slice[M]) Cut(start int, end int) {
	start_index := max(0, start)
	end_index := min(len(s.internal)-1, max(start_index, end))

	s.internal = append(s.internal[:start_index], s.internal[end_index:]...)
}

func (s *Slice[M]) Append(list []M) {
	s.internal = append(s.internal, list...)
}

func (s *Slice[M]) Pop() M {
	x := s.internal[len(s.internal)-1]
	s.DeleteClean(s.GetLength() - 1)

	return x
}

func (s *Slice[M]) Shift() M {
	x := s.internal[0]
	s.DeleteClean(0)

	return x
}

func (s *Slice[M]) Push(element M) {
	s.internal = append(s.internal, element)
}

func (s *Slice[M]) Expand(from_index int, element_number int) {
	expanded_s := append(make([]M, element_number), s.internal[from_index:]...)
	s.internal = append(s.internal[:from_index], expanded_s...)
}

func (s *Slice[M]) Insert(val M, at int) {
	last_part := append([]M{val}, s.internal[at:]...)
	s.internal = append(s.internal[:at], last_part...)
}

func (s *Slice[M]) Extend(element_number int) {
	s.internal = append(s.internal, make([]M, element_number)...)
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
