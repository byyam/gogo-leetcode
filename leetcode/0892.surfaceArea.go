package leetcode

import "log"

func surfaceArea(grid [][]int) int {
	area := 0
	high := len(grid)
	if high == 0 {
		return 0
	}
	wide := len(grid[0])
	log.Printf("high:%d, wide:%d", high, wide)
	if high != wide {
		return 0
	}

	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			if grid[i][j] == 0 {
				continue
			}
			singleArea := 1 + 1 + grid[i][j]*4
			if i-1 >= 0 {
				minArea := grid[i][j]
				if grid[i-1][j] < minArea {
					minArea = grid[i-1][j]
				}
				singleArea = singleArea - minArea
			}
			if j-1 >= 0 {
				minArea := grid[i][j]
				if grid[i][j-1] < minArea {
					minArea = grid[i][j-1]
				}
				singleArea = singleArea - minArea
			}
			if i+1 < wide {
				minArea := grid[i][j]
				if grid[i+1][j] < minArea {
					minArea = grid[i+1][j]
				}
				singleArea = singleArea - minArea
			}
			if j+1 < high {
				minArea := grid[i][j]
				if grid[i][j+1] < minArea {
					minArea = grid[i][j+1]
				}
				singleArea = singleArea - minArea
			}
			area += singleArea
			log.Printf("get grid[%d][%d]=%d single area=%d, area=%d", i, j, grid[i][j], singleArea, area)
		}
	}

	return area
}
