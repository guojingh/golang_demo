package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add：", result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(3, 1)
	fmt.Println("add：", result)
}
