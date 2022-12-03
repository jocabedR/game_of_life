package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Grid struct {
	Grid [][]bool `json:"grid"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/game", HandleGame).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	log.Println("Listening at port " + srv.Addr)
	srv.ListenAndServe()
}

func HandleGame(w http.ResponseWriter, r *http.Request) {
	var input Grid

	body := json.NewDecoder(r.Body)
	if err := body.Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	grid := input.Grid
	gridResponse := aplyRules(grid)

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
							if z >= 0 && y >= 0 {
								if z < len(grid) && y < len(grid) {
									grid[z][y] = false
								}
							}
						}
					}
					grid[i][j] = true
				}
			}
		}
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
