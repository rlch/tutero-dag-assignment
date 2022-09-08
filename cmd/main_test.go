package main

import "testing"

func BenchmarkDrive(b *testing.B) {
	steps := 0.0
	for i := 0; i < b.N; i++ {
		if n, err := drive(); err != nil {
			b.Error(err)
		} else {
			steps += float64(n) / float64(b.N)
		}
	}
	b.ReportMetric(float64(steps), "steps/op")
}
