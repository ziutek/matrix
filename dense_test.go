package matrix

import (
	"testing"
)

func TestDenseI(t *testing.T) {
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

func TestDenseSlice(t *testing.T) {
	m := DenseZero(5, 6)
	m.Rand(-1, 1)
	start, stop := 1, 4

	s := m.Hslice(start, stop)
	if s.Rows() != stop-start || s.Cols() != m.Cols() {
		t.Fatal("Hslice: wrong size")
	}
	for i := 0; i < s.Rows(); i++ {
		for k := 0; k < s.Cols(); k++ {
			if s.Get(i, k) != m.Get(i+start, k) {
				t.Fatal("Hslice: elements don't match")
			}
		}
	}

	s = m.Vslice(start, stop)
	if s.Rows() != m.Rows() || s.Cols() != stop-start {
		t.Fatal("Vslice: wrong size")
	}
	for i := 0; i < s.Rows(); i++ {
		for k := 0; k < s.Cols(); k++ {
			if s.Get(i, k) != m.Get(i, k+start) {
				t.Fatal("Vslice: elements don't match")
			}
		}
	}

}

func TestDenseSetAll(t *testing.T) {
	m := DenseZero(5, 6)
	start, stop := 1, 4
	m.Vslice(start, stop).SetAll(1)
	for i := 0; i < m.Rows(); i++ {
		for k := 0; k < m.Cols(); k++ {
			v := m.Get(i, k)
			if k >= start && k < stop && v == 1 {
				continue
			}
			if (k < start || k >= stop) && v == 0 {
				continue
			}
			wrongValue(t, m, i, k)
		}
	}
}

func TestDenseMulBy(t *testing.T) {
	m := DenseZero(6, 5)
	m.SetAll(2)
	start, stop := 1, 4
	m.Hslice(start, stop).SetAll(3)
	s := m.Vslice(start, stop)
	s.MulBy(s)
	for i := 0; i < m.Rows(); i++ {
		for k := 0; k < m.Cols(); k++ {
			v := m.Get(i, k)
			if k >= start && k < stop {
				if i >= start && i < stop && v == 9 {
					continue
				}
				if (i < start || i >= stop) && v == 4 {
					continue
				}
			} else {
				if i >= start && i < stop && v == 3 {
					continue
				}
				if (i < start || i >= stop) && v == 2 {
					continue
				}
			}
			wrongValue(t, m, i, k)
		}
	}
}

// Utils

func wrongValue(t *testing.T, m *Dense, i, k int) {
	t.Fatalf("wrong value %g at (%d,%d)\n%s", m.Get(i, k), i, k, m)
}
