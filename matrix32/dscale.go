package matrix32

// Scale performs: d = a * s
func (d Dense) Scale(a Dense, s float32) {
	d.checkDim(a)
	for i := 0; i < d.numrow; i++ {
		dr := d.v[i*d.stride:]
		ar := a.v[i*a.stride:]
		k, n := 0, d.numcol-1
		for k < n {
			dr[k+0] = ar[k+0] * s
			dr[k+1] = ar[k+1] * s
			k += 2
		}
		if k == n {
			dr[k] = ar[k] * s
		}
	}
}
