package comm

func If(condition bool, x, y any) any {
	if condition {
		return x
	}
	return y
}
