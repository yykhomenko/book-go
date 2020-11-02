package eval

import (
	"fmt"
	"testing"
)

func TestCallString(t *testing.T) {
	const input = "sin(x)+-25+sin(2)*2/10-3.14/sin(1/2)"
	expr, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(expr)

}
