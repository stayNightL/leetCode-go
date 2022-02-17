package offer

func movingCount(m int, n int, k int) int {
	roboot_map := make([][]int, m)
	for i := range roboot_map {
		roboot_map[i] = make([]int,n)
	}
	if k == 0 {
		return 1
	}
	max :=dfs(roboot_map,0,0,k)
return max

}

func dfs(roboot_map [][]int, x int, y int,k int) int {
	if x < 0 || y < 0 || x > len(roboot_map[0])-1 || y > len(roboot_map)-1 || roboot_map[y][x] ==1 {
		return 0
	}
	if sum(x)+sum(y) > k {
		roboot_map[y][x] = 1
		return 0
	}
	roboot_map[y][x] = 1
	res := 1+dfs(roboot_map,x-1,y,k)+dfs(roboot_map,x+1,y,k)+dfs(roboot_map,x,y-1,k)+dfs(roboot_map,x,y+1,k)
	return res
}

func sum(m int) int {
	 n := 0
	for m != 0 {
		n = n + m % 10
		m = m / 10
	}
	return n
}
