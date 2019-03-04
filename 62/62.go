package uniquePaths

func uniquePaths(m int, n int) int {
	orderx, ordery := m, n
	startx, starty := 1, 1
	count := 0
	run(&count, startx, starty, orderx, ordery)
	return count
}
func run(count *int, x int, y int, ox int, oy int) {
	next_x := x + 1
	next_y := y + 1
	//--
	if ox == x && y == oy {
		*count += 1
		return
	}
	if ox == x && oy == y {
		*count += 1
		return
	}
	//--
	if ox == x {
		run(count, x, next_y, ox, oy)
		return
	}
	if oy == y {
		run(count, next_x, y, ox, oy)
		return
	}
	run(count, next_x, y, ox, oy)
	run(count, x, next_y, ox, oy)

}
