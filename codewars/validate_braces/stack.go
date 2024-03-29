package validate_braces

type Stack []string

func (s *Stack) Add(val string) {
	*s = append(*s, val)
}

func (s *Stack) DeleteLast() {
	if l := len(*s); l > 0 {
		*s = (*s)[:l-1]
	}
}

func (s Stack) GetLast() (res *string) {
	if l := len(s); l > 0 {
		res = &s[len(s)-1]
	}
	return
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
