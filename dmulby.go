package matrix

// m *= a
func (m *Dense) MulBy(a *Dense) {
	m.checkEqualDims(a)
	for i := 0; i < m.rows; i++ {
		mr := m.v[i*m.stride:]
		ar := a.v[i*a.stride:]
		k := m.cols
		for k >= 2 {
			k--
			mr[k] *= ar[k]
			k--
			mr[k] *= ar[k]
		}
		if k != 0 {
			mr[0] *= ar[0]
		}
	}
}
