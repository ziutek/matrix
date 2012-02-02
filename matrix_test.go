package matrix

import (
	"testing"
	"fmt"
)

func TestIdentity(t *testing.T) {
	m := DenseI(11)
	rows, cols := m.Size()
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

func TestSlice(t *testing.T) {
	m := DenseZero(5, 6)
	m.Rand(-1, 1)
	s := m.Hslice(1, 4)
	fmt.Printf("m = %s\n", m)
	fmt.Printf("s = %s\n", s)
	s = m.Vslice(1, 4)
	fmt.Printf("m = %s\n", m)
	fmt.Printf("s = %s\n", s)
}


func TestAddTo(t *testing.T) {
	
}
