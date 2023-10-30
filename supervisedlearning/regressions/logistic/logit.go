package logit

import (
	"errors"
	"math"
	"math/rand"

	"github.com/bayesiangopher/bayesiangopher/core"
	"gonum.org/v1/gonum/mat"
)

var (
	BatchSizeInputError = errors.New("Batch size exceeds the size of data")
)

type LGRsolver int

const (
	// Gradient Descending solver
	GD LGRsolver = 1 << iota
	// Stochastic Gradient Descending solver
	SGD
)

type LGR struct {
	bias        *mat.VecDense //bias; default: random vetor
	lrate       float64       //learning rate
	iter        int           //number of iteration
	batchs      int           //batch size
	weights     *mat.VecDense
	solver      LGRsolver //solving method: GD, SGD
	multi_class bool
	res         *mat.Dense
}

func (lgr *LGR) Fit(train core.Train, targetColumn int) (err error) {
	// Prepare data:
	x_train := mat.NewDense(len(*train), (*train)[0].Elements-1, nil)
	y_train := mat.NewVecDense(len(*train), nil)
	for index, row := range *train {
		for idx, elem := range row.Data {
			if idx == targetColumn {
				y_train.SetVec(index, elem)
			} else {
				if idx > targetColumn {
					idx -= 1
				}
				x_train.Set(index, idx, elem)
			}
		}
	}

	xrn, xcn := x_train.Dims()
	//Set optinal parametrs
	if lgr.bias == nil {
		rbias := make([]float64, xrn)
		for i := range rbias {
			rbias[i] = rand.NormFloat64()
		}
		lgr.bias = mat.NewVecDense(len(rbias), rbias)
	}
	if lgr.lrate == 0.0 {
		lg