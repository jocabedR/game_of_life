package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/game", HandleGame).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	log.Println("Listening at port " + srv.Addr)
	srv.ListenAndServe()

	//grid = aplyRules(grid)

}

func HandleGame(w http.ResponseWriter, r *http.Request) {
	/* w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var grid [][]bool
	err := json.NewDecoder(r.Body).Decode(&grid)
	if err != nil {
		log.Fatalln("There was an error decoding the request body.")
	}

	print(grid) */

	decoder := json.NewDecoder(r.Body)

	var grid [][]bool
	var gridResponse [][]bool

	decoder.Decode(&grid)

	gridResponse = aplyRules(grid)

	fmt.Println(gridResponse)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(gridResponse); err != nil {
		panic(err)
	}

}

func aplyRules(grid [][]bool) [][]bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {

			is_alive := grid[i][j]
			live_neighbors := liveNeighbors(j, i, grid)
			if is_alive {
				if live_neighbors <= 1 || live_neighbors >= 4 {
					grid[i][j] = false
				}
			} else {
				if live_neighbors == 3 {
					for z := i - 1; z <= i+1; z++ {
						for y := i - 1; y <= j+1; y++ {
							grid[z][y] = false
						}
					}
					grid[i][j] = true
				}
			}

			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
	return grid
}

func liveNeighbors(x, y int, grid [][]bool) int {
	count := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if (i == y && j == x) || (i < 0 || j < 0) || (i >= len(grid) || j >= len(grid)) {
				continue
			}
			if grid[i][j] {
				count++
			}
		}
	}
	return count
}

func limit(cor, lim int) int {
	if cor >= 0 && cor < lim {
		return 0
	} else {
		return 1
	}
}
