package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snake struct {
	body *Deque[rl.Vector2]
	dir  rl.Vector2
}

func (s *Snake) Draw() {
	for _, pos := range s.body.AllElements() {
		segment := rl.NewRectangle(OFFSET+pos.X*CELL_WIDTH,
			OFFSET+pos.Y*CELL_WIDTH, CELL_WIDTH, CELL_WIDTH)

		rl.DrawRectangleRounded(segment, 0.5, 6, DarkGreen)
	}
}
