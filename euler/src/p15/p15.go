package main

type Key struct {
	row, col int
}

var routes_map = map[Key]int{}

// tl: top left; br: bottom right

func routes_tl_to_br(row int, col int) int {
	if val, ok := routes_map[Key{row, col}]; ok {
		return val
	}

	var routes int
	if row < 0 || col < 0 {
		routes = 0
	} else if row == 0 || col == 0 {
		routes = 1
	} else {

		routes = routes_tl_to_br(row-1, col) + routes_tl_to_br(col-1, row)
		routes_map[Key{row, col}] = routes
	}

	return routes
}

func routes_tl_to_br_loop(row int, col int) int {
	routes_map[Key{-1, -1}] = 0
	routes_map[Key{-1, 0}] = 0
	routes_map[Key{0, -1}] = 0
	routes_map[Key{0, 0}] = 1

	for r := 0; r <= row; r++ {
		for c := 0; c <= col; c++ {
			if _, ok := routes_map[Key{r, c}]; !ok {
				routes_map[Key{r, c}] = routes_map[Key{r - 1, c}] + routes_map[Key{r, c - 1}]
			}
		}
	}

	return routes_map[Key{row, col}]
}
