package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dim int = 60

type Particle struct {
	X, Y int
}

func (p *Particle) has_neighbor(grid [][]int) bool {
	has_neighbor := false
	res := [][]int{
		{p.X, p.Y},
		{wrap(p.X, 1), p.Y},
		{wrap(p.X, -1), p.Y},
		{p.X, wrap(p.Y, 1)},
		{p.X, wrap(p.Y, -1)},
	}
	for i := range res {
		if grid[res[i][0]][res[i][1]] != 0 {
			has_neighbor = true
			break
		}
	}
	return has_neighbor
}

func (p *Particle) move() {
	n := rand.Intn(4)
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

func draw(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", grid[i][j])
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func main() {
	grid := make([][]int, dim)
	for i := range grid {
		grid[i] = make([]int, dim)
	}

	seed := time.Now().Unix()
	rand.Seed(seed)

	//for i := 0; i < 10; i++ {
	//	grid[rand.Intn(dim)][rand.Intn(dim)] += 1
	//}
	grid[dim/2][dim/2] = 1

	for i := 0; i < 500; i++ {
		walker := &Particle{X: rand.Intn(dim), Y: rand.Intn(dim)}
		steps := 0
		for steps < 5000 {
			if walker.has_neighbor(grid) {
				grid[walker.X][walker.Y] += 1
				break
			}
			walker.move()
			steps += 1

		}
	}
	draw(grid)
}
