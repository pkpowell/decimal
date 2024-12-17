package decimal

import "testing"

func TestNew(t *testing.T) {
	n := New(1.234, 3)
	t.Log("new", n)
	t.Log("new", n.Decimal())
}
