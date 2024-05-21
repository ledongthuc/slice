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
