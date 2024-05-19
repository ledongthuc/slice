package slice

type Slice[M any] struct {
	internal []M
}

func (s Slice[M]) Copy() Slice[M] {
	newS := make([]M, len(s.internal))
	copy(newS, s.internal)
	return Slice[M]{newS}
}
