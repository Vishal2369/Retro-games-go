package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CELL_WIDTH = 30
	COLUMNS    = 25
	OFFSET     = 75
	FPS        = 60
	INTERVAL   = 0.2
)

var (
	Green     = rl.NewColor(173, 204, 96, 255)
	DarkGreen = rl.NewColor(43, 51, 24, 255)
)
