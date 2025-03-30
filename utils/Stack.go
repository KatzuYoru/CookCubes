package utils

type Stack []string

func (st *Stack) Push(val string) {
	*st = append(*st, val)
}

func (st *Stack) Pop() (string, bool) {
	if len(*st) == 0 {
		return "", false
	}

	index := len(*st) - 1
	element := (*st)[index]
	*st = (*st)[:index]
	return element, true
}

func (st *Stack) Peek() (string, bool) {
	if len(*st) == 0 {
		return "", false
	}

	return (*st)[len(*st)-1], true
}

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Size() int {
	return len(*st)
}
