package matrix

// m = a * b
func (m *Dense) Mul(a, b *Dense) {
	checkEqualDims(m, a)
	checkEqualDims(m, b)

	var mi, ai, bi int
	n := m.cols * m.rows

	for n := m.cols * m.rows; n >= 2; n -= 2 {
		m.v[mi] = a.v[ai] * b.v[bi]
		mi += m.stride
		ai += a.stride
		bi += b.stride

		m.v[mi] = a.v[ai] * b.v[bi]
		mi += m.stride
		ai += a.stride
		bi += b.stride
	}
	if n != 0 {
		m.v[mi] = a.v[ai] * b.v[bi]
	}
}
