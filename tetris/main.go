package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Tetris")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	g := NewGame()
	g.LoadAssests()

	defer g.UnloadAssests()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}
