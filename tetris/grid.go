package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	rows     int
	cols     int
	cellSize int
	cells    [][]int
}

func NewGrid() *Grid {
	g := &Grid{}
	g.rows = 20
	g.cols = 10
	g.cellSize = 30
	g.cells = make([][]int, g.rows)

	for i := range g.rows {
		g.cells[i] = make([]int, g.cols)
	}

	return g
}

func (g *Grid) Draw() {
	for r := range g.rows {
		for c := range g.cols {
			x := (int32)(c*g.cellSize + 11)
			y := (int32)(r*g.cellSize + 11)
			cellSize := (int32)(g.cellSize - 1)

			rl.DrawRectangle(x, y, cellSize, cellSize, CellColors[g.cells[r][c]])
		}
	}
}

func (g *Grid) IsCellOutside(rows, cols int) bool {
	return rows < 0 || rows >= g.rows || cols < 0 || cols >= g.cols
}

func (g *Grid) IsCellEmpty(rows, cols int) bool {
	return g.cells[rows][cols] == 0
}

func (g *Grid) IsRowFull(row int) bool {
	for c := range g.cols {
		if g.cells[row][c] == 0 {
			return false
		}
	}
	return true
}

func (g *Grid) ClearRow(row int) {
	for c := range g.cols {
		g.cells[row][c] = 0
	}
}

func (g *Grid) MoveRowDown(row, numRows int) {
	for c := range g.cols {
		g.cells[row+numRows][c] = g.cells[row][c]
		g.cells[row][c] = 0
	}
}

func (g *Grid) ClearFullRows() int {
	completed := 0
	for r := g.rows - 1; r >= 0; r-- {
		if g.IsRowFull(r) {
			g.ClearRow(r)
			completed++
		} else if completed > 0 {
			g.MoveRowDown(r, completed)
		}
	}

	return completed
}
