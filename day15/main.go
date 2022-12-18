package main

import (
	"fmt"
	"sort"
)

func main() {
	var intervals [][2]int

	for _, row := range input {
		sx, sy := row[0][0], row[0][1]
		bx, by := row[1][0], row[1][1]
		dist := abs(sx-bx) + abs(sy-by)
		yDistToTarget := abs(sy - targetRow)
		xDist := dist - yDistToTarget
		if xDist >= 0 {
			intervals = append(intervals, [2]int{sx - xDist, sx + xDist})
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	// Delete intervals that lie completely in another interval.
	for {
		deleted := func() bool {
			for i := range intervals {
				a := intervals[i]
				for j := range intervals {
					if i != j {
						b := intervals[j]
						if a[0] >= b[0] && a[1] <= b[1] {
							intervals = append(intervals[:i], intervals[i+1:]...)
							return true
						}
					}
				}
			}
			return false
		}()

		if !deleted {
			break
		}
	}

	for {
		merged := false

		for i := 1; i < len(intervals); i++ {
			a, b := intervals[i-1], intervals[i]
			if a[1] >= b[0] {
				c := [2]int{min(a[0], b[0]), max(a[1], b[1])}
				// Replace i-1.
				intervals[i-1] = c
				// Remove i.
				intervals = append(intervals[:i], intervals[i+1:]...)
				merged = true
				break
			}
		}

		if !merged {
			break
		}
	}

	sum := 0
	for _, in := range intervals {
		sum += in[1] - in[0]
	}
	fmt.Println("part 1)", sum)

	// Part 2: we know that a single free spot on the map must have a signal
	// directly next of it to either side.
	// Look at this pair of signals:
	//
	// 	  1
	// 	 111   2
	// 	11X11 222
	// 	 111 22222
	// 	  1 222X222
	// 	     22222
	// 	      222
	// 	       2
	//
	// There is one-square-wide line between signal 1 and 2. For this to be just
	// one free spot, it needs a diamond to the other 2 sides as well.
	// Thus we can look at all the borders around all signals (each signal has
	// 4 lines around it) and look where four of these borders meet. This should
	// be our spot.
	//
	// Turns out that a map enumerating all possible border points around all
	// signals is just fast enough. Takes under a minute and does the job. A
	// faster way (to compute, not to code up) would be to intersect all these
	// border line segments with each other.
	borderCounts := make(map[position]int)
	for _, row := range input {
		sx, sy := row[0][0], row[0][1]
		bx, by := row[1][0], row[1][1]
		dist := abs(sx-bx) + abs(sy-by)

		{
			x := sx
			y := sy - (dist + 1)
			for y < sy {
				borderCounts[pos(x, y)]++
				x++
				y++
			}
		}
		{
			x := sx + (dist + 1)
			y := sy
			for x > sx {
				borderCounts[pos(x, y)]++
				x--
				y++
			}
		}
		{
			x := sx
			y := sy + (dist + 1)
			for y > sy {
				borderCounts[pos(x, y)]++
				x--
				y--
			}
		}
		{
			x := sx - (dist + 1)
			y := sy
			for x < sx {
				borderCounts[pos(x, y)]++
				x++
				y--
			}
		}
	}
	max := 0
	best := 0
	for p, n := range borderCounts {
		if n > max {
			max = n
			best = p.x*4000000 + p.y
		}
	}
	fmt.Println("part 2)", best)
}

type position struct {
	x, y int
}

func pos(x, y int) position {
	return position{x: x, y: y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const targetRow = 2000000

var input = [][2][2]int{
	{{1259754, 1927417}, {1174860, 2000000}},
	{{698360, 2921616}, {1174860, 2000000}},
	{{2800141, 2204995}, {3151616, 2593677}},
	{{3257632, 2621890}, {3336432, 2638865}},
	{{3162013, 3094407}, {3151616, 2593677}},
	{{748228, 577603}, {849414, -938539}},
	{{3624150, 2952930}, {3336432, 2638865}},
	{{2961687, 2430611}, {3151616, 2593677}},
	{{142293, 3387807}, {45169, 4226343}},
	{{3309479, 1598941}, {3336432, 2638865}},
	{{1978235, 3427616}, {2381454, 3683743}},
	{{23389, 1732536}, {1174860, 2000000}},
	{{1223696, 3954547}, {2381454, 3683743}},
	{{3827517, 3561118}, {4094575, 3915146}},
	{{3027894, 3644321}, {2381454, 3683743}},
	{{3523333, 3939956}, {4094575, 3915146}},
	{{2661743, 3988507}, {2381454, 3683743}},
	{{2352285, 2877820}, {2381454, 3683743}},
	{{3214853, 2572272}, {3151616, 2593677}},
	{{3956852, 2504216}, {3336432, 2638865}},
	{{219724, 3957089}, {45169, 4226343}},
	{{1258233, 2697879}, {1174860, 2000000}},
	{{3091374, 215069}, {4240570, 610698}},
	{{3861053, 889064}, {4240570, 610698}},
	{{2085035, 1733247}, {1174860, 2000000}},
}
