package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
	"github.com/PhillipNordwal/gobitrand"
)

var (
	dim              int = 400
	particle_count   int = 40000
	nucleation_sites int = 10
)

type Particle struct {
	X, Y int
}

func (p *Particle) has_neighbor(grid [][]int) bool {
	res := [][]int{
		{p.X, p.Y},
		{wrap(p.X, 1), p.Y},
		{wrap(p.X, -1), p.Y},
		{p.X, wrap(p.Y, 1)},
		{p.X, wrap(p.Y, -1)},
	}
	for _, val := range res {
		if grid[val[0]][val[1]] != 0 {
			return true
		}
	}
	return false
}

func (p *Particle) move() {
	n := gobitrand.Two_bits()
	switch n {
	case 0:
		p.X = wrap(p.X, 1)
	case 1:
		p.X = wrap(p.X, -1)
	case 2:
		p.Y = wrap(p.Y, 1)
	case 3:
		p.Y = wrap(p.Y, -1)
	}
}

func wrap(n, offset int) int {
	res := n + offset
	switch {
	case res < 0:
		return dim - 1
	case res == dim:
		return 0
	}
	return res
}

func writeToFile(grid [][]int) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				fmt.Fprintln(w, i, j, grid[i][j])
			}
		}
	}
	return w.Flush()
}

func main() {
	f, _ := os.Create("my_profile.file")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	grid := make([][]int, dim)
	for i := range grid {
		grid[i] = make([]int, dim)
	}

	seed := time.Now().Unix()
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
	writeToFile(grid)
}
