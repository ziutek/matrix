package matrix

import (
	"fmt"
	"math/rand"
)

type Dense struct {
	v          []float64 // [row, row, ..., row]
	rows, cols int
	stride     int // distance between vertically adjacent values
}

// Retuns zero matrix
func DenseZero(rows, cols int) *Dense {
	return &Dense{
		v:      make([]float64, cols*rows),
		rows:   rows,
		cols:   cols,
		stride: cols,
	}
}

// Returns identity matrix
func DenseI(n int) *Dense {
	m := DenseZero(n, n)
	inc := n + 1
	for i := 0; i < len(m.v); i += inc {
		m.v[i] = 1
	}
	return m
}

// Returns dimensions of the matrix (rows, cols)
func (m *Dense) Size() (int, int) {
	return m.rows, m.cols
}

func (m *Dense) Rows() int {
	return m.rows
}

func (m *Dense) Cols() int {
	return m.cols
}

func (m *Dense) Get(i, k int) float64 {
	m.checkIndexes(i, k)
	return m.v[i*m.stride+k]
}

func (m *Dense) Set(i, k int, a float64) {
	m.checkIndexes(i, k)
	m.v[i*m.stride+k] = a
}

func (m *Dense) SetAll(a float64) {
	o := m.rows * m.stride
	for o != 0 {
		o -= m.stride
		row := m.v[o : o+m.cols]
		for k := 0; k < m.cols; k++ {
			row[k] = a
		}
	}
}

// Sets all elements to: min + rand.Float64() * (max - min)
func (m *Dense) Rand(min, max float64) {
	s := max - min
	o := m.rows * m.stride
	for o != 0 {
		o -= m.stride
		row := m.v[o : o+m.cols]
		for k := 0; k < m.cols; k++ {
			row[k] = min + s*rand.Float64()
		}
	}
}

func (m *Dense) RandNorm(mean, stdDev float64) {
	o := m.rows * m.stride
	for o != 0 {
		o -= m.stride
		row := m.v[o : o+m.cols]
		for k := 0; k < m.cols; k++ {
			row[k] = mean + stdDev*rand.Float64()
		}
	}
}

// Returns a slice of a matrix that contains rows from start to stop - 1
func (m *Dense) Hslice(start, stop int) *Dense {
	if start > stop || start < 0 || stop > m.rows {
		panic("bad indexes for horizontal slice")
	}
	return &Dense{
		v:      m.v[start*m.stride : stop*m.stride],
		rows:   stop - start,
		cols:   m.cols,
		stride: m.stride,
	}
}

// Returns a slice of a matrix that contains cols from start to stop - 1
func (m *Dense) Vslice(start, stop int) *Dense {
	if start > stop || start < 0 || stop > m.cols {
		panic("bad indexes for vertical slice")
	}
	return &Dense{
		v:      m.v[start : (m.rows-1)*m.stride+stop],
		rows:   m.rows,
		cols:   stop - start,
		stride: m.stride,
	}

}

func (m *Dense) String() string {
	rows, cols := m.Size()
	s := "["
	for i := 0; i < rows; i++ {
		if i != 0 {
			s += "\n"
		}
		for k := 0; k < cols; k++ {
			s += fmt.Sprintf(" %-g ", m.Get(i, k))
		}
	}
	return s + "]"
}
