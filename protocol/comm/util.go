package comm

func If(condition bool, x, y any) any {
	if condition {
		return x
	}
	return y
}

// DecimalToBinary 二进制转换
func DecimalToBinary(num int) (binary []int) {
	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}

	//返回数据
	return
}
