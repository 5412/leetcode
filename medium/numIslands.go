package medium

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cells := len(grid[0])
	nums := 0
	visited := make(map[int]map[int]bool, 0)
	for i := 0; i < rows; i++ {
		if visited[i] == nil {
			visited[i] = make(map[int]bool, cells)
		}
		for j := 0; j < cells; j++ {
			if visited[i][j] {
				continue
			}
			if grid[i][j] == '1' {
				nums++
				numIslandsHelper(grid, visited, i, j)
			}
		}
	}
	return nums
}

func numIslandsHelper(grid [][]byte, m map[int]map[int]bool, i, j int) {
	//fmt.Println(m)
	if grid[i][j] == '1' && !m[i][j] {
		if m[i] == nil {
			m[i] = make(map[int]bool, len(grid[0]))
		}
		m[i][j] = true
		if j+1 < len(grid[0]) {
			numIslandsHelper(grid, m, i, j+1)
		}
		if j-1 >= 0 && !m[i][j-1] {
			numIslandsHelper(grid, m, i, j-1)
		}
		if i+1 < len(grid) {
			numIslandsHelper(grid, m, i+1, j)
		}
		if i-1 >=0 && !m[i-1][j] {
			numIslandsHelper(grid, m, i-1, j)
		}
	}
}