package matrix

import (
	"fmt"
)

type Dense struct {
	v      []float64 // [row, row, ..., row]
	numrow int
	numcol int
	stride int // The distance between vertically adjacent elements.
}

// AsDense makes new dense matrix that refers to v
func AsDense(numrow, numcol int, v []float64) Dense {
	n := numrow * numcol
	if n > len(v) {
		panic("matrix: AsDense: numrow*numcol > len(v)")
	}
	return Dense{v: v[:n], numrow: numrow, numcol: numcol, stride: numcol}
}

// MakeDense allocates new dense matrix and initializes its first elements to
// values specified by iv.
func MakeDense(numrow, numcol int, iv ...float64) Dense {
	v := make([]float64, numrow*numcol)
	copy(v, iv)
	return AsDense(numrow, numcol, v)
}

// IsValid reports whether d represents valid matrix value.
func (d Dense) IsValid() bool {
	return len(d.v) > 0 && d.numrow > 0 && d.numcol > 0 && d.stride > 0
}

// SetAll sets all elements of d to a.
func (d Dense) SetAll(a float64) {
	for i := 0; i < d.numrow; i++ {
		row := d.v[i*d.stride:]
		k := d.numcol - 1
		for k > 0 {
			row[k] = a
			row[k-1] = a
			k -= 2
		}
		if k == 0 {
			row[0] = a
		}
	}
}

// SetI sets elements of d to create identity matrix (panics if d is not square).
func (d Dense) SetIdentity() {
	if d.numrow != d.numcol {
		panic("matrix: SetI on non square matrix")
	}
	d.SetAll(0)
	for i := 0; i < len(d.v); i += d.stride + 1 {
		d.v[i] = 1
	}
}

// Size returns dimensions of the matrix (rows, cols).
func (d Dense) Size() (int, int) {
	return d.numrow, d.numcol
}

// NumRow returns number of rows.
func (d Dense) NumRow() int {
	return d.numrow
}

// NumCol returns number of columns.
func (d Dense) NumCol() int {
	return d.numcol
}

// Stride returns distance between vertically adjacent elements.
func (d Dense) Stride() int {
	return d.stride
}

// Elems returns internal buffer of elements. Be careful when use returned slice.
// For example you can not assume that its length corresponds to dimensions of d.
func (d Dense) Elems() []float64 {
	return d.v
}

// Get returns element from row i, column k.
func (d Dense) Get(i, k int) float64 {
	return d.v[i*d.stride+k]
}

// Set sets element in row i and column k.
func (d Dense) Set(i, k int, a float64) {
	d.v[i*d.stride+k] = a
}

// Rows returns a slice of a matrix that contains rows from start to stop-1.
func (d Dense) Rows(start, stop int) Dense {
	if start > stop || start < 0 || stop > d.numrow {
		panic("matrix: bad indexes for horizontal slice")
	}
	return Dense{
		v:      d.v[start*d.stride : stop*d.stride],
		numrow: stop - start,
		numcol: d.numcol,
		stride: d.stride,
	}
}

// Cols returns a slice of a matrix that contains columns from start to stop-1.
func (d Dense) Cols(start, stop int) Dense {
	if start > stop || start < 0 || stop > d.numcol {
		panic("matrix: bad indexes for vertical slice")
	}
	return Dense{
		v:      d.v[start : (d.numrow-1)*d.stride+stop],
		numrow: d.numrow,
		numcol: stop - start,
		stride: d.stride,
	}

}

// AsRow returns horizontal vector that refers to d. Panics if cols != stride.
func (d Dense) AsRow() Dense {
	if d.numcol != d.stride {
		panic("matrix: AsRow: numcol != stride")
	}
	return Dense{v: d.v, numrow: 1, numcol: len(d.v), stride: len(d.v)}
}

// AsCol returns vertical vector that refers to d. Panics if numcol != stride.
func (d Dense) AsCol() Dense {
	if d.numcol != d.stride {
		panic("matrix: AsCol: numcol != stride")
	}
	return Dense{v: d.v, numrow: len(d.v), numcol: 1, stride: 1}
}

// Equal returns true if matrices are equal
func (d Dense) Equal(a Dense) bool {
	if d.numrow != a.numrow || d.numcol != a.numcol {
		return false
	}
	for i := 0; i < d.numrow; i++ {
		dr := d.v[i*d.stride:]
		ar := a.v[i*a.stride:]
		for k := 0; k < d.numcol; k++ {
			if dr[k] != ar[k] {
				return false
			}
		}
	}
	return true
}

func (d Dense) Format(f fmt.State, c rune) {
	format := "%g"
	switch c {
	case 'f':
		format = "%f"
	case 'e':
		format = "%e"
	case 'G':
		format = "%G"
	case 'E':
		format = "%E"
	case 'F':
		format = "%F"
	}
	lastrow, lastcol := d.Size()
	lastrow--
	lastcol--
	f.Write([]byte{'{'})
	for i := 0; i <= lastrow; i++ {
		if i > 0 {
			f.Write([]byte{'\n'})
		}
		for k := 0; k <= lastcol; k++ {
			fmt.Fprintf(f, format, d.Get(i, k))
			if k < lastcol {
				f.Write([]byte{','})
			}
		}
		if i < lastrow {
			f.Write([]byte{','})
		}
	}
	f.Write([]byte{'}'})
}

// Utils

func (d Dense) checkDim(a Dense) {
	if d.numrow != a.numrow || d.numcol != a.numcol {
		panic("matrix: dimensions not equal")
	}
}
