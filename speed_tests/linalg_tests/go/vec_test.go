package speedtests

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

// TestCreateRandomVector testing speed and result of work
// createRandomVector function
func TestCreateRandomVector(t *testing.T) {
	for key, value := range mapVecTest {
		t.Run(key, testCreateRandomVectorFunc(value))
	}
}

func testCreateRandomVectorFunc(N int) func(t *testing.T) {
	return func(t *testing.T) {
		v := createRandomVector(N)
		fmt.Printf("Created vector 'v' with dimension: %d \n", v.Len()); vecPrint(v)
	}
}

// TestScaleVector testing correctness of
// scaleVector (gonum: ScaleVec) function
func TestScaleVector(t *testing.T) {
	v := createRandomVector(1024)
	check := v.AtVec(5)
	fmt.Printf("Created vector 'v' with dimension: %d\n", v.Len()); vecPrint(v)
	scaleVector(v, 5.25)
	fmt.Println("Vector 'v' after scaling:"); vecPrint(v)
	if v.AtVec(5) / check != 5.25 {
		t.Fatal("Ошибка масштабирования.")
	}
}

// TestFrobeniusNormOfVector testing correctness of
// frobeniusNormOfVector (gonum: Norm(v, 2) function
func TestFrobeniusNormOfVector(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v := mat.NewVecDense(len(data), data)
	fmt.Printf("Created vector 'v' with dimension: %d\n", v.Len()); vecPrint(v)
	f := frobeniusNormOfVector(v)
	if f != 19.621416870348583  {
		t.Fatal("Неправильно найдена норма Фробениуса.")
	}
	fmt.Printf("Frobenius norm of vector 'v': %v \n", f)
}

// TestAdditionOfVectors testing correctness of
// additionOfVectors (gonum: AddVec) function
func TestAdditionOfVectors(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataTest := []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	v := mat.NewVecDense(len(data), data)
	u := mat.NewVecDense(len(data), data)
	resultTest := mat.NewVecDense(len(dataTest), dataTest)
	w := additionOfVectors(u, v, 0.0)
	if froTest := frobeniusNormOfVector(w); froTest != frobeniusNormOfVector(resultTest) {
		t.Fatal("Векторы сложились неправильно.")
	}
}

// TestSubtractOfVectors testing correctness of
// subtractOfVectors (gonum: SubVec) function
func TestSubtractOfVectors(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataTest := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	v := mat.NewVecDense(len(data), data)
	u := mat.NewVecDense(len(data), data)
	resultTest := mat.NewVecDense(len(dataTest), dataTest)
	w := subtractOfVectors(u, v)
	if froTest := frobeniusNormOfVector(w); froTest != frobeniusNormOfVector(resultTest) {
		t.Fatal("Вычитание выполнено с ошибкой.")
	}
}

// TestDotOfVectors testing correctness of
// dotOfVectors (gonum: Dot) function
func TestDotOfVectors(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testResult := 385.0
	v := mat.NewVecDense(len(data), data)
	u := mat.NewVecDense(len(data), data)
	resultTest := dotOfVectors(v, u)
	if resultTest != testResult {
		t.Fatal("Умножение выполнено с ошибкой.")
	}
}

// BenchmarkCreateRandomVector testing speed and memory use of
// cre