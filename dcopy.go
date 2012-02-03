package matrix

// m = a
func (m *Dense) Copy(a *Dense) {
	m.checkEqualDims(a)
	for i := 0; i < m.rows; i++ {
		ao := i * a.stride
		copy(m.v[i*m.stride:], a.v[ao:ao+m.cols])
	}
}
