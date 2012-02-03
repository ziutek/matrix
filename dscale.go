package matrix

// m = a * s
func (m *Dense) Scale(a *Dense, s float64) {
	m.checkEqualDims(a)
	for i := 0; i < m.rows; i++ {
		mr := m.v[i*m.stride:]
		ar := a.v[i*a.stride:]
		k := m.cols
		for k >= 2 {
			k--
			mr[k] = ar[k] * s
			k--
			mr[k] = ar[k] * s
		}
		if k != 0 {
			mr[0] = ar[0] * s
		}
	}
}
