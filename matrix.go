package matrix

type Matrix struct {
	v                  []float64 // [row, row, ..., row]
	rows, cols, stride int
}

// Retuns zero matrix
func Zero(rows, cols int) *Matrix {
	return &Matrix{
		v:      make([]float64, cols*rows),
		rows:   rows,
		cols:   cols,
		stride: 1,
	}
}

func Identity(n int) *Matrix {
	m := Zero(n, n)
	inc := n + 1
	for i := 0; i < len(m.v); i += inc {
		m.v[i] = 1
	}
	return m
}

func Ones(n int) *Matrix {
	m := Zero(n, n)
	m.SetAll(1)
	return m
}

// Returns dimensions of the matrix (rows, cols)
func (m *Matrix) Dims() (int, int) {
	return m.rows, m.cols
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) Get(i, k int) float64 {
	m.checkIndexes(i, k)
	return m.v[m.stride*(i*m.cols+k)]
}

func (m *Matrix) Set(i, k int, a float64) {
	m.checkIndexes(i, k)
	m.v[m.stride*(i*m.cols+k)] = a
}

func (m *Matrix) SetAll(a float64) {
	i := m.rows * m.cols * m.stride
	for i > 0 {
		i--
		m.v[i] = a
	}
}

// Returns a slice of a matrix that contains rows from start to stop - 1
func (m *Matrix) Hslice(start, stop int) *Matrix {
	if start > stop || start < 0 || stop > m.rows {
		panic("bad indexes for horizontal slice")
	}
	inc := m.cols * m.stride
	return &Matrix{
		v:      m.v[start*inc : stop*inc],
		rows:   stop - start,
		cols:   m.cols,
		stride: m.stride,
	}
}
