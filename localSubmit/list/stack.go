package main

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, true
}

func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

// func (s *Stack) Peek() (int, bool) {
// 	if len(s.data) == 0 {
// 		return 0, false
// 	}
// 	return s.data[len(s.data)-1], true
// }

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}