package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func StrToInt(s string) (int, error) {
	// 1. 验证长度是否为10位
	if len(s) != 10 {
		return 0, errors.New("输入必须是10位字符串")
	}

	// 2. 验证是否为纯数字（使用正则）
	if !regexp.MustCompile(`^\d+$`).MatchString(s) {
		return 0, errors.New("输入必须是纯数字字符串")
	}

	// 3. 去除高位0（转换为int后会自动忽略前导0）
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("转换失败: %v", err)
	}

	return num, nil
}

// 将int转换为10位字符串，不足补0，超过10位返回错误
func IntTo10DigitStr(num uint) (string, error) {
	// 检查数值是否超过10位最大范围（9999999999）
	if num > 9999999999 {
		return "", fmt.Errorf("数值超出10位范围（0~9999999999）")
	}

	// 使用fmt格式化：%010d表示不足10位时高位补0，共10位
	return fmt.Sprintf("%010d", num), nil
}
