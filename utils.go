package matrix

import "fmt"

func (m *Matrix) checkIndexes(i, k int) {
	if i < 0 || i >= m.rows {
		panic(fmt.Sprintf("row index (%d) out of range [0, %d]", i, m.rows))
	}
	if k < 0 || k >= m.cols {
		panic(fmt.Sprintf("column index (%d) out of range [0, %d]", k, m.cols))
	}
}

func checkEqualDims(m1, m2 *Matrix) {
	if m1.Rows() != m2.Rows() || m1.Cols() != m2.Cols() {
		panic("dimensions of matrices are not equal")
	}
}
