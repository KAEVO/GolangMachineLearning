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
		fmt.Printf("Created matrix 'A' with dimension: %d, %d \n", Rows, Cols); matPrint(A)
	}
}

// TestScaleMatrix testing speed and result of work
// scaleMatrix (gonum: Scale) function
func TestScaleMatrix(t *testing.T) {
	A := createRandomMatrix(16, 16)
	check := A.At(0, 0) * 5.25
	scaleMatrix(5.25, A)
	if A.At(0, 0) != check {
		t.Fatal("Масштабирование матрицы реализовано неправильно.")
	}
}

// TestTransposeMatrix testing speed and result of work
// transposeMatrix (gonum: T) function
func TestTransposeMatrix(t *testing.T) {
	A := createRandomMatrix(16, 16)
	B := transposeMatrix(A)
	if A.At(2, 3) != B.At(3 , 2) {
		t.Fatal("Транспонирование матрицы реализовано неправильно.")
	}
}

// TestAdditionOfMatrices testing speed and result of work
// additionOfMatrices (gonum: Add) function
func TestAdditionOfMatrices(t *testing.T) {
	A := createRandomMatrix(16, 16)
	B := createRandomMatrix(16, 16)
	check := A.At(0, 0) + B.At(0, 0)
	C := additionOfMatrices(A, B)
	if C.At(0, 0) != check {
		t.Fatal("Матрицы сложились неправильно.")
	}
}

// TestSubtractOfMatrices testing speed and result of work
// suntractOfMatrices (gonum: Sub) function
func TestSubtractOfMatrices(t *testing.T) {
	A := createRandomMatrix(16, 16)
	B := createRandomMatrix(16, 16)
	check := A.At(0, 0) - B.At(0, 0)
	C := subtractOfMatrices(A, B)
	if C.At(0, 0) != check {
		t.Fatal("Вычитание матриц реализовано неправильно.")
	}
}

// TestDotOfMatrices testing speed and result of work
// dotOfMatrices (gonum: Product) function
func TestDotOfMatrices(t *testing.T