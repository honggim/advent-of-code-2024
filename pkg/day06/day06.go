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

type Position struct {
	Point
	Direction
}

type Game struct {
	puzzle     [][]rune
	yMax, xMax int

	position Position

	traveled map[string]bool
	rocks    map[string]bool
}

func New(puzzle [][]rune, x, y int) *Game {
	position := Position{
		Point: Point{
			x: x,
			y: y,
		},
		Direction: N,
	}
	traveled := map[string]bool{
		position.Token(): true,
	}

	return &Game{
		puzzle:   puzzle,
		yMax:     len(puzzle) - 1,
		xMax:     len(puzzle[0]) - 1,
		position: position,
		traveled: traveled,
		rocks:    make(map[string]bool),
	}
}

func (g *Game) TravelDistance() int {
	return len(g.traveled)
}

func (g *Game) Simulate() (int, bool) {
	for !g.isEnd() {
		next := g.next()
		switch g.peek(next) {
		case '.':
			g.move()
		case '^':
			g.move()
		case '#':
			rockAddy := fmt.Sprintf("%s|%v", next.Token(), g.position.Direction)
			if ok := g.rocks[rockAddy]; ok {
				return 0, false
			}
			g.rocks[rockAddy] = true
			g.turn()
		}
	}

	return g.TravelDistance(), true
}

/*
// what if next '.' was changed to '#'
// check right turn direction for previous rock
func (g *Game) potentialLoop() bool {
	switch g.position.Direction {
	case N: // check E, x+
		for x := g.position.Point.x; x <= g.xMax; x++ {
			token := Point{x: x, y: g.position.y}.Token()
			if g.rocks[token] {
				return true
			}
		}
	case E: // check S, y+
		for y := g.position.Point.y; y <= g.yMax; y++ {
			token := Point{x: g.position.x, y: y}.Token()
			if g.rocks[token] {
				return true
			}
		}
	case S: // check W, x-
		for x := g.position.Point.x; x >= 0; x-- {
			token := Point{x: x, y: g.position.y}.Token()
			if g.rocks[token] {
				return true
			}
		}
	case W: // check N, y-
		for y := g.position.Point.y; y >= 0; y-- {
			token := Point{x: g.position.x, y: y}.Token()
			if g.rocks[token] {
				return true
			}
		}
	}

	return false
}
*/

func (g *Game) next() (p Point) {
	next := Point{
		x: g.position.x,
		y: g.position.y,
	}
	switch g.position.Direction {
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
	switch g.position.Direction {
	case N:
		g.position.Direction = E
	case E:
		g.position.Direction = S
	case S:
		g.position.Direction = W
	case W:
		g.position.Direction = N
	}
}

func (g *Game) peek(next Point) rune {
	char := g.puzzle[next.y][next.x]
	// fmt.Printf("next: %v, char: %s\n", next, string(char))

	return char
}

func (g *Game) move() {
	g.position.Point = g.next()
	g.traveled[g.position.Token()] = true
}

func (g *Game) isEnd() bool {
	if g.position.Point.x == 0 || g.position.Point.y == 0 {
		return true
	}
	if g.position.Point.x == g.xMax || g.position.Point.y == g.yMax {
		return true
	}

	return false
}

func part1(puzzle [][]rune, x, y int) int {
	game := New(puzzle, x, y)
	count, _ := game.Simulate()
	return count
}

func part2(puzzle [][]rune, xStart, yStart int) int {
	count := 0
	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[0]); x++ {
			// skip if already a rock
			if puzzle[y][x] == '#' {
				continue
			}
			puzzle[y][x] = '#'
			game := New(puzzle, xStart, yStart)
			if _, ok := game.Simulate(); !ok {
				// printPuzzle(p)
				count++
			}
			puzzle[y][x] = '.' // revert
		}
	}

	return count
}

func printPuzzle(puzzle [][]rune) {
	for _, row := range puzzle {
		fmt.Println(string(row))
	}
}
