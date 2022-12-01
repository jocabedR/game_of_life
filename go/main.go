package main

import "fmt"

func main() {
	fmt.Println("holis")

	width := 5
	height := 5

	nombres := [height][width]string{{" ", " ", " ", "*", "*"}, {" ", " ", " ", " ", " "}, {" ", " ", " ", " ", " "}, {" ", " ", " ", " ", " "}, {" ", " ", " ", " ", " "}}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			n := liveNeighbors(i, j, width, height, nombres)
			nombres[i][j] = aplyRules(i, j, n, nombres)
		}
	}

}

func aplyRules(x, y, liveNeighbors int, grid [][]string) string {
	if isAlive(grid, x, y) == 1 {
		switch liveNeighbors {
		case 2:
		case 3:
			return "*"
			break
		default:
			return " "
			break
		}
	} else if liveNeighbors == 3 {
		return "*"
	}

}

func liveNeighbors(x, y, width, height int, grid [][]string) int {
	cont := 0
	for i := y - 1; i < y+2; i++ {
		for j := x - 1; j < x+2; j++ {
			if limit(j, width) == 0 && limit(i, height) == 0 {
				cont = cont + isAlive(grid, j, i)
			}
		}
	}
	return cont
}

func limit(cor, lim int) int {
	if cor >= 0 && cor < lim {
		return 0
	} else {
		return 1
	}
}

func isAlive(arr [][]string, x, y int) int {
	if arr[y][x] == "*" {
		return 1
	} else {
		return 0
	}
}
