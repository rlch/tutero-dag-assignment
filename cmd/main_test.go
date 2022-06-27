package main

import (
	"fmt"
	"testing"
)

func BenchmarkDrive(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		if n, err := drive(); err != nil {
			b.Error(err)
		} else if err == nil {
			b.ReportMetric(float64(n), "steps")
		}
	}
}
