package main

import (
	"errors"
	"fmt"
)

// errors.Is() 函数的应用：用来判断是否匹配某个具体的错误值或被它包裹的错误，适合做等值匹配。
// errors.Is(err, targetErr)
var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
	err1 := fmt.Errorf("wrap err1: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err2: %w", err1)

	if errors.Is(err2, ErrSentinel) {
		println("err is ErrSentinel")
		return
	}

	println("err is not ErrSentinel")
}
