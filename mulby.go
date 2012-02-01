package matrix

// m *= a
func (m *Dense) MulBy(a *Dense) {
	checkEqualDims(m, a)

	var mi, ai int
	n := m.cols * m.rows

	for n := m.cols * m.rows; n >= 2; n -= 2 {
		m.v[mi] *= a.v[ai]
		mi += m.stride
		ai += a.stride

		m.v[mi] *= a.v[ai]
		mi += m.stride
		ai += a.stride
	}
	if n != 0 {
		m.v[mi] *= a.v[ai]
	}
}
