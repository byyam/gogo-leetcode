package leetcode

import (
	"fmt"
	toolkit "github.com/byyam/go-leetcode-toolkit"
)

func maxDistance(grid [][]int) int {
	highLen := len(grid)
	if highLen == 0 {
		return -1
	}
	wideLen := len(grid[0])
	if highLen != wideLen {
		return -1
	}
	// 腐烂橘子问题：从每个陆地作为BFS的源开始扩散，直到周围没有海水可以被覆盖
	type Point struct {
		X int
		Y int
		V int
	}
	printQueue := func(name string, queue []*Point) {
		fmt.Printf("%s:", name)
		for i := 0; i < len(queue); i++ {
			fmt.Printf("[%d][%d]=%d ", queue[i].X, queue[i].Y, queue[i].V)
		}
		fmt.Printf("\n")
	}
	// 把每个源放在队列里, 初始的陆地都标记为1
	var queue []*Point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, &Point{i, j, 1})
			}
		}
	}
	// 全为海水或陆地
	if len(queue) == 0 || len(queue) == len(grid)*len(grid[0]) {
		return -1
	}
	ans := 0
	d := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) > 0 { // 从每个源出发
		printQueue("queue", queue)
		ans++
		tmpQueue := queue
		queue = nil
		for len(tmpQueue) > 0 {
			printQueue("tmp queue", tmpQueue)
			p := tmpQueue[0]
			value := p.V + 1
			tmpQueue = tmpQueue[1:] // pop
			// 以p为原点，检查4个方向
			for i := 0; i < 4; i++ {
				x := p.X + d[i][0]
				y := p.Y + d[i][1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != 0 {
					continue
				}
				queue = append(queue, &Point{x, y, value}) // push
				grid[x][y] = value
			}
		}
		toolkit.PrintGrid(grid)
	}

	return ans - 1
}
