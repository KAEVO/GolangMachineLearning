package speedtests

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"math/rand"
	"time"
)

// createRandomMatrix create NxM random Matrix [0;100)
func createRandomMatrix(N, M int) (A *mat.Dense) {
	v := make([]float64, N*M)
	for i := 0; i < N*M; i++ {
		rand.Seed(int64(time.Now().UnixNano()))
		v[i] = rand.Float64() * 100
	}
	A = mat.NewDense(N, M, v)
	return
}

// scaleMatrix make inplace scaling of given matrix
func scaleMatrix(alpha float64, A *mat.Dense) {
	A.Scale(alpha, A)
}

// transposingMatrix return interface Matrix with structure of
// transposed matrix
func transposeMatrix(A *mat.Dense) mat.Matrix {
	return A.T()
}

// additionOfMatrices return result of addition of two matrices
func additionOfMatrices(A, B *mat.Dense) (C *mat.Dense) {
	Rows, Cols := A.Dims()
	C = mat.NewDense(Rows, Cols, nil)
	C.Add(A, B)
	return
}

// subtractOfMatrices return result of subtract of two matrices
func subtractOfMatrices(A, B *mat.Dense) (C *mat.Dense) {
	Rows, Cols := A.Dims()
	C = mat.NewDense(Rows, Cols, nil)
	C.Sub(A, B)
	return
}

// dotOfMatrices return dot of matrices
func dotOfMatrices(A, B *mat.Dense) (C *mat.Dense) {
	Rows, Cols := A.Dims()
	C = mat.NewDense(Rows, Cols, nil)
	C.Product(A, B)
	return
}

// determinantOfMatrix return determinant of given matrix
func determinantOfMatrix(A *mat.Dense) (d float64) {
	d = mat.Det(A)
	return
}

// eigensOfMatrix returns right eigen vectors and values
func eigensOfMatrix(A *mat.Dense) (eigenValues []complex128, eigenVectors *mat.CDense){
	var eig mat.Eigen
	if ok := eig.Factorize(A, mat.EigenRight); !ok {
		log.Fatal("Ei