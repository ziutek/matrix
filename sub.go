package matrix

// m = (a - b) * s
func (m *Matrix) Sub(a, b *Matrix, s float64) {
	checkEqualDims(m, a)
	checkEqualDims(m, b)

	var mi, ai, bi int
	n := m.cols * m.rows

	if s == 1.0 {
		for n := m.cols * m.rows; n >= 2; n -= 2 {
			m.v[mi] = a.v[ai] - b.v[bi]
			mi += m.stride
			ai += a.stride
			bi += b.stride

			m.v[mi] = a.v[ai] - b.v[bi]
			mi += m.stride
			ai += a.stride
			bi += b.stride
		}
		if n != 0 {
			m.v[mi] = a.v[ai] - b.v[bi]
		}
	} else {
		for n := m.cols * m.rows; n >= 2; n -= 2 {
			m.v[mi] = (a.v[ai] - b.v[bi]) * s
			mi += m.stride
			ai += a.stride
			bi += b.stride

			m.v[mi] = (a.v[ai] - b.v[bi]) * s
			mi += m.stride
			ai += a.stride
			bi += b.stride
		}
		if n != 0 {
			m.v[mi] = (a.v[ai] - b.v[bi]) * s
		}
	}
}
