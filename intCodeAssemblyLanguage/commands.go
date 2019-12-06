package main

const (
	VALUE    uint8 = 0
	POINTER  uint8 = 1
	VARIABLE uint8 = 2
)

type Word string

func convStrSlicetoWordSlice(in []string) []Word {
	out := make([]Word, 0)
	for _, v := range in {
		out = append(out, Word(v))
	}
	return out
}

func (w Word) getType() uint8 {
	switch w[0] {
	case '*':
		return POINTER
	case '_':
		return VARIABLE
	default:
		return VALUE
	}
}

func add(args []Word) []int {
	return []int{99}
}
func multi(args []Word) []int {
	return []int{99}
}
func input(args []Word) []int {
	return []int{99}
}
func output(args []Word) []int {
	return []int{99}
}

func end(args []Word) []int {
	return []int{99}
}
