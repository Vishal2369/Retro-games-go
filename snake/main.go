package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(2*OFFSET+CELL_WIDTH*COLUMNS, 2*OFFSET+CELL_WIDTH*COLUMNS, "classic-snake")
	defer rl.CloseWindow()

	rl.SetTargetFPS(FPS)

	game := NewGame()

	game.LoadAssests()
	defer game.UnloadAssests()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		game.Update()
		game.Draw()
		rl.EndDrawing()
	}
}
