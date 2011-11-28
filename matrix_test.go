package matrix

import (
	"testing"
)

func TestIdentity(t *testing.T) {
	m := Identity(11)
	rows, cols := m.Dim()
	for i := 0; i < rows; i++ {
		for k := 0; k < cols; k++ {
			a := m.Get(i, k)
			if i == k {
				if a != 1 {
					t.Fatalf("(%d, %d) == %f != 1", i, k, a)
				}
			} else {
				if a != 0 {
					t.Fatalf("(%d, %d) == %f != 0", i, k, a)
				}
			}
		}
	}
}
