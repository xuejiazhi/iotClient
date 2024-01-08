package comm

import (
	"bytes"
	"encoding/json"
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

func StructToMap(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}
