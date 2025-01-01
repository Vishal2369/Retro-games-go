package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 500
	SCREEN_HEIGHT = 620
	FPS           = 60
)

var (
	DarkBlue  = rl.NewColor(44, 44, 127, 255)
	LightBlue = rl.NewColor(59, 85, 162, 255)

	CellColors = []rl.Color{
		rl.NewColor(26, 31, 40, 255),
		rl.NewColor(47, 230, 23, 255),
		rl.NewColor(232, 18, 18, 255),
		rl.NewColor(226, 116, 17, 255),
		rl.NewColor(237, 234, 4, 255),
		rl.NewColor(166, 0, 247, 255),
		rl.NewColor(21, 204, 209, 255),
		rl.NewColor(13, 64, 216, 255),
	}
)
