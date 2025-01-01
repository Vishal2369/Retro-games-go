package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	grid          *Grid
	blocks        []Block
	currBlock     Block
	nextBlock     Block
	rand          *rand.Rand
	lastUpdatedAt float64
	gameOver      bool
	font          rl.Font
	score         int
	music         rl.Music
	clearSound    rl.Sound
	rotateSound   rl.Sound
}

func NewGame() *Game {
	g := &Game{
		grid:   NewGrid(),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		blocks: getAllBlocks(),
	}

	g.currBlock = g.GetRandomBlock()
	g.nextBlock = g.GetRandomBlock()

	return g
}

func (g *Game) LoadAssests() {
	g.font = rl.LoadFontEx("font/monogram.ttf", 64, nil, 0)

	rl.InitAudioDevice()

	g.music = rl.LoadMusicStream("sound/music.mp3")
	rl.PlayMusicStream(g.music)

	g.rotateSound = rl.LoadSound("sound/rotate.mp3")
	g.clearSound = rl.LoadSound("sound/clear.mp3")
}

func (g *Game) UnloadAssests() {
	rl.UnloadFont(g.font)

	rl.UnloadMusicStream(g.music)
	rl.UnloadSound(g.clearSound)
	rl.UnloadSound(g.rotateSound)

	rl.CloseAudioDevice()
}

func (g *Game) ShouldTriggerEvent(sec float64) bool {
	if rl.GetTime()-g.lastUpdatedAt >= sec {
		g.lastUpdatedAt = rl.GetTime()
		return true
	}
	return false
}

func (g *Game) Update() {
	rl.UpdateMusicStream(g.music)
	g.handleInput()

	if g.ShouldTriggerEvent(0.2) {
		g.MoveBlock(1, 0)
	}
}

func (g *Game) Draw() {
	rl.ClearBackground(DarkBlue)
	rl.DrawTextEx(g.font, "Score", rl.NewVector2(365, 15), 38, 2, rl.White)
	rl.DrawTextEx(g.font, "Next", rl.NewVector2(370, 175), 38, 2, rl.White)
	if g.gameOver {
		rl.DrawTextEx(g.font, "Game Over", rl.NewVector2(320, 450), 38, 2, rl.White)
	}

	rl.DrawRectangleRounded(rl.NewRectangle(320, 55, 170, 60), 0.3, 6, LightBlue)
	rl.DrawRectangleRounded(rl.NewRectangle(320, 215, 170, 180), 0.3, 6, LightBlue)

	scoreText := fmt.Sprintf("%d", g.score)
	textSize := rl.MeasureTextEx(g.font, scoreText, 38, 2)
	rl.DrawTextEx(g.font, scoreText, rl.NewVector2(320+(170-textSize.X)/2, 65), 38, 2, rl.White)

	g.grid.Draw()
	g.currBlock.Draw(11, 11)

	switch g.nextBlock.GetId() {
	case 3:
		g.nextBlock.Draw(255, 290)
	case 5:
		g.nextBlock.Draw(255, 280)
	default:
		g.nextBlock.Draw(270, 270)
	}
}

func (g *Game) Reset() {
	g.grid = NewGrid()
	g.blocks = getAllBlocks()
	g.currBlock = g.GetRandomBlock()
	g.nextBlock = g.GetRandomBlock()
	g.score = 0
}

func (g *Game) GetRandomBlock() Block {
	if len(g.blocks) == 0 {
		g.blocks = getAllBlocks()
	}

	randomIdx := g.rand.Intn(len(g.blocks))
	randomBlock := g.blocks[randomIdx]

	g.blocks = append(g.blocks[:randomIdx], g.blocks[randomIdx+1:]...)

	return randomBlock
}

func getAllBlocks() []Block {
	return []Block{
		NewIBlock(),
		NewJBlock(),
		NewLBlock(),
		NewSBlock(),
		NewSBlock(),
		NewTBlock(),
		NewZBlock(),
		NewOBlock(),
	}
}

func (g *Game) handleInput() {
	keyPressed := rl.GetKeyPressed()

	if g.gameOver && keyPressed != 0 {
		g.gameOver = false
		g.Reset()
	}

	switch keyPressed {
	case rl.KeyLeft, rl.KeyA:
		g.MoveBlock(0, -1)
	case rl.KeyRight, rl.KeyD:
		g.MoveBlock(0, 1)
	case rl.KeyUp, rl.KeyW:
		g.RotateBlock(1)
	case rl.KeyDown, rl.KeyS:
		g.MoveBlock(1, 0)
	}
}

func (g *Game) MoveBlock(rows, cols int) {
	if g.gameOver {
		return
	}
	g.currBlock.Move(rows, cols)
	g.ValidateMove(rows, cols)

	if rows == 1 && cols == 0 {
		g.UpdateScore(0, 1)
	}
}

func (g *Game) RotateBlock(dir int) {
	if g.gameOver {
		return
	}
	g.currBlock.Rotate(dir)
	g.ValidateRotate(dir)
}

func (g *Game) ValidateMove(rows, cols int) {
	if g.isBlockOutside() || !g.BlockFits() {
		g.currBlock.Move(-rows, -cols)

		if rows == 1 && cols == 0 {
			g.LockBlock()
		}
	}
}

func (g *Game) ValidateRotate(dir int) {
	if g.isBlockOutside() || !g.BlockFits() {
		g.currBlock.Rotate(-dir)
	} else {
		rl.PlaySound(g.rotateSound)
	}
}

func (g *Game) isBlockOutside() bool {
	for _, pos := range g.currBlock.GetCellPositions() {
		if g.grid.IsCellOutside(pos.row, pos.col) {
			return true
		}
	}
	return false
}

func (g *Game) BlockFits() bool {
	for _, pos := range g.currBlock.GetCellPositions() {
		if !g.grid.IsCellEmpty(pos.row, pos.col) {
			return false
		}
	}
	return true
}

func (g *Game) LockBlock() {
	tiles := g.currBlock.GetCellPositions()
	for _, tile := range tiles {
		g.grid.cells[tile.row][tile.col] = g.currBlock.GetId()
	}

	g.currBlock = g.nextBlock
	if !g.BlockFits() {
		g.gameOver = true
	}

	g.nextBlock = g.GetRandomBlock()

	lineCleared := g.grid.ClearFullRows()

	if lineCleared > 0 {
		rl.PlaySound(g.clearSound)
		g.UpdateScore(lineCleared, 0)
	}
}

func (g *Game) UpdateScore(lineCleared, moveDownPoints int) {
	switch lineCleared {
	case 1:
		g.score += 100
	case 2:
		g.score += 300
	case 3:
		g.score += 500
	}
	g.score += moveDownPoints
}
