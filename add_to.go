package matrix

// m += a * s
func (m *Dense) AddTo(a *Dense, s float64) {
	checkEqualDims(m, a)

	var mi, ai int
	n := m.cols * m.rows

	switch s {
	case 1.0:
		for n := m.cols * m.rows; n >= 2; n -= 2 {
			m.v[mi] += a.v[ai]
			mi += m.stride
			ai += a.stride

			m.v[mi] += a.v[ai]
			mi += m.stride
			ai += a.stride
		}
		if n != 0 {
			m.v[mi] += a.v[ai]
		}
	case -1:
		for n := m.cols * m.rows; n >= 2; n -= 2 {
			m.v[mi] -= a.v[ai]
			mi += m.stride
			ai += a.stride

			m.v[mi] -= a.v[ai]
			mi += m.stride
			ai += a.stride
		}
		if n != 0 {
			m.v[mi] -= a.v[ai]
		}
	default:
		for n := m.cols * m.rows; n >= 2; n -= 2 {
			m.v[mi] += a.v[ai] * s
			mi += m.stride
			ai += a.stride

			m.v[mi] += a.v[ai] * s
			mi += m.stride
			ai += a.stride
		}
		if n != 0 {
			m.v[mi] += a.v[ai] * s
		}
	}
}
