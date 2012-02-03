package matrix

// m *= s
func (m *Dense) ScaleBy(s float64) {
	for i := 0; i < m.rows; i++ {
		mr := m.v[i*m.stride:]
		k := m.cols
		for k >= 2 {
			k--
			mr[k] *= s
			k--
			mr[k] *= s
		}
		if k != 0 {
			mr[0] *= s
		}
	}
}
