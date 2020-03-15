package leetcode

import "fmt"

var highLen = 0
var wideLen = 0

func PrintGrid(grid [][]int) {
	for _, list := range grid {
		fmt.Printf("%v\n", list)

	}
}

func maxAreaOfIsland(grid [][]int) int {
	// 获取二维数组的宽高
	highLen = len(grid)
	if highLen == 0 {
		return 0
	}
	wideLen = len(grid[0])
	fmt.Printf("high=%d, wide=%d\n", highLen, wideLen)
	PrintGrid(grid)

	maxArea := 0
	thisArea := 0
	for i := 0; i < highLen; i++ { // 外层数组
		for j := 0; j < wideLen; j++ { // 内层数组
			// 遍历每一块可以组成的岛屿面积
			fmt.Printf("before i=%d, j=%d\n", i, j)
			thisArea = maxAreaOfIslandDFS(grid, i, j)
			fmt.Printf("area=%d, i=%d, j=%d\n", thisArea, i, j)
			PrintGrid(grid)

			if thisArea > maxArea {
				maxArea = thisArea
			}
		}
	}

	return maxArea
}

func maxAreaOfIslandDFS(grid [][]int, i, j int) int {
	fmt.Printf("DFS: i=%d,j=%d,high=%d,wide=%d\n", i, j, highLen, wideLen)
	if i < 0 || j < 0 || i >= highLen || j >= wideLen {
		fmt.Printf("DFS:return\n\n")
		return 0
	} else if grid[i][j] == 0 {
		fmt.Printf("DFS:return grid==0,i=%d,j=%d\n\n", i, j)
		return 0
	}
	// 遍历过这块则置0避免再次计算
	grid[i][j] = 0
	count := 1
	fmt.Printf("set i=%d,j=%d to 0\n", i, j)

	// 计算四周
	count = count + maxAreaOfIslandDFS(grid, i-1, j)
	count = count + maxAreaOfIslandDFS(grid, i, j-1)
	count = count + maxAreaOfIslandDFS(grid, i+1, j)
	count = count + maxAreaOfIslandDFS(grid, i, j+1)

	return count
}
