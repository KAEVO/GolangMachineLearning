package speedtests

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"testing"
)

// TestCreateRandomMatrix testing speed and result of work
// createRandomMatrix function
func TestCreateRandomMatrix(t *testing.T) {
	for key, value := range mapMatTest {
		t.Run(key, testCreateRandomMatrixFunc(value[0], value[1]))
	}
}

func testCreateRandomMatrixFunc(N, M int) func(t *testing.T) {
	return func(t *testing.T) {
		A := createRandomMatrix(N, M)
		Rows, Cols := A.Dims()
		fmt.Printf("Created matrix 'A' with d