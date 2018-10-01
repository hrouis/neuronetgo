package math

type Matrix interface {
	// returns the dimensions of a matrix
	Dims() (r, c int)
	// returns value of (i,j) point in the matrix
	At(i, j int) float64
	// returns the transpose for the matrix.
	T() Matrix
	Put(i,j int, val float64) ()
}


// General Matrix structure .

type General struct {
	Rows, Cols int
	Data []float64
	Stride int
}

/*
* Function sum returns the sum of all the matrix elements.
 */
func Sum(a Matrix) float64 {
	r, c := a.Dims()
	var sum float64

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			sum += a.At(i, j)
		}
	}
	return sum
}

/*
* Function multiply two matrices.

 */
func Mul(a Matrix, b Matrix, result Matrix)  {
	r, c := a.Dims()
	_, c1 := b.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c1; j++ {
			var val float64
			for k := 0; k < c; k++ {
				 val += a.At(i,k) * b.At(k,j)
			}
			result.Put(i,j, val)
		}
	}
}
