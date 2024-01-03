package comm

import (
	"fmt"
	"runtime"
)

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

func InStringArray(value string, inData []string) bool {
	if len(inData) > 0 {
		for _, datum := range inData {
			if datum == value {
				return true
			}
		}
	}
	return false
}

func GetOs() {
	os := runtime.GOOS // 获取当前操作系统名称
	switch os {
	case "darwin":
		fmt.Println("当前操作系统为 macOS")
	case "linux":
		fmt.Println("当前操作系统为 Linux")
	case "windows":
		fmt.Println("当前操作系统为 Windows")
	default:
		fmt.Printf("未知操作系统 %s\n", os)
	}
}
