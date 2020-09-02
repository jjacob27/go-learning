package main

type Stack []string

func (s *Stack) Push(d string) {
	*s = append(*s, d)
}

func (s *Stack) Pop() (bool, string) {
	if s.IsEmpty() {
		return false, ""
	}

	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return true, e
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}

	return (*s)[len(*s)-1]
}
