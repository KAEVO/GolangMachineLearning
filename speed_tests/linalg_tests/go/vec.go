
// Set of wrappers for gonum vector algebra

package speedtests

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

// createRandomVector create random N vector [0;100)
func createRandomVector(N int) *mat.VecDense {
	v := mat.NewVecDense(N, nil)
	for i := 0; i < N; i++ {
		rand.Seed(int64(time.Now().UnixNano()))
		v.SetVec(i, rand.Float64() * 100)
	}
	return v
}

// scaleVector scale vector v
func scaleVector(v *mat.VecDense, alpha float64) {
	v.ScaleVec(alpha, v)
}

// frobeniusNormOfVector return Frobenius norm of vector
func frobeniusNormOfVector(v *mat.VecDense) (f float64) {
	f = mat.Norm(v, 2)
	return
}

// additionOfVectors return sum of vectors v, u,
// if alpha is not 0 return scaled sum of vectors (v + alpha * u)
func additionOfVectors(v, u *mat.VecDense, alpha float64) (w *mat.VecDense) {
	w = mat.NewVecDense(v.Len(), nil)
	if alpha == 0.0 {
		w.AddVec(u, v)
	}
	if alpha != 0.0 {
		w.AddScaledVec(v, alpha, u)
	}
	return