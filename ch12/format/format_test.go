package format

import (
	"fmt"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	tests := []struct {
		args interface{}
		want string
	}{
		{int64(1), "1"},
		{time.Nanosecond, "1"},
	}
	for _, tt := range tests {
		if got := Any(tt.args); got != tt.want {
			t.Errorf("Any() = %v, want %v", got, tt.want)
		}
	}
}

func ExampleAny() {
	fmt.Println(Any([]int64{1}))
	fmt.Println(Any([]time.Duration{1}))
}
