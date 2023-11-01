package kmeans

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/bayesiangopher/bayesiangopher/core"
)

//Speed of fitting model for Xclara dataset
func BenchmarkXclaraFitKmeans(b *testing.B) {
	r := core.CSVReader{Path: "../../../datasets/the_xclara_cluster_2.5k_dataset.csv"}
	train := r.Read(true)
	k := 3
	kmns := KMNS{}
	b.ReportAllocs()
	b.N = 10
	b.ResetTimer