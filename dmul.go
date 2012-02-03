package matrix

// m = a * b
func (m *Dense) Mul(a, b *Dense) {
	m.checkEqualDims(a)
	m.checkEqualDims(b)
	for i := 0; i < m.rows; i++ {
		mr := m.v[i*m.stride:]
		ar := a.v[i*a.stride:]
		br := b.v[i*b.stride:]
		k := m.cols
		for k >= 2 {
			k--
			mr[k] = ar[k] * br[k]
			k--
			mr[k] = ar[k] * br[k]
		}
		if k != 0 {
			mr[0] = ar[0] * br[0]
		}
	}
}
