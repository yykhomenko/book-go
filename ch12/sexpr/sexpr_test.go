package sexpr

import (
	"fmt"
	"os"
	"testing"
)

func TestMarshal(t *testing.T) {
	b, err := Marshal(os.Stdin)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}
