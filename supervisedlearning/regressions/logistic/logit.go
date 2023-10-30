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
		lgr.lrate = 0.01
	}
	if lgr.iter == 0 {
		lgr.iter = 10000
	}
	if lgr.batchs == 0 {
		lgr.batchs = int(math.Round(float64(xrn) / 5.0))
	} else if lgr.batchs >= xrn {
		return BatchSizeInputError
	}
	if lgr.weights == nil {
		rweights := make([]float64, xcn)
		for i := range rweights {
			rweights[i] = rand.NormFloat64()
		}
		lgr.weights = mat.NewVecDense(len(rweights), rweights)
	}

	switch {
	case lgr.solver&GD != 0:
		gdSolver(x_train, y_train, lgr.weights, lgr.bias, lgr.lrate, lgr.iter)
	case lgr.solver&SGD != 0:
		sgdSolver(x_train, y_train, lgr.weights, lgr.bias, lgr.lrate, lgr.iter, lgr.batchs)
	default:
		sgdSolver(x_train, y_train, lgr.weights, lgr.bias, lgr.lrate, lgr.iter, lgr.batchs)

	}
	return nil
}

func (lgr *LGR) Predict(train core.Train, w *mat.VecDense) {
	b := mat.NewVecDense(len(*train), nil)

	x_test := mat.NewDense(len(*train), (*train)[0].Elements-1, nil)
	for index, row := range *train {
		for idx, elem := range row.Data {
			if idx != 0 {
				if idx > 0 {
					idx -= 1
				}
				x_test.Set(index, idx, elem)
			}
		}
	}

	lgr.res = mat.NewDense(len(*train), 1, nil)
	sigmoid(x_test, w, b, lgr.res)
}

func gdSolver(X *mat.Dense, y *mat.VecDense, w *mat.VecDense, b *mat.VecDense, lr float64, iter int) {
	h := mat.NewDense(y.Len(), 1, nil)
	grad := mat.NewVecDense(w.Len(), nil)
	for i := 0; i < iter; i++ {
		sigmoid(X, w, b, h)
		subg := mat.NewVecDense(y.Len(), nil)
		subg.SubVec(y, h.ColV