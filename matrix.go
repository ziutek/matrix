package matrix

type Matrix struct {
	a                  []float64 // [row, row, ..., row]
	rows, cols, stride int
}

// Retuns zero matrix
func Zero(rows, cols int) *Matrix {
	return &Matrix{
		a:      make([]float64, cols*rows),
		rows:   rows,
		cols:   cols,
		stride: 1,
	}
}

func Identity(n int) *Matrix {
	m := Zero(n, n)
	inc := n + 1
	for i := 0; i < len(m.a); i += inc {
		m.a[i] = 1
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
	return m.a[m.stride*(i*m.cols+k)]
}

func (m *Matrix) Set(i, k int, a float64) {
	m.checkIndexes(i, k)
	m.a[m.stride*(i*m.cols+k)] = a
}

func (m *Matrix) SetAll(a float64) {
	i := m.rows * m.cols * m.stride
	for i > 0 {
		i--
		m.a[i] = a
	}
}

// Returns a slice of a matrix that contains rows from i1 to i2 - 1
func (m *Matrix) Hslice(i1, i2 int) *Matrix {
	if i1 > i2 || i1 < 0 || i2 > m.rows {
		panic("bad indexes for horizontal slice")
	}
	inc := m.cols * m.stride
	return &Matrix{
		a:      m.a[i1*inc : i2*inc],
		rows:   i2 - i1,
		cols:   m.cols,
		stride: m.stride,
	}
}
