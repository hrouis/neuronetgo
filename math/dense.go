package math

/**
 * Dense matrix Type.
 */
type Dense struct {
	mat General
	capRows, capCols int
}


/**
 * Function to create a new dense matrix.
 */
func NewDense(r, c int, data []float64) *Dense {
	if r < 0 || c < 0 {
		panic("mat: negative dimension")
	}
	if data != nil && r*c != len(data) {
		panic(ErrShape)
	}
	if data == nil {
		data = make([]float64, r*c)
	}
	return &Dense{
		mat: General{
			Rows:   r,
			Cols:   c,
			Stride: c,
			Data:   data,
		},
		capRows: r,
		capCols: c,
	}
}