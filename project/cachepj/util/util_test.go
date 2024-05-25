package util

import (
	"fmt"
	"testing"
)

// 测试获取内存容量函数
func TestParseSize(t *testing.T) {

	size, s := ParseSize("200GB")

	t.Logf(fmt.Sprintf("%d,%s", size, s))
}
