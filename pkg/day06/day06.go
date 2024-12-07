package day06

import "fmt"

type Point struct {
	x, y int
}

func (p Point) Token() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

type Game struct {
	puzzle     [][]rune
	yMax, xMax int
	position   Point
	facing     Direction

	traveled map[string]bool
}

func New(puzzle [][]rune, x, y int) *Game {
	start := Point{x, y}
	traveled := map[string]bool{
		start.Token(): true,
	}

	return &Game{
		puzzle:   puzzle,
		yMax:     len(puzzle) - 1,
		xMax:     len(puzzle[0]) - 1,
		position: start,
		facing:   N,
		traveled: traveled,
	}
}

func (g *Game) Count() int {
	return len(g.traveled)
}

func (g *Game) Simulate() int {
	for !g.isEnd() {
		switch g.peek() {
		case '.':
			g.move()
		case '^':
			g.move()
		case '#':
			g.turn()
		}
	}

	return g.Count()
}

func (g *Game) next() (p Point) {
	next := Point{
		x: g.position.x,
		y: g.position.y,
	}
	switch g.facing {
	case N:
		next.y--
	case E:
		next.x++
	case S:
		next.y++
	case W:
		next.x--
	}

	return next
}

func (g *Game) turn() {
	switch g.facing {
	case N:
		g.facing = E
	case E:
		g.facing = S
	case S:
		g.facing = W
	case W:
		g.facing = N
	}
}

func (g *Game) peek() rune {
	next := g.next()

	val := g.puzzle[next.y][next.x]
	// fmt.Printf("next: %v, val: %s\n", next, string(val))

	return val
}

func (g *Game) move() {
	g.position = g.next()
	g.traveled[g.position.Token()] = true
}

func (g *Game) isEnd() bool {
	if g.position.x == 0 || g.position.y == 0 {
		return true
	}
	if g.position.x == g.xMax || g.position.y == g.yMax {
		return true
	}

	return false
}

func part1(puzzle [][]rune, x, y int) int {
	game := New(puzzle, x, y)
	return game.Simulate()
}
