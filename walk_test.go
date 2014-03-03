package main

import (
	"testing"
	"math/rand"
)

// bmark is used to run benchmarks with known seeds
func bmark(seed int64) {
	grid := make([][]int, dim)
	for i := range grid {
		grid[i] = make([]int, dim)
	}
	rand.Seed(seed)

	for i := 0; i < nucleation_sites; i++ {
		grid[rand.Intn(dim)][rand.Intn(dim)] += 1
	}

	for i := 0; i < particle_count; i++ {
		walker := &Particle{X: rand.Intn(dim), Y: rand.Intn(dim)}
		steps := 0
		for {
			if walker.has_neighbor(grid) {
				grid[walker.X][walker.Y] += 1
				break
			}
			walker.move()
			steps += 1
		}
	}
}

// TestWrap4n1 tests that wrapping works on numbers that don't wrap wrap(4, -1)
func TestWrap4n1(t *testing.T){
	if wrap(4,-1) != 3 {
		t.Error("Test Wrap4n1, failed.")
	} else {
		t.Log("Test Wrap4n1, passed.")
	}
}

// BenchmarkS01 runs a benchmark with a seed of 01
func BenchmarkS01(b *testing.B){
	for i := 0; i < b.N; i++ {
		bmark(01)
	}
}

// BenchmarkS02 runs a benchmark with a seed of 02
func BenchmarkS02(b *testing.B){
	for i := 0; i < b.N; i++ {
		bmark(02)
	}
}

// BenchmarkS01 runs a benchmark with a seed of 03
func BenchmarkS03(b *testing.B){
	for i := 0; i < b.N; i++ {
		bmark(03)
	}
}
