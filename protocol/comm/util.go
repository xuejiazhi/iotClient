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

func InIntArray(value int, inData []int) bool {
	if len(inData) > 0 {
		for _, datum := range inData {
			if datum == value {
				return true
			}
		}
	}
	return false
}
