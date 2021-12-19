package helpers

type IntStack struct {
	data []int
}

func (s *IntStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *IntStack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	idx := len(s.data) - 1
	res := s.data[idx]
	s.data = s.data[:idx]
	return res, true
}

func (s *IntStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}

type RuneStack struct {
	data []rune
}

func (s *RuneStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *RuneStack) Push(i rune) {
	s.data = append(s.data, i)
}

func (s *RuneStack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	idx := len(s.data) - 1
	res := s.data[idx]
	s.data = s.data[:idx]
	return res, true
}

func (s *RuneStack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}
