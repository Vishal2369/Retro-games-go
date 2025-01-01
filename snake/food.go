package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	position rl.Vector2
	texture  rl.Texture2D
}

func (f *Food) LoadTexture() {
	image := rl.LoadImage("Graphics/food.png")
	defer rl.UnloadImage(image)

	f.texture = rl.LoadTextureFromImage(image)
}

func (f *Food) UnloadTexture() {
	rl.UnloadTexture(f.texture)
}

func (f *Food) Draw() {
	rl.DrawTexture(f.texture, (int32)(f.position.X)*CELL_WIDTH+OFFSET,
		(int32)(f.position.Y)*CELL_WIDTH+OFFSET, rl.White)
}

func GenerateRandomPosition(snakeBodyPositions []rl.Vector2) rl.Vector2 {
	position := generateRandomCell()
	for elementInArr(snakeBodyPositions, position) {
		position = generateRandomCell()
	}
	return position
}

func elementInArr(arr []rl.Vector2, elem rl.Vector2) bool {
	for _, v := range arr {
		if rl.Vector2Equals(v, elem) {
			return true
		}
	}
	return false
}

func generateRandomCell() rl.Vector2 {
	x := (float32)(rl.GetRandomValue(0, COLUMNS-1))
	y := (float32)(rl.GetRandomValue(0, COLUMNS-1))
	return rl.NewVector2(x, y)
}
