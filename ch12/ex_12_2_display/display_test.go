package display

import (
	"testing"
)

func TestDisplayCycle(t *testing.T) {
	type cycle struct {
		val  int
		tail *cycle
	}
	var c cycle
	c = cycle{42, &c}
	Display("c", c)
}
