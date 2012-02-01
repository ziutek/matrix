package matrix

import (
	"math/rand"
)

type Dense struct {
	v                  []float64 // [row, row, ..., row]
	rows, cols, stride int
}

// Retuns zero matrix
func DenseZero(rows, cols int) *Dense {
	return &Dense{
		v:      make([]float64, cols*rows),
		rows:   rows,
		cols:   cols,
		stride: 1,
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
func (m *Dense) Dims() (int, int) {
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
	return m.v[m.stride*(i*m.cols+k)]
}

func (m *Dense) Set(i, k int, a float64) {
	m.checkIndexes(i, k)
	m.v[m.stride*(i*m.cols+k)] = a
}

func (m *Dense) SetAll(a float64) {
	i := m.rows * m.cols * m.stride
	for i > 0 {
		i--
		m.v[i] = a
	}
}

// Sets all elements to: min + rand.Float64() * (max - min)
func (m *Dense) Rand(min, max float64) {
	i := m.rows * m.cols * m.stride
	s := max - min
	for i > 0 {
		i--
		m.v[i] = min + s*rand.Float64()
	}
}

func (m *Dense) RandNorm(mean, stdDev float64) {
	i := m.rows * m.cols * m.stride
	for i > 0 {
		i--
		m.v[i] = mean + stdDev*rand.Float64()
	}
}


// Returns a slice of a matrix that contains rows from start to stop - 1
func (m *Dense) Hslice(start, stop int) *Dense {
	if start > stop || start < 0 || stop > m.rows {
		panic("bad indexes for horizontal slice")
	}
	inc := m.cols * m.stride
	return &Dense{
		v:      m.v[start*inc : stop*inc],
		rows:   stop - start,
		cols:   m.cols,
		stride: m.stride,
	}
}
