package calculator

type Basic struct {
	No1    int
	No2    int
	Result int
}

// Return the sum of first(No1) and second item (No2)
func (b *Basic) GetSum() int {
	return b.No1 + b.No2
}

func (b *Basic) GetDiff() int {
	return b.No1 - b.No2
}

// Return the mul
func (b *Basic) GetMul() (int, int, int) {
	return b.No1 * b.No2, b.No1, b.No2
}
