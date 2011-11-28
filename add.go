package matrix

func (m *Matrix) Add(m1 *Matrix) {
	checkEqualDims(m, m1)

	var i, i1 int
	n := m.cols * m.rows

	for n := m.cols * m.rows; n >= 2; n -= 2 {
		m.a[i] += m1.a[i1]
		i += m.stride
		i1 += m1.stride

		m.a[i] += m1.a[i1]
		i += m.stride
		i1 += m1.stride
	}
	if n != 0 {
		m.a[i] += m1.a[i1]
	}
}
