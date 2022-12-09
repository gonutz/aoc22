package main

import "fmt"

func main() {
	traceTail := func(ropeLength int) int {
		rope := make([]pos, ropeLength)
		visited := make(map[pos]bool)
		for _, move := range input {
			dir := move.r
			steps := move.n

			for i := 0; i < steps; i++ {

				head := &rope[0]
				switch dir {
				case 'L':
					head.x--
				case 'R':
					head.x++
				case 'U':
					head.y++
				case 'D':
					head.y--
				}

				for r := 1; r < len(rope); r++ {
					head := rope[r-1]
					tail := &rope[r]

					if head.x < tail.x-1 {
						tail.x--
						tail.y += sign(head.y - tail.y)
					}

					if head.x > tail.x+1 {
						tail.x++
						tail.y += sign(head.y - tail.y)
					}

					if head.y < tail.y-1 {
						tail.y--
						tail.x += sign(head.x - tail.x)
					}

					if head.y > tail.y+1 {
						tail.y++
						tail.x += sign(head.x - tail.x)
					}

				}

				tail := rope[len(rope)-1]
				visited[tail] = true

			}
		}
		return len(visited)
	}
	fmt.Println("part 1)", traceTail(2))
	fmt.Println("part 2)", traceTail(10))
}

type pos struct {
	x, y int
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

var input = []struct {
	r rune
	n int
}{
	{r: 'L', n: 1},
	{r: 'R', n: 2},
	{r: 'L', n: 1},
	{r: 'U', n: 2},
	{r: 'D', n: 1},
	{r: 'R', n: 1},
	{r: 'D', n: 1},
	{r: 'L', n: 2},
	{r: 'D', n: 1},
	{r: 'U', n: 2},
	{r: 'D', n: 2},
	{r: 'U', n: 2},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'D', n: 1},
	{r: 'U', n: 2},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 1},
	{r: 'L', n: 1},
	{r: 'R', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'U', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 1},
	{r: 'U', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 1},
	{r: 'D', n: 1},
	{r: 'L', n: 2},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 2},
	{r: 'D', n: 2},
	{r: 'R', n: 1},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 2},
	{r: 'D', n: 1},
	{r: 'R', n: 1},
	{r: 'U', n: 1},
	{r: 'R', n: 2},
	{r: 'D', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 1},
	{r: 'L', n: 1},
	{r: 'U', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'U', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 2},
	{r: 'R', n: 1},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'R', n: 1},
	{r: 'L', n: 2},
	{r: 'D', n: 1},
	{r: 'U', n: 2},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'R', n: 2},
	{r: 'U', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'U', n: 1},
	{r: 'R', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'R', n: 1},
	{r: 'D', n: 1},
	{r: 'L', n: 2},
	{r: 'U', n: 2},
	{r: 'D', n: 3},
	{r: 'R', n: 3},
	{r: 'L', n: 2},
	{r: 'R', n: 3},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 2},
	{r: 'U', n: 3},
	{r: 'D', n: 2},
	{r: 'L', n: 2},
	{r: 'U', n: 2},
	{r: 'R', n: 3},
	{r: 'D', n: 3},
	{r: 'R', n: 1},
	{r: 'D', n: 3},
	{r: 'U', n: 3},
	{r: 'L', n: 3},
	{r: 'R', n: 1},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 3},
	{r: 'U', n: 3},
	{r: 'L', n: 1},
	{r: 'U', n: 2},
	{r: 'L', n: 3},
	{r: 'U', n: 1},
	{r: 'R', n: 3},
	{r: 'D', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 1},
	{r: 'L', n: 3},
	{r: 'D', n: 3},
	{r: 'U', n: 1},
	{r: 'R', n: 1},
	{r: 'U', n: 2},
	{r: 'R', n: 1},
	{r: 'L', n: 1},
	{r: 'U', n: 2},
	{r: 'L', n: 3},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 3},
	{r: 'L', n: 2},
	{r: 'D', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'L', n: 3},
	{r: 'R', n: 1},
	{r: 'U', n: 2},
	{r: 'L', n: 1},
	{r: 'R', n: 1},
	{r: 'U', n: 3},
	{r: 'D', n: 3},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 3},
	{r: 'R', n: 3},
	{r: 'D', n: 3},
	{r: 'L', n: 2},
	{r: 'R', n: 1},
	{r: 'U', n: 3},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'U', n: 3},
	{r: 'D', n: 3},
	{r: 'U', n: 1},
	{r: 'R', n: 2},
	{r: 'D', n: 1},
	{r: 'U', n: 2},
	{r: 'D', n: 1},
	{r: 'L', n: 2},
	{r: 'D', n: 3},
	{r: 'L', n: 3},
	{r: 'U', n: 2},
	{r: 'R', n: 1},
	{r: 'L', n: 3},
	{r: 'D', n: 1},
	{r: 'R', n: 2},
	{r: 'L', n: 2},
	{r: 'D', n: 3},
	{r: 'R', n: 2},
	{r: 'D', n: 1},
	{r: 'U', n: 2},
	{r: 'R', n: 3},
	{r: 'U', n: 3},
	{r: 'D', n: 1},
	{r: 'U', n: 2},
	{r: 'L', n: 1},
	{r: 'R', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 3},
	{r: 'U', n: 1},
	{r: 'R', n: 3},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'U', n: 4},
	{r: 'R', n: 3},
	{r: 'L', n: 2},
	{r: 'R', n: 1},
	{r: 'L', n: 4},
	{r: 'R', n: 2},
	{r: 'L', n: 4},
	{r: 'R', n: 3},
	{r: 'D', n: 3},
	{r: 'R', n: 2},
	{r: 'L', n: 4},
	{r: 'U', n: 4},
	{r: 'R', n: 3},
	{r: 'L', n: 3},
	{r: 'R', n: 1},
	{r: 'L', n: 4},
	{r: 'D', n: 3},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'R', n: 2},
	{r: 'D', n: 1},
	{r: 'L', n: 3},
	{r: 'U', n: 2},
	{r: 'R', n: 1},
	{r: 'U', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 1},
	{r: 'U', n: 4},
	{r: 'D', n: 2},
	{r: 'R', n: 3},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'L', n: 3},
	{r: 'U', n: 2},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 4},
	{r: 'U', n: 2},
	{r: 'D', n: 4},
	{r: 'L', n: 3},
	{r: 'U', n: 1},
	{r: 'L', n: 4},
	{r: 'R', n: 1},
	{r: 'D', n: 1},
	{r: 'R', n: 4},
	{r: 'U', n: 4},
	{r: 'D', n: 2},
	{r: 'R', n: 3},
	{r: 'D', n: 4},
	{r: 'R', n: 4},
	{r: 'D', n: 4},
	{r: 'U', n: 3},
	{r: 'D', n: 1},
	{r: 'U', n: 4},
	{r: 'R', n: 1},
	{r: 'D', n: 4},
	{r: 'U', n: 2},
	{r: 'R', n: 2},
	{r: 'D', n: 4},
	{r: 'R', n: 4},
	{r: 'L', n: 4},
	{r: 'R', n: 2},
	{r: 'D', n: 4},
	{r: 'U', n: 3},
	{r: 'L', n: 3},
	{r: 'D', n: 3},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'R', n: 1},
	{r: 'L', n: 2},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 4},
	{r: 'U', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'R', n: 1},
	{r: 'U', n: 4},
	{r: 'D', n: 1},
	{r: 'L', n: 1},
	{r: 'U', n: 4},
	{r: 'L', n: 1},
	{r: 'R', n: 2},
	{r: 'U', n: 4},
	{r: 'L', n: 3},
	{r: 'U', n: 2},
	{r: 'D', n: 2},
	{r: 'U', n: 4},
	{r: 'L', n: 3},
	{r: 'R', n: 2},
	{r: 'L', n: 3},
	{r: 'U', n: 4},
	{r: 'L', n: 4},
	{r: 'U', n: 2},
	{r: 'L', n: 4},
	{r: 'D', n: 3},
	{r: 'L', n: 2},
	{r: 'U', n: 3},
	{r: 'D', n: 4},
	{r: 'U', n: 1},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'D', n: 3},
	{r: 'R', n: 1},
	{r: 'L', n: 4},
	{r: 'U', n: 1},
	{r: 'R', n: 4},
	{r: 'U', n: 2},
	{r: 'R', n: 2},
	{r: 'U', n: 5},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'R', n: 1},
	{r: 'U', n: 5},
	{r: 'L', n: 5},
	{r: 'U', n: 1},
	{r: 'R', n: 4},
	{r: 'U', n: 1},
	{r: 'L', n: 1},
	{r: 'D', n: 5},
	{r: 'L', n: 3},
	{r: 'D', n: 1},
	{r: 'L', n: 4},
	{r: 'D', n: 4},
	{r: 'R', n: 4},
	{r: 'U', n: 3},
	{r: 'D', n: 3},
	{r: 'R', n: 5},
	{r: 'U', n: 1},
	{r: 'L', n: 3},
	{r: 'D', n: 1},
	{r: 'R', n: 4},
	{r: 'D', n: 1},
	{r: 'L', n: 3},
	{r: 'U', n: 4},
	{r: 'R', n: 4},
	{r: 'U', n: 4},
	{r: 'L', n: 5},
	{r: 'U', n: 4},
	{r: 'L', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 5},
	{r: 'U', n: 2},
	{r: 'R', n: 4},
	{r: 'U', n: 3},
	{r: 'R', n: 2},
	{r: 'D', n: 1},
	{r: 'U', n: 5},
	{r: 'R', n: 1},
	{r: 'D', n: 4},
	{r: 'U', n: 3},
	{r: 'L', n: 1},
	{r: 'R', n: 4},
	{r: 'D', n: 5},
	{r: 'U', n: 5},
	{r: 'R', n: 5},
	{r: 'U', n: 3},
	{r: 'D', n: 3},
	{r: 'L', n: 1},
	{r: 'U', n: 4},
	{r: 'D', n: 1},
	{r: 'U', n: 1},
	{r: 'D', n: 3},
	{r: 'R', n: 4},
	{r: 'L', n: 5},
	{r: 'D', n: 4},
	{r: 'L', n: 3},
	{r: 'U', n: 5},
	{r: 'D', n: 3},
	{r: 'L', n: 5},
	{r: 'R', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 2},
	{r: 'R', n: 5},
	{r: 'D', n: 4},
	{r: 'U', n: 2},
	{r: 'D', n: 2},
	{r: 'R', n: 5},
	{r: 'U', n: 3},
	{r: 'R', n: 4},
	{r: 'L', n: 5},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'U', n: 3},
	{r: 'D', n: 2},
	{r: 'L', n: 3},
	{r: 'R', n: 2},
	{r: 'U', n: 3},
	{r: 'D', n: 3},
	{r: 'U', n: 5},
	{r: 'L', n: 5},
	{r: 'R', n: 2},
	{r: 'U', n: 4},
	{r: 'R', n: 2},
	{r: 'D', n: 5},
	{r: 'U', n: 3},
	{r: 'D', n: 4},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 3},
	{r: 'R', n: 4},
	{r: 'U', n: 1},
	{r: 'D', n: 3},
	{r: 'U', n: 1},
	{r: 'L', n: 3},
	{r: 'R', n: 3},
	{r: 'U', n: 2},
	{r: 'R', n: 3},
	{r: 'U', n: 5},
	{r: 'L', n: 4},
	{r: 'R', n: 2},
	{r: 'L', n: 3},
	{r: 'D', n: 4},
	{r: 'R', n: 3},
	{r: 'U', n: 5},
	{r: 'R', n: 4},
	{r: 'L', n: 5},
	{r: 'R', n: 5},
	{r: 'L', n: 5},
	{r: 'D', n: 1},
	{r: 'U', n: 4},
	{r: 'D', n: 1},
	{r: 'U', n: 6},
	{r: 'D', n: 3},
	{r: 'U', n: 5},
	{r: 'R', n: 3},
	{r: 'L', n: 2},
	{r: 'U', n: 3},
	{r: 'R', n: 4},
	{r: 'U', n: 1},
	{r: 'R', n: 6},
	{r: 'U', n: 2},
	{r: 'R', n: 3},
	{r: 'L', n: 4},
	{r: 'U', n: 2},
	{r: 'D', n: 1},
	{r: 'R', n: 5},
	{r: 'U', n: 1},
	{r: 'D', n: 1},
	{r: 'R', n: 3},
	{r: 'D', n: 3},
	{r: 'L', n: 5},
	{r: 'D', n: 1},
	{r: 'R', n: 3},
	{r: 'U', n: 3},
	{r: 'D', n: 4},
	{r: 'U', n: 4},
	{r: 'R', n: 5},
	{r: 'D', n: 2},
	{r: 'L', n: 6},
	{r: 'D', n: 4},
	{r: 'U', n: 2},
	{r: 'R', n: 5},
	{r: 'U', n: 5},
	{r: 'D', n: 5},
	{r: 'R', n: 5},
	{r: 'L', n: 3},
	{r: 'R', n: 5},
	{r: 'D', n: 1},
	{r: 'U', n: 3},
	{r: 'L', n: 5},
	{r: 'U', n: 5},
	{r: 'R', n: 5},
	{r: 'D', n: 6},
	{r: 'L', n: 3},
	{r: 'R', n: 1},
	{r: 'U', n: 4},
	{r: 'L', n: 4},
	{r: 'D', n: 5},
	{r: 'L', n: 3},
	{r: 'D', n: 5},
	{r: 'L', n: 1},
	{r: 'R', n: 6},
	{r: 'U', n: 1},
	{r: 'D', n: 1},
	{r: 'U', n: 4},
	{r: 'L', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 2},
	{r: 'R', n: 3},
	{r: 'L', n: 6},
	{r: 'R', n: 4},
	{r: 'L', n: 6},
	{r: 'U', n: 3},
	{r: 'D', n: 5},
	{r: 'L', n: 4},
	{r: 'R', n: 1},
	{r: 'L', n: 6},
	{r: 'R', n: 4},
	{r: 'L', n: 6},
	{r: 'D', n: 4},
	{r: 'U', n: 4},
	{r: 'L', n: 4},
	{r: 'R', n: 3},
	{r: 'L', n: 4},
	{r: 'R', n: 2},
	{r: 'L', n: 2},
	{r: 'R', n: 6},
	{r: 'U', n: 1},
	{r: 'D', n: 6},
	{r: 'L', n: 6},
	{r: 'R', n: 3},
	{r: 'D', n: 6},
	{r: 'U', n: 3},
	{r: 'D', n: 5},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 5},
	{r: 'D', n: 6},
	{r: 'R', n: 6},
	{r: 'D', n: 3},
	{r: 'U', n: 2},
	{r: 'R', n: 6},
	{r: 'U', n: 4},
	{r: 'D', n: 4},
	{r: 'U', n: 2},
	{r: 'R', n: 1},
	{r: 'D', n: 2},
	{r: 'R', n: 4},
	{r: 'U', n: 1},
	{r: 'L', n: 6},
	{r: 'R', n: 5},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'L', n: 6},
	{r: 'D', n: 2},
	{r: 'U', n: 2},
	{r: 'L', n: 2},
	{r: 'U', n: 2},
	{r: 'R', n: 3},
	{r: 'D', n: 5},
	{r: 'U', n: 3},
	{r: 'L', n: 2},
	{r: 'R', n: 7},
	{r: 'L', n: 6},
	{r: 'D', n: 6},
	{r: 'L', n: 4},
	{r: 'R', n: 6},
	{r: 'D', n: 7},
	{r: 'L', n: 3},
	{r: 'R', n: 4},
	{r: 'U', n: 1},
	{r: 'D', n: 5},
	{r: 'R', n: 1},
	{r: 'L', n: 1},
	{r: 'U', n: 2},
	{r: 'R', n: 4},
	{r: 'U', n: 3},
	{r: 'L', n: 6},
	{r: 'R', n: 4},
	{r: 'D', n: 3},
	{r: 'U', n: 7},
	{r: 'D', n: 1},
	{r: 'L', n: 6},
	{r: 'D', n: 6},
	{r: 'R', n: 1},
	{r: 'D', n: 4},
	{r: 'R', n: 3},
	{r: 'D', n: 2},
	{r: 'R', n: 1},
	{r: 'U', n: 3},
	{r: 'L', n: 7},
	{r: 'R', n: 7},
	{r: 'L', n: 2},
	{r: 'U', n: 2},
	{r: 'D', n: 3},
	{r: 'L', n: 5},
	{r: 'D', n: 3},
	{r: 'U', n: 2},
	{r: 'R', n: 7},
	{r: 'L', n: 6},
	{r: 'U', n: 3},
	{r: 'R', n: 1},
	{r: 'U', n: 7},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'L', n: 5},
	{r: 'D', n: 1},
	{r: 'U', n: 6},
	{r: 'L', n: 3},
	{r: 'U', n: 6},
	{r: 'D', n: 7},
	{r: 'U', n: 5},
	{r: 'L', n: 2},
	{r: 'R', n: 3},
	{r: 'L', n: 5},
	{r: 'R', n: 4},
	{r: 'D', n: 3},
	{r: 'U', n: 5},
	{r: 'L', n: 6},
	{r: 'R', n: 1},
	{r: 'U', n: 4},
	{r: 'L', n: 4},
	{r: 'U', n: 5},
	{r: 'L', n: 7},
	{r: 'D', n: 6},
	{r: 'U', n: 4},
	{r: 'R', n: 4},
	{r: 'D', n: 7},
	{r: 'U', n: 5},
	{r: 'D', n: 7},
	{r: 'R', n: 6},
	{r: 'L', n: 5},
	{r: 'U', n: 4},
	{r: 'R', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 6},
	{r: 'D', n: 6},
	{r: 'U', n: 6},
	{r: 'R', n: 7},
	{r: 'L', n: 3},
	{r: 'R', n: 6},
	{r: 'D', n: 1},
	{r: 'R', n: 5},
	{r: 'L', n: 6},
	{r: 'D', n: 5},
	{r: 'L', n: 4},
	{r: 'R', n: 4},
	{r: 'L', n: 6},
	{r: 'R', n: 6},
	{r: 'L', n: 3},
	{r: 'D', n: 2},
	{r: 'L', n: 6},
	{r: 'D', n: 4},
	{r: 'R', n: 5},
	{r: 'L', n: 5},
	{r: 'U', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 6},
	{r: 'R', n: 7},
	{r: 'L', n: 6},
	{r: 'U', n: 6},
	{r: 'D', n: 5},
	{r: 'U', n: 7},
	{r: 'L', n: 4},
	{r: 'U', n: 6},
	{r: 'D', n: 7},
	{r: 'U', n: 3},
	{r: 'D', n: 6},
	{r: 'L', n: 6},
	{r: 'U', n: 7},
	{r: 'D', n: 5},
	{r: 'U', n: 1},
	{r: 'R', n: 2},
	{r: 'L', n: 1},
	{r: 'R', n: 7},
	{r: 'L', n: 6},
	{r: 'D', n: 5},
	{r: 'U', n: 8},
	{r: 'D', n: 3},
	{r: 'U', n: 8},
	{r: 'D', n: 1},
	{r: 'U', n: 7},
	{r: 'D', n: 2},
	{r: 'L', n: 4},
	{r: 'R', n: 2},
	{r: 'U', n: 5},
	{r: 'D', n: 6},
	{r: 'L', n: 1},
	{r: 'R', n: 1},
	{r: 'L', n: 5},
	{r: 'R', n: 7},
	{r: 'D', n: 8},
	{r: 'U', n: 2},
	{r: 'R', n: 7},
	{r: 'U', n: 4},
	{r: 'D', n: 5},
	{r: 'U', n: 8},
	{r: 'D', n: 7},
	{r: 'U', n: 4},
	{r: 'R', n: 3},
	{r: 'D', n: 2},
	{r: 'L', n: 8},
	{r: 'D', n: 4},
	{r: 'L', n: 8},
	{r: 'D', n: 5},
	{r: 'U', n: 7},
	{r: 'R', n: 8},
	{r: 'U', n: 2},
	{r: 'L', n: 1},
	{r: 'U', n: 6},
	{r: 'D', n: 3},
	{r: 'U', n: 6},
	{r: 'L', n: 3},
	{r: 'U', n: 2},
	{r: 'D', n: 8},
	{r: 'U', n: 1},
	{r: 'R', n: 4},
	{r: 'L', n: 3},
	{r: 'R', n: 1},
	{r: 'L', n: 6},
	{r: 'U', n: 6},
	{r: 'L', n: 4},
	{r: 'R', n: 7},
	{r: 'D', n: 8},
	{r: 'U', n: 4},
	{r: 'R', n: 1},
	{r: 'D', n: 7},
	{r: 'L', n: 4},
	{r: 'D', n: 6},
	{r: 'R', n: 4},
	{r: 'L', n: 3},
	{r: 'D', n: 5},
	{r: 'R', n: 7},
	{r: 'L', n: 5},
	{r: 'D', n: 3},
	{r: 'U', n: 8},
	{r: 'L', n: 8},
	{r: 'D', n: 4},
	{r: 'R', n: 1},
	{r: 'U', n: 3},
	{r: 'L', n: 5},
	{r: 'R', n: 4},
	{r: 'L', n: 1},
	{r: 'R', n: 8},
	{r: 'U', n: 7},
	{r: 'D', n: 3},
	{r: 'U', n: 2},
	{r: 'D', n: 7},
	{r: 'L', n: 4},
	{r: 'U', n: 8},
	{r: 'L', n: 7},
	{r: 'U', n: 2},
	{r: 'D', n: 1},
	{r: 'R', n: 7},
	{r: 'L', n: 4},
	{r: 'U', n: 1},
	{r: 'R', n: 2},
	{r: 'D', n: 5},
	{r: 'L', n: 1},
	{r: 'U', n: 5},
	{r: 'L', n: 8},
	{r: 'R', n: 7},
	{r: 'U', n: 1},
	{r: 'D', n: 1},
	{r: 'U', n: 5},
	{r: 'D', n: 7},
	{r: 'L', n: 6},
	{r: 'D', n: 4},
	{r: 'U', n: 7},
	{r: 'D', n: 8},
	{r: 'R', n: 1},
	{r: 'D', n: 3},
	{r: 'U', n: 4},
	{r: 'L', n: 6},
	{r: 'R', n: 7},
	{r: 'L', n: 8},
	{r: 'R', n: 1},
	{r: 'D', n: 7},
	{r: 'L', n: 3},
	{r: 'D', n: 6},
	{r: 'U', n: 5},
	{r: 'L', n: 6},
	{r: 'R', n: 7},
	{r: 'L', n: 9},
	{r: 'D', n: 2},
	{r: 'R', n: 1},
	{r: 'L', n: 4},
	{r: 'D', n: 5},
	{r: 'L', n: 6},
	{r: 'D', n: 8},
	{r: 'U', n: 7},
	{r: 'L', n: 7},
	{r: 'U', n: 5},
	{r: 'D', n: 8},
	{r: 'U', n: 8},
	{r: 'L', n: 6},
	{r: 'U', n: 6},
	{r: 'D', n: 8},
	{r: 'U', n: 9},
	{r: 'D', n: 1},
	{r: 'R', n: 4},
	{r: 'L', n: 9},
	{r: 'D', n: 4},
	{r: 'R', n: 4},
	{r: 'L', n: 3},
	{r: 'U', n: 8},
	{r: 'D', n: 2},
	{r: 'R', n: 3},
	{r: 'U', n: 6},
	{r: 'L', n: 8},
	{r: 'D', n: 1},
	{r: 'R', n: 7},
	{r: 'D', n: 5},
	{r: 'R', n: 9},
	{r: 'D', n: 1},
	{r: 'L', n: 7},
	{r: 'D', n: 2},
	{r: 'U', n: 2},
	{r: 'L', n: 9},
	{r: 'R', n: 5},
	{r: 'L', n: 9},
	{r: 'D', n: 6},
	{r: 'U', n: 6},
	{r: 'L', n: 2},
	{r: 'D', n: 5},
	{r: 'R', n: 3},
	{r: 'L', n: 4},
	{r: 'D', n: 7},
	{r: 'U', n: 2},
	{r: 'R', n: 6},
	{r: 'U', n: 1},
	{r: 'L', n: 8},
	{r: 'R', n: 3},
	{r: 'L', n: 6},
	{r: 'U', n: 5},
	{r: 'R', n: 9},
	{r: 'U', n: 8},
	{r: 'L', n: 2},
	{r: 'D', n: 9},
	{r: 'U', n: 9},
	{r: 'D', n: 7},
	{r: 'L', n: 1},
	{r: 'U', n: 6},
	{r: 'R', n: 4},
	{r: 'L', n: 2},
	{r: 'R', n: 1},
	{r: 'D', n: 1},
	{r: 'R', n: 7},
	{r: 'L', n: 7},
	{r: 'U', n: 6},
	{r: 'D', n: 8},
	{r: 'U', n: 6},
	{r: 'D', n: 5},
	{r: 'L', n: 3},
	{r: 'D', n: 8},
	{r: 'U', n: 6},
	{r: 'L', n: 2},
	{r: 'U', n: 8},
	{r: 'R', n: 8},
	{r: 'D', n: 7},
	{r: 'U', n: 9},
	{r: 'R', n: 6},
	{r: 'U', n: 1},
	{r: 'R', n: 2},
	{r: 'D', n: 5},
	{r: 'R', n: 4},
	{r: 'L', n: 2},
	{r: 'U', n: 3},
	{r: 'D', n: 3},
	{r: 'L', n: 5},
	{r: 'U', n: 1},
	{r: 'L', n: 5},
	{r: 'D', n: 8},
	{r: 'U', n: 9},
	{r: 'L', n: 8},
	{r: 'U', n: 9},
	{r: 'L', n: 9},
	{r: 'U', n: 5},
	{r: 'R', n: 7},
	{r: 'U', n: 7},
	{r: 'D', n: 4},
	{r: 'R', n: 2},
	{r: 'L', n: 8},
	{r: 'U', n: 9},
	{r: 'D', n: 6},
	{r: 'U', n: 5},
	{r: 'L', n: 9},
	{r: 'U', n: 3},
	{r: 'L', n: 5},
	{r: 'U', n: 9},
	{r: 'L', n: 7},
	{r: 'D', n: 5},
	{r: 'L', n: 8},
	{r: 'R', n: 9},
	{r: 'L', n: 10},
	{r: 'D', n: 1},
	{r: 'L', n: 10},
	{r: 'R', n: 4},
	{r: 'L', n: 7},
	{r: 'U', n: 5},
	{r: 'L', n: 10},
	{r: 'U', n: 2},
	{r: 'R', n: 10},
	{r: 'L', n: 3},
	{r: 'U', n: 10},
	{r: 'R', n: 2},
	{r: 'U', n: 6},
	{r: 'L', n: 2},
	{r: 'R', n: 8},
	{r: 'L', n: 8},
	{r: 'R', n: 4},
	{r: 'U', n: 1},
	{r: 'R', n: 6},
	{r: 'L', n: 2},
	{r: 'D', n: 6},
	{r: 'L', n: 10},
	{r: 'D', n: 9},
	{r: 'L', n: 2},
	{r: 'U', n: 5},
	{r: 'R', n: 1},
	{r: 'D', n: 9},
	{r: 'R', n: 9},
	{r: 'U', n: 3},
	{r: 'D', n: 6},
	{r: 'L', n: 10},
	{r: 'U', n: 4},
	{r: 'R', n: 6},
	{r: 'U', n: 9},
	{r: 'R', n: 6},
	{r: 'U', n: 6},
	{r: 'L', n: 4},
	{r: 'R', n: 8},
	{r: 'D', n: 3},
	{r: 'U', n: 8},
	{r: 'L', n: 3},
	{r: 'U', n: 6},
	{r: 'R', n: 2},
	{r: 'L', n: 2},
	{r: 'D', n: 3},
	{r: 'U', n: 3},
	{r: 'R', n: 3},
	{r: 'D', n: 10},
	{r: 'R', n: 9},
	{r: 'D', n: 1},
	{r: 'R', n: 6},
	{r: 'L', n: 4},
	{r: 'U', n: 4},
	{r: 'D', n: 1},
	{r: 'R', n: 2},
	{r: 'U', n: 6},
	{r: 'R', n: 3},
	{r: 'U', n: 5},
	{r: 'R', n: 8},
	{r: 'L', n: 5},
	{r: 'U', n: 2},
	{r: 'R', n: 5},
	{r: 'D', n: 10},
	{r: 'L', n: 7},
	{r: 'R', n: 4},
	{r: 'D', n: 9},
	{r: 'U', n: 6},
	{r: 'L', n: 7},
	{r: 'R', n: 10},
	{r: 'D', n: 3},
	{r: 'U', n: 5},
	{r: 'R', n: 1},
	{r: 'U', n: 2},
	{r: 'R', n: 9},
	{r: 'U', n: 10},
	{r: 'R', n: 8},
	{r: 'L', n: 8},
	{r: 'U', n: 4},
	{r: 'R', n: 9},
	{r: 'L', n: 6},
	{r: 'U', n: 7},
	{r: 'L', n: 4},
	{r: 'U', n: 1},
	{r: 'L', n: 6},
	{r: 'U', n: 4},
	{r: 'L', n: 7},
	{r: 'R', n: 4},
	{r: 'U', n: 9},
	{r: 'D', n: 9},
	{r: 'R', n: 9},
	{r: 'L', n: 4},
	{r: 'R', n: 3},
	{r: 'D', n: 7},
	{r: 'L', n: 3},
	{r: 'U', n: 6},
	{r: 'D', n: 2},
	{r: 'U', n: 8},
	{r: 'R', n: 8},
	{r: 'L', n: 6},
	{r: 'D', n: 5},
	{r: 'L', n: 6},
	{r: 'R', n: 1},
	{r: 'L', n: 3},
	{r: 'U', n: 10},
	{r: 'L', n: 2},
	{r: 'D', n: 7},
	{r: 'R', n: 7},
	{r: 'L', n: 8},
	{r: 'U', n: 4},
	{r: 'R', n: 1},
	{r: 'D', n: 5},
	{r: 'L', n: 10},
	{r: 'R', n: 6},
	{r: 'U', n: 1},
	{r: 'R', n: 7},
	{r: 'L', n: 6},
	{r: 'R', n: 11},
	{r: 'U', n: 11},
	{r: 'R', n: 3},
	{r: 'L', n: 5},
	{r: 'U', n: 7},
	{r: 'D', n: 7},
	{r: 'L', n: 9},
	{r: 'U', n: 4},
	{r: 'L', n: 5},
	{r: 'D', n: 11},
	{r: 'L', n: 4},
	{r: 'R', n: 1},
	{r: 'L', n: 10},
	{r: 'U', n: 9},
	{r: 'L', n: 8},
	{r: 'U', n: 10},
	{r: 'R', n: 7},
	{r: 'U', n: 3},
	{r: 'R', n: 3},
	{r: 'D', n: 1},
	{r: 'U', n: 4},
	{r: 'L', n: 5},
	{r: 'D', n: 6},
	{r: 'R', n: 8},
	{r: 'D', n: 10},
	{r: 'R', n: 4},
	{r: 'L', n: 9},
	{r: 'U', n: 6},
	{r: 'R', n: 10},
	{r: 'D', n: 2},
	{r: 'U', n: 9},
	{r: 'R', n: 10},
	{r: 'U', n: 5},
	{r: 'R', n: 4},
	{r: 'U', n: 3},
	{r: 'R', n: 4},
	{r: 'L', n: 1},
	{r: 'U', n: 3},
	{r: 'R', n: 7},
	{r: 'L', n: 3},
	{r: 'D', n: 3},
	{r: 'U', n: 6},
	{r: 'L', n: 8},
	{r: 'U', n: 7},
	{r: 'L', n: 7},
	{r: 'R', n: 4},
	{r: 'L', n: 1},
	{r: 'D', n: 5},
	{r: 'U', n: 1},
	{r: 'R', n: 4},
	{r: 'D', n: 1},
	{r: 'R', n: 1},
	{r: 'D', n: 6},
	{r: 'L', n: 1},
	{r: 'D', n: 1},
	{r: 'L', n: 3},
	{r: 'D', n: 7},
	{r: 'R', n: 1},
	{r: 'U', n: 7},
	{r: 'L', n: 11},
	{r: 'U', n: 7},
	{r: 'L', n: 9},
	{r: 'D', n: 7},
	{r: 'R', n: 5},
	{r: 'L', n: 11},
	{r: 'D', n: 3},
	{r: 'R', n: 5},
	{r: 'U', n: 3},
	{r: 'L', n: 9},
	{r: 'R', n: 4},
	{r: 'U', n: 2},
	{r: 'L', n: 9},
	{r: 'U', n: 1},
	{r: 'R', n: 4},
	{r: 'U', n: 8},
	{r: 'D', n: 1},
	{r: 'R', n: 1},
	{r: 'D', n: 5},
	{r: 'R', n: 5},
	{r: 'U', n: 11},
	{r: 'L', n: 10},
	{r: 'U', n: 5},
	{r: 'D', n: 4},
	{r: 'R', n: 7},
	{r: 'U', n: 4},
	{r: 'D', n: 10},
	{r: 'L', n: 4},
	{r: 'D', n: 6},
	{r: 'U', n: 6},
	{r: 'L', n: 8},
	{r: 'U', n: 8},
	{r: 'D', n: 1},
	{r: 'L', n: 4},
	{r: 'R', n: 1},
	{r: 'L', n: 10},
	{r: 'D', n: 5},
	{r: 'R', n: 6},
	{r: 'U', n: 4},
	{r: 'L', n: 4},
	{r: 'U', n: 9},
	{r: 'L', n: 7},
	{r: 'R', n: 6},
	{r: 'U', n: 8},
	{r: 'L', n: 4},
	{r: 'D', n: 5},
	{r: 'R', n: 4},
	{r: 'L', n: 1},
	{r: 'U', n: 3},
	{r: 'D', n: 7},
	{r: 'R', n: 3},
	{r: 'U', n: 3},
	{r: 'L', n: 2},
	{r: 'D', n: 4},
	{r: 'R', n: 2},
	{r: 'L', n: 6},
	{r: 'R', n: 11},
	{r: 'D', n: 2},
	{r: 'L', n: 6},
	{r: 'U', n: 2},
	{r: 'R', n: 12},
	{r: 'L', n: 6},
	{r: 'D', n: 12},
	{r: 'R', n: 11},
	{r: 'L', n: 9},
	{r: 'D', n: 9},
	{r: 'R', n: 5},
	{r: 'U', n: 2},
	{r: 'L', n: 5},
	{r: 'D', n: 10},
	{r: 'R', n: 4},
	{r: 'U', n: 4},
	{r: 'R', n: 8},
	{r: 'U', n: 2},
	{r: 'R', n: 6},
	{r: 'U', n: 3},
	{r: 'D', n: 7},
	{r: 'U', n: 11},
	{r: 'L', n: 8},
	{r: 'D', n: 8},
	{r: 'U', n: 11},
	{r: 'D', n: 2},
	{r: 'U', n: 3},
	{r: 'R', n: 7},
	{r: 'D', n: 9},
	{r: 'R', n: 3},
	{r: 'L', n: 11},
	{r: 'U', n: 3},
	{r: 'L', n: 7},
	{r: 'D', n: 6},
	{r: 'L', n: 3},
	{r: 'D', n: 2},
	{r: 'R', n: 1},
	{r: 'U', n: 9},
	{r: 'L', n: 6},
	{r: 'D', n: 6},
	{r: 'L', n: 5},
	{r: 'R', n: 8},
	{r: 'L', n: 4},
	{r: 'R', n: 9},
	{r: 'U', n: 5},
	{r: 'D', n: 11},
	{r: 'R', n: 7},
	{r: 'D', n: 11},
	{r: 'L', n: 6},
	{r: 'R', n: 2},
	{r: 'L', n: 5},
	{r: 'U', n: 6},
	{r: 'L', n: 5},
	{r: 'D', n: 4},
	{r: 'R', n: 3},
	{r: 'D', n: 8},
	{r: 'R', n: 8},
	{r: 'D', n: 2},
	{r: 'U', n: 1},
	{r: 'R', n: 6},
	{r: 'U', n: 5},
	{r: 'L', n: 7},
	{r: 'D', n: 11},
	{r: 'U', n: 4},
	{r: 'L', n: 1},
	{r: 'R', n: 8},
	{r: 'D', n: 6},
	{r: 'R', n: 6},
	{r: 'U', n: 11},
	{r: 'L', n: 8},
	{r: 'D', n: 2},
	{r: 'L', n: 6},
	{r: 'R', n: 10},
	{r: 'L', n: 4},
	{r: 'D', n: 4},
	{r: 'U', n: 3},
	{r: 'D', n: 12},
	{r: 'U', n: 10},
	{r: 'L', n: 3},
	{r: 'U', n: 3},
	{r: 'D', n: 7},
	{r: 'R', n: 9},
	{r: 'L', n: 2},
	{r: 'U', n: 11},
	{r: 'L', n: 11},
	{r: 'U', n: 4},
	{r: 'R', n: 2},
	{r: 'L', n: 5},
	{r: 'D', n: 11},
	{r: 'U', n: 5},
	{r: 'L', n: 10},
	{r: 'U', n: 4},
	{r: 'R', n: 1},
	{r: 'U', n: 11},
	{r: 'R', n: 12},
	{r: 'U', n: 2},
	{r: 'L', n: 7},
	{r: 'U', n: 12},
	{r: 'L', n: 1},
	{r: 'D', n: 9},
	{r: 'U', n: 4},
	{r: 'L', n: 4},
	{r: 'D', n: 12},
	{r: 'L', n: 5},
	{r: 'D', n: 1},
	{r: 'R', n: 9},
	{r: 'U', n: 11},
	{r: 'L', n: 6},
	{r: 'D', n: 8},
	{r: 'L', n: 11},
	{r: 'U', n: 9},
	{r: 'D', n: 6},
	{r: 'L', n: 6},
	{r: 'R', n: 6},
	{r: 'U', n: 5},
	{r: 'D', n: 10},
	{r: 'R', n: 4},
	{r: 'D', n: 4},
	{r: 'U', n: 10},
	{r: 'D', n: 2},
	{r: 'L', n: 12},
	{r: 'D', n: 3},
	{r: 'L', n: 7},
	{r: 'D', n: 13},
	{r: 'U', n: 11},
	{r: 'D', n: 4},
	{r: 'L', n: 6},
	{r: 'U', n: 4},
	{r: 'D', n: 10},
	{r: 'L', n: 2},
	{r: 'R', n: 4},
	{r: 'D', n: 7},
	{r: 'U', n: 7},
	{r: 'R', n: 9},
	{r: 'L', n: 8},
	{r: 'R', n: 3},
	{r: 'U', n: 8},
	{r: 'D', n: 10},
	{r: 'L', n: 2},
	{r: 'D', n: 2},
	{r: 'L', n: 2},
	{r: 'D', n: 11},
	{r: 'L', n: 8},
	{r: 'U', n: 12},
	{r: 'L', n: 1},
	{r: 'U', n: 9},
	{r: 'L', n: 6},
	{r: 'R', n: 7},
	{r: 'L', n: 6},
	{r: 'D', n: 9},
	{r: 'U', n: 10},
	{r: 'R', n: 2},
	{r: 'L', n: 3},
	{r: 'U', n: 5},
	{r: 'L', n: 3},
	{r: 'D', n: 9},
	{r: 'R', n: 8},
	{r: 'D', n: 1},
	{r: 'L', n: 11},
	{r: 'D', n: 1},
	{r: 'U', n: 10},
	{r: 'D', n: 5},
	{r: 'L', n: 10},
	{r: 'U', n: 1},
	{r: 'L', n: 2},
	{r: 'D', n: 11},
	{r: 'L', n: 8},
	{r: 'R', n: 6},
	{r: 'D', n: 6},
	{r: 'U', n: 13},
	{r: 'L', n: 1},
	{r: 'D', n: 3},
	{r: 'L', n: 11},
	{r: 'D', n: 13},
	{r: 'R', n: 9},
	{r: 'D', n: 6},
	{r: 'L', n: 7},
	{r: 'U', n: 7},
	{r: 'R', n: 8},
	{r: 'L', n: 3},
	{r: 'D', n: 3},
	{r: 'L', n: 9},
	{r: 'D', n: 5},
	{r: 'U', n: 12},
	{r: 'D', n: 6},
	{r: 'L', n: 8},
	{r: 'U', n: 8},
	{r: 'L', n: 12},
	{r: 'U', n: 10},
	{r: 'R', n: 2},
	{r: 'L', n: 8},
	{r: 'D', n: 12},
	{r: 'R', n: 4},
	{r: 'L', n: 3},
	{r: 'R', n: 7},
	{r: 'U', n: 12},
	{r: 'R', n: 2},
	{r: 'L', n: 2},
	{r: 'R', n: 10},
	{r: 'L', n: 12},
	{r: 'R', n: 1},
	{r: 'L', n: 1},
	{r: 'R', n: 3},
	{r: 'L', n: 2},
	{r: 'R', n: 2},
	{r: 'D', n: 3},
	{r: 'R', n: 8},
	{r: 'L', n: 10},
	{r: 'R', n: 8},
	{r: 'L', n: 11},
	{r: 'U', n: 3},
	{r: 'D', n: 12},
	{r: 'R', n: 8},
	{r: 'U', n: 12},
	{r: 'D', n: 7},
	{r: 'L', n: 6},
	{r: 'D', n: 14},
	{r: 'L', n: 11},
	{r: 'U', n: 12},
	{r: 'L', n: 10},
	{r: 'R', n: 4},
	{r: 'U', n: 11},
	{r: 'R', n: 5},
	{r: 'U', n: 13},
	{r: 'R', n: 3},
	{r: 'U', n: 9},
	{r: 'R', n: 2},
	{r: 'L', n: 2},
	{r: 'D', n: 3},
	{r: 'L', n: 3},
	{r: 'R', n: 1},
	{r: 'D', n: 11},
	{r: 'L', n: 1},
	{r: 'D', n: 6},
	{r: 'R', n: 9},
	{r: 'L', n: 10},
	{r: 'R', n: 8},
	{r: 'D', n: 5},
	{r: 'U', n: 6},
	{r: 'D', n: 12},
	{r: 'U', n: 12},
	{r: 'L', n: 10},
	{r: 'D', n: 8},
	{r: 'R', n: 4},
	{r: 'D', n: 3},
	{r: 'L', n: 10},
	{r: 'R', n: 14},
	{r: 'L', n: 5},
	{r: 'D', n: 5},
	{r: 'L', n: 1},
	{r: 'U', n: 3},
	{r: 'R', n: 10},
	{r: 'U', n: 2},
	{r: 'D', n: 2},
	{r: 'L', n: 12},
	{r: 'R', n: 7},
	{r: 'U', n: 1},
	{r: 'D', n: 13},
	{r: 'R', n: 11},
	{r: 'L', n: 13},
	{r: 'U', n: 8},
	{r: 'L', n: 10},
	{r: 'R', n: 2},
	{r: 'U', n: 8},
	{r: 'D', n: 10},
	{r: 'R', n: 14},
	{r: 'U', n: 2},
	{r: 'L', n: 14},
	{r: 'D', n: 12},
	{r: 'R', n: 8},
	{r: 'D', n: 8},
	{r: 'U', n: 5},
	{r: 'R', n: 7},
	{r: 'L', n: 14},
	{r: 'R', n: 6},
	{r: 'U', n: 13},
	{r: 'D', n: 7},
	{r: 'U', n: 2},
	{r: 'L', n: 7},
	{r: 'R', n: 4},
	{r: 'L', n: 8},
	{r: 'U', n: 7},
	{r: 'R', n: 3},
	{r: 'D', n: 1},
	{r: 'R', n: 3},
	{r: 'D', n: 14},
	{r: 'U', n: 2},
	{r: 'R', n: 1},
	{r: 'D', n: 7},
	{r: 'L', n: 4},
	{r: 'D', n: 14},
	{r: 'U', n: 8},
	{r: 'L', n: 9},
	{r: 'U', n: 11},
	{r: 'L', n: 5},
	{r: 'R', n: 2},
	{r: 'D', n: 4},
	{r: 'L', n: 3},
	{r: 'U', n: 9},
	{r: 'R', n: 4},
	{r: 'U', n: 4},
	{r: 'R', n: 1},
	{r: 'U', n: 3},
	{r: 'L', n: 6},
	{r: 'D', n: 14},
	{r: 'U', n: 12},
	{r: 'R', n: 11},
	{r: 'D', n: 6},
	{r: 'U', n: 8},
	{r: 'D', n: 14},
	{r: 'R', n: 14},
	{r: 'L', n: 13},
	{r: 'R', n: 11},
	{r: 'D', n: 2},
	{r: 'U', n: 6},
	{r: 'L', n: 7},
	{r: 'D', n: 9},
	{r: 'R', n: 7},
	{r: 'D', n: 14},
	{r: 'L', n: 2},
	{r: 'R', n: 7},
	{r: 'L', n: 13},
	{r: 'R', n: 3},
	{r: 'D', n: 8},
	{r: 'U', n: 5},
	{r: 'L', n: 14},
	{r: 'D', n: 7},
	{r: 'U', n: 5},
	{r: 'D', n: 12},
	{r: 'R', n: 11},
	{r: 'U', n: 1},
	{r: 'D', n: 10},
	{r: 'L', n: 7},
	{r: 'R', n: 7},
	{r: 'D', n: 2},
	{r: 'L', n: 4},
	{r: 'U', n: 13},
	{r: 'R', n: 10},
	{r: 'D', n: 1},
	{r: 'L', n: 9},
	{r: 'D', n: 14},
	{r: 'U', n: 6},
	{r: 'D', n: 5},
	{r: 'R', n: 8},
	{r: 'U', n: 7},
	{r: 'L', n: 9},
	{r: 'R', n: 15},
	{r: 'L', n: 5},
	{r: 'U', n: 15},
	{r: 'R', n: 9},
	{r: 'U', n: 3},
	{r: 'L', n: 12},
	{r: 'R', n: 2},
	{r: 'D', n: 8},
	{r: 'U', n: 8},
	{r: 'D', n: 11},
	{r: 'R', n: 14},
	{r: 'D', n: 14},
	{r: 'U', n: 5},
	{r: 'R', n: 13},
	{r: 'D', n: 13},
	{r: 'L', n: 7},
	{r: 'U', n: 14},
	{r: 'D', n: 8},
	{r: 'R', n: 12},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 6},
	{r: 'R', n: 4},
	{r: 'U', n: 2},
	{r: 'R', n: 14},
	{r: 'L', n: 5},
	{r: 'R', n: 2},
	{r: 'D', n: 11},
	{r: 'U', n: 1},
	{r: 'L', n: 14},
	{r: 'D', n: 1},
	{r: 'R', n: 5},
	{r: 'D', n: 12},
	{r: 'U', n: 4},
	{r: 'D', n: 8},
	{r: 'U', n: 6},
	{r: 'D', n: 12},
	{r: 'L', n: 5},
	{r: 'U', n: 4},
	{r: 'R', n: 15},
	{r: 'D', n: 4},
	{r: 'R', n: 4},
	{r: 'U', n: 4},
	{r: 'R', n: 3},
	{r: 'L', n: 1},
	{r: 'U', n: 8},
	{r: 'R', n: 4},
	{r: 'D', n: 7},
	{r: 'U', n: 13},
	{r: 'R', n: 8},
	{r: 'L', n: 1},
	{r: 'D', n: 3},
	{r: 'L', n: 3},
	{r: 'U', n: 11},
	{r: 'L', n: 12},
	{r: 'U', n: 10},
	{r: 'L', n: 12},
	{r: 'U', n: 13},
	{r: 'D', n: 2},
	{r: 'L', n: 7},
	{r: 'U', n: 1},
	{r: 'R', n: 1},
	{r: 'U', n: 15},
	{r: 'D', n: 9},
	{r: 'L', n: 1},
	{r: 'R', n: 6},
	{r: 'L', n: 10},
	{r: 'D', n: 7},
	{r: 'U', n: 5},
	{r: 'L', n: 8},
	{r: 'U', n: 9},
	{r: 'R', n: 5},
	{r: 'U', n: 15},
	{r: 'D', n: 7},
	{r: 'L', n: 9},
	{r: 'R', n: 15},
	{r: 'L', n: 15},
	{r: 'R', n: 6},
	{r: 'D', n: 2},
	{r: 'R', n: 2},
	{r: 'L', n: 5},
	{r: 'U', n: 5},
	{r: 'L', n: 11},
	{r: 'D', n: 8},
	{r: 'R', n: 15},
	{r: 'L', n: 5},
	{r: 'D', n: 10},
	{r: 'L', n: 13},
	{r: 'D', n: 13},
	{r: 'R', n: 13},
	{r: 'D', n: 14},
	{r: 'L', n: 11},
	{r: 'R', n: 2},
	{r: 'D', n: 6},
	{r: 'U', n: 16},
	{r: 'D', n: 13},
	{r: 'L', n: 8},
	{r: 'D', n: 2},
	{r: 'L', n: 5},
	{r: 'U', n: 8},
	{r: 'L', n: 4},
	{r: 'R', n: 9},
	{r: 'D', n: 4},
	{r: 'U', n: 5},
	{r: 'D', n: 9},
	{r: 'U', n: 7},
	{r: 'L', n: 11},
	{r: 'D', n: 8},
	{r: 'U', n: 7},
	{r: 'D', n: 10},
	{r: 'U', n: 11},
	{r: 'R', n: 5},
	{r: 'U', n: 14},
	{r: 'L', n: 12},
	{r: 'R', n: 14},
	{r: 'D', n: 3},
	{r: 'R', n: 3},
	{r: 'D', n: 1},
	{r: 'U', n: 8},
	{r: 'D', n: 2},
	{r: 'R', n: 14},
	{r: 'D', n: 8},
	{r: 'R', n: 6},
	{r: 'L', n: 9},
	{r: 'D', n: 6},
	{r: 'R', n: 4},
	{r: 'L', n: 5},
	{r: 'D', n: 14},
	{r: 'R', n: 11},
	{r: 'L', n: 4},
	{r: 'R', n: 12},
	{r: 'D', n: 15},
	{r: 'R', n: 16},
	{r: 'U', n: 13},
	{r: 'R', n: 7},
	{r: 'U', n: 15},
	{r: 'D', n: 8},
	{r: 'U', n: 11},
	{r: 'L', n: 11},
	{r: 'D', n: 14},
	{r: 'U', n: 14},
	{r: 'D', n: 13},
	{r: 'U', n: 14},
	{r: 'L', n: 14},
	{r: 'D', n: 11},
	{r: 'U', n: 9},
	{r: 'R', n: 5},
	{r: 'L', n: 15},
	{r: 'R', n: 5},
	{r: 'D', n: 6},
	{r: 'U', n: 10},
	{r: 'R', n: 10},
	{r: 'D', n: 16},
	{r: 'R', n: 13},
	{r: 'D', n: 9},
	{r: 'U', n: 10},
	{r: 'R', n: 5},
	{r: 'L', n: 6},
	{r: 'D', n: 15},
	{r: 'U', n: 10},
	{r: 'L', n: 15},
	{r: 'R', n: 7},
	{r: 'L', n: 13},
	{r: 'D', n: 6},
	{r: 'U', n: 7},
	{r: 'L', n: 11},
	{r: 'U', n: 4},
	{r: 'R', n: 4},
	{r: 'L', n: 11},
	{r: 'R', n: 12},
	{r: 'D', n: 6},
	{r: 'L', n: 8},
	{r: 'D', n: 8},
	{r: 'L', n: 12},
	{r: 'D', n: 1},
	{r: 'R', n: 15},
	{r: 'D', n: 12},
	{r: 'L', n: 8},
	{r: 'R', n: 11},
	{r: 'L', n: 10},
	{r: 'D', n: 9},
	{r: 'U', n: 6},
	{r: 'L', n: 2},
	{r: 'R', n: 9},
	{r: 'L', n: 7},
	{r: 'R', n: 4},
	{r: 'U', n: 6},
	{r: 'D', n: 11},
	{r: 'L', n: 6},
	{r: 'D', n: 11},
	{r: 'L', n: 14},
	{r: 'R', n: 10},
	{r: 'U', n: 4},
	{r: 'R', n: 12},
	{r: 'U', n: 10},
	{r: 'D', n: 7},
	{r: 'R', n: 7},
	{r: 'L', n: 7},
	{r: 'U', n: 1},
	{r: 'D', n: 3},
	{r: 'L', n: 1},
	{r: 'U', n: 9},
	{r: 'D', n: 4},
	{r: 'U', n: 15},
	{r: 'L', n: 11},
	{r: 'U', n: 11},
	{r: 'D', n: 11},
	{r: 'L', n: 9},
	{r: 'D', n: 13},
	{r: 'L', n: 15},
	{r: 'D', n: 11},
	{r: 'R', n: 16},
	{r: 'U', n: 11},
	{r: 'D', n: 6},
	{r: 'L', n: 10},
	{r: 'R', n: 10},
	{r: 'D', n: 4},
	{r: 'R', n: 8},
	{r: 'U', n: 17},
	{r: 'D', n: 12},
	{r: 'L', n: 16},
	{r: 'U', n: 14},
	{r: 'R', n: 15},
	{r: 'U', n: 15},
	{r: 'D', n: 6},
	{r: 'R', n: 7},
	{r: 'U', n: 5},
	{r: 'L', n: 9},
	{r: 'U', n: 6},
	{r: 'D', n: 14},
	{r: 'U', n: 4},
	{r: 'L', n: 8},
	{r: 'D', n: 6},
	{r: 'R', n: 16},
	{r: 'U', n: 4},
	{r: 'R', n: 7},
	{r: 'D', n: 3},
	{r: 'U', n: 3},
	{r: 'D', n: 15},
	{r: 'L', n: 1},
	{r: 'U', n: 1},
	{r: 'D', n: 1},
	{r: 'R', n: 1},
	{r: 'U', n: 7},
	{r: 'D', n: 12},
	{r: 'R', n: 1},
	{r: 'D', n: 2},
	{r: 'U', n: 5},
	{r: 'L', n: 13},
	{r: 'U', n: 7},
	{r: 'R', n: 10},
	{r: 'U', n: 15},
	{r: 'D', n: 17},
	{r: 'L', n: 16},
	{r: 'D', n: 8},
	{r: 'L', n: 15},
	{r: 'U', n: 2},
	{r: 'R', n: 2},
	{r: 'D', n: 14},
	{r: 'L', n: 4},
	{r: 'U', n: 7},
	{r: 'D', n: 2},
	{r: 'L', n: 15},
	{r: 'R', n: 5},
	{r: 'U', n: 11},
	{r: 'R', n: 2},
	{r: 'U', n: 8},
	{r: 'D', n: 12},
	{r: 'R', n: 13},
	{r: 'L', n: 7},
	{r: 'D', n: 4},
	{r: 'R', n: 6},
	{r: 'U', n: 2},
	{r: 'D', n: 9},
	{r: 'U', n: 6},
	{r: 'L', n: 10},
	{r: 'R', n: 1},
	{r: 'L', n: 15},
	{r: 'R', n: 12},
	{r: 'L', n: 2},
	{r: 'D', n: 16},
	{r: 'L', n: 2},
	{r: 'R', n: 7},
	{r: 'U', n: 15},
	{r: 'L', n: 3},
	{r: 'D', n: 12},
	{r: 'U', n: 10},
	{r: 'R', n: 14},
	{r: 'L', n: 7},
	{r: 'D', n: 9},
	{r: 'U', n: 13},
	{r: 'R', n: 8},
	{r: 'U', n: 10},
	{r: 'R', n: 7},
	{r: 'L', n: 12},
	{r: 'R', n: 15},
	{r: 'U', n: 10},
	{r: 'D', n: 5},
	{r: 'L', n: 11},
	{r: 'R', n: 9},
	{r: 'D', n: 4},
	{r: 'L', n: 9},
	{r: 'D', n: 8},
	{r: 'R', n: 13},
	{r: 'L', n: 13},
	{r: 'D', n: 7},
	{r: 'L', n: 3},
	{r: 'D', n: 3},
	{r: 'U', n: 14},
	{r: 'D', n: 14},
	{r: 'U', n: 11},
	{r: 'D', n: 17},
	{r: 'R', n: 1},
	{r: 'D', n: 10},
	{r: 'L', n: 12},
	{r: 'R', n: 1},
	{r: 'U', n: 4},
	{r: 'D', n: 16},
	{r: 'U', n: 18},
	{r: 'L', n: 9},
	{r: 'U', n: 2},
	{r: 'R', n: 13},
	{r: 'L', n: 7},
	{r: 'U', n: 7},
	{r: 'R', n: 10},
	{r: 'L', n: 17},
	{r: 'R', n: 7},
	{r: 'U', n: 4},
	{r: 'L', n: 8},
	{r: 'U', n: 1},
	{r: 'L', n: 13},
	{r: 'U', n: 12},
	{r: 'D', n: 7},
	{r: 'R', n: 10},
	{r: 'L', n: 18},
	{r: 'U', n: 12},
	{r: 'D', n: 2},
	{r: 'U', n: 7},
	{r: 'D', n: 12},
	{r: 'R', n: 13},
	{r: 'D', n: 9},
	{r: 'R', n: 13},
	{r: 'U', n: 1},
	{r: 'L', n: 9},
	{r: 'R', n: 3},
	{r: 'L', n: 3},
	{r: 'D', n: 15},
	{r: 'R', n: 11},
	{r: 'L', n: 2},
	{r: 'R', n: 18},
	{r: 'L', n: 7},
	{r: 'R', n: 10},
	{r: 'L', n: 10},
	{r: 'D', n: 8},
	{r: 'L', n: 4},
	{r: 'R', n: 6},
	{r: 'U', n: 12},
	{r: 'L', n: 8},
	{r: 'U', n: 16},
	{r: 'D', n: 18},
	{r: 'L', n: 9},
	{r: 'U', n: 15},
	{r: 'R', n: 13},
	{r: 'D', n: 18},
	{r: 'U', n: 6},
	{r: 'R', n: 7},
	{r: 'U', n: 7},
	{r: 'L', n: 16},
	{r: 'U', n: 18},
	{r: 'R', n: 10},
	{r: 'L', n: 4},
	{r: 'R', n: 2},
	{r: 'U', n: 2},
	{r: 'L', n: 5},
	{r: 'R', n: 8},
	{r: 'U', n: 14},
	{r: 'L', n: 5},
	{r: 'U', n: 2},
	{r: 'D', n: 17},
	{r: 'R', n: 7},
	{r: 'L', n: 13},
	{r: 'D', n: 15},
	{r: 'R', n: 2},
	{r: 'D', n: 2},
	{r: 'R', n: 6},
	{r: 'L', n: 10},
	{r: 'R', n: 4},
	{r: 'U', n: 3},
	{r: 'R', n: 18},
	{r: 'D', n: 16},
	{r: 'L', n: 13},
	{r: 'U', n: 8},
	{r: 'D', n: 5},
	{r: 'L', n: 9},
	{r: 'R', n: 4},
	{r: 'D', n: 7},
	{r: 'R', n: 9},
	{r: 'L', n: 5},
	{r: 'D', n: 13},
	{r: 'L', n: 9},
	{r: 'D', n: 13},
	{r: 'U', n: 15},
	{r: 'L', n: 3},
	{r: 'D', n: 15},
	{r: 'U', n: 16},
	{r: 'R', n: 2},
	{r: 'U', n: 18},
	{r: 'R', n: 18},
	{r: 'L', n: 6},
	{r: 'D', n: 12},
	{r: 'R', n: 10},
	{r: 'U', n: 10},
	{r: 'L', n: 3},
	{r: 'R', n: 3},
	{r: 'D', n: 7},
	{r: 'L', n: 11},
	{r: 'R', n: 1},
	{r: 'D', n: 7},
	{r: 'L', n: 8},
	{r: 'R', n: 15},
	{r: 'L', n: 3},
	{r: 'U', n: 12},
	{r: 'R', n: 7},
	{r: 'D', n: 19},
	{r: 'R', n: 13},
	{r: 'D', n: 11},
	{r: 'U', n: 15},
	{r: 'D', n: 3},
	{r: 'U', n: 6},
	{r: 'D', n: 15},
	{r: 'R', n: 8},
	{r: 'U', n: 5},
	{r: 'L', n: 2},
	{r: 'D', n: 12},
	{r: 'U', n: 3},
	{r: 'D', n: 3},
	{r: 'L', n: 3},
	{r: 'D', n: 9},
	{r: 'L', n: 8},
	{r: 'R', n: 19},
	{r: 'L', n: 12},
	{r: 'U', n: 13},
	{r: 'L', n: 1},
	{r: 'D', n: 14},
	{r: 'R', n: 9},
	{r: 'L', n: 15},
	{r: 'R', n: 15},
	{r: 'L', n: 14},
	{r: 'D', n: 2},
	{r: 'U', n: 3},
	{r: 'L', n: 17},
	{r: 'U', n: 15},
	{r: 'D', n: 13},
	{r: 'R', n: 4},
	{r: 'D', n: 18},
	{r: 'L', n: 4},
	{r: 'R', n: 2},
	{r: 'D', n: 16},
	{r: 'R', n: 9},
	{r: 'L', n: 3},
	{r: 'U', n: 14},
	{r: 'D', n: 1},
	{r: 'R', n: 19},
	{r: 'U', n: 11},
	{r: 'D', n: 6},
	{r: 'U', n: 12},
	{r: 'R', n: 5},
	{r: 'U', n: 9},
	{r: 'L', n: 8},
	{r: 'D', n: 14},
	{r: 'R', n: 5},
	{r: 'U', n: 16},
	{r: 'L', n: 12},
	{r: 'U', n: 17},
	{r: 'R', n: 17},
	{r: 'U', n: 19},
	{r: 'L', n: 1},
	{r: 'D', n: 9},
	{r: 'U', n: 10},
	{r: 'L', n: 4},
	{r: 'R', n: 13},
	{r: 'U', n: 15},
	{r: 'D', n: 7},
	{r: 'R', n: 14},
	{r: 'U', n: 15},
	{r: 'R', n: 7},
	{r: 'D', n: 2},
	{r: 'R', n: 7},
	{r: 'L', n: 10},
	{r: 'D', n: 17},
	{r: 'U', n: 5},
	{r: 'D', n: 1},
	{r: 'L', n: 18},
	{r: 'D', n: 9},
	{r: 'R', n: 16},
	{r: 'L', n: 18},
	{r: 'R', n: 3},
	{r: 'D', n: 7},
	{r: 'R', n: 11},
	{r: 'D', n: 14},
	{r: 'U', n: 15},
	{r: 'R', n: 9},
	{r: 'D', n: 9},
	{r: 'R', n: 7},
	{r: 'U', n: 17},
	{r: 'R', n: 3},
	{r: 'D', n: 19},
	{r: 'R', n: 4},
	{r: 'U', n: 10},
	{r: 'L', n: 11},
	{r: 'D', n: 2},
	{r: 'L', n: 2},
	{r: 'D', n: 11},
	{r: 'R', n: 14},
	{r: 'D', n: 15},
	{r: 'R', n: 13},
	{r: 'D', n: 10},
	{r: 'R', n: 4},
	{r: 'U', n: 12},
	{r: 'L', n: 13},
	{r: 'U', n: 12},
	{r: 'R', n: 10},
	{r: 'L', n: 4},
	{r: 'U', n: 17},
	{r: 'D', n: 4},
	{r: 'R', n: 4},
	{r: 'U', n: 7},
	{r: 'L', n: 10},
	{r: 'R', n: 12},
	{r: 'L', n: 18},
	{r: 'U', n: 11},
	{r: 'D', n: 8},
	{r: 'R', n: 13},
	{r: 'U', n: 19},
}