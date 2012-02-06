package matrix

import (
	"fmt"
	"encoding/json"
	"math/rand"
)

type Dense struct {
	v          []float64 // [row, row, ..., row]
	rows, cols int
	stride     int // distance between vertically adjacent values
}

// Creates new matrix that refers to v
func NewDense(rows, cols, stride int, v ...float64) *Dense {
	n := rows*stride
	if n > len(v) {
		panic("v is to small")
	}
	return &Dense{v: v[:n], rows: rows, cols: cols, stride: stride}
}

// Retuns new zero matrix
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

// Returns number of rows
func (m *Dense) Rows() int {
	return m.rows
}

// Returns number of columns
func (m *Dense) Cols() int {
	return m.cols
}

// Returns stride
func (m *Dense) Stride() int {
	return m.stride
}

// Returns internal buffer of values
func (m *Dense) Elems() []float64 {
	return m.v
}

// Returns value from row i and column k
func (m *Dense) Get(i, k int) float64 {
	m.checkIndexes(i, k)
	return m.v[i*m.stride+k]
}

// Sets value in row i and column k
func (m *Dense) Set(i, k int, a float64) {
	m.checkIndexes(i, k)
	m.v[i*m.stride+k] = a
}

// Sets all elements to a
func (m *Dense) SetAll(a float64) {
	for i := 0; i < m.rows; i++ {
		row := m.v[i*m.stride:]
		k := m.cols
		for k >= 2 {
			k--
			row[k] = a
			k--
			row[k] = a
		}
		if k != 0 {
			row[0] = a
		}
	}
}

// Sets all elements to: min + rand.Float64() * (max - min)
func (m *Dense) Rand(min, max float64) {
	s := max - min
	for i := 0; i < m.rows; i++ {
		row := m.v[i*m.stride:]
		for k := 0; k < m.cols; k++ {
			row[k] = min + s*rand.Float64()
		}
	}
}

// Sets all elements to: mean + stdDev*rand.NormFloat64()
func (m *Dense) RandNorm(mean, stdDev float64) {
	for i := 0; i < m.rows; i++ {
		row := m.v[i*m.stride:]
		for k := 0; k < m.cols; k++ {
			row[k] = mean + stdDev*rand.NormFloat64()
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

// Returns matrix as horizontal vector. Panics if cols != stride.
func (m *Dense) Hvec() *Dense {
	if m.cols != m.stride {
		panic("cols != stride")
	}
	return &Dense{v: m.v, rows: 1, cols: len(m.v), stride: len(m.v)}
}

// Returns matrix as vertical vector. Panics if cols != stride.
func (m *Dense) Vvec() *Dense {
	if m.cols != m.stride {
		panic("cols != stride")
	}
	return &Dense{v: m.v, rows: len(m.v), cols: 1, stride: 1}
}

// Returns true if matrices are equal
func (m *Dense) Equal(a *Dense) bool {
	if m.rows != a.rows || m.cols != a.cols {
		return false
	}
	for i := 0; i < m.rows; i++ {
		mr := m.v[i*m.stride:]
		ar := a.v[i*a.stride:]
		for k := 0; k < m.cols; k++ {
			if mr[k] != ar[k] {
				return false
			}
		}
	}
	return true
}

func (m *Dense) String() string {
	rows, cols := m.Size()
	s := "["
	for i := 0; i < rows; i++ {
		if i != 0 {
			s += "\n "
		}
		o := ""
		for k := 0; k < cols; k++ {
			s += fmt.Sprintf("%s%-g", o, m.Get(i, k))
			if k == 0 {
				o = " "
			}
		}
	}
	return s + "]"
}

func (m *Dense) MarshalJSON() ([]byte, error) {
	s := struct{Cols, Stride int; Elems []float64}{m.cols, m.stride, m.v}
	return json.Marshal(s)
}

// Utils

func (m *Dense) checkIndexes(i, k int) {
	if i < 0 || i >= m.rows {
		panic(fmt.Sprintf("row index (%d) out of range [0, %d]", i, m.rows))
	}
	if k < 0 || k >= m.cols {
		panic(fmt.Sprintf("column index (%d) out of range [0, %d]", k, m.cols))
	}
}

func (m *Dense) checkEqualDims(a *Dense) {
	if m.Rows() != a.Rows() || m.Cols() != a.Cols() {
		panic("dimensions of matrices are not equal")
	}
}
