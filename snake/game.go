package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	food              *Food
	snake             *Snake
	updateTriggeredAt float64
	running           bool
	points            int
	eatSound          rl.Sound
	wallSound         rl.Sound
}

func NewGame() *Game {
	deque := getDefaultSnakePos()

	return &Game{
		food:  &Food{position: GenerateRandomPosition(deque.AllElements())},
		snake: &Snake{body: deque, dir: rl.NewVector2(1, 0)},
	}
}

func (g *Game) Reset() {
	deque := getDefaultSnakePos()

	g.running = false
	g.food.position = GenerateRandomPosition(deque.AllElements())
	g.snake.body = deque
	g.snake.dir = rl.NewVector2(1, 0)
	g.points = 0
}

func getDefaultSnakePos() *Deque[rl.Vector2] {
	deque := NewDeque[rl.Vector2]()
	deque.PushFront(rl.NewVector2(5, 9))
	deque.PushFront(rl.NewVector2(6, 9))
	deque.PushFront(rl.NewVector2(7, 9))
	return deque
}

func (g *Game) LoadAssests() {
	g.food.LoadTexture()

	rl.InitAudioDevice()
	g.eatSound = rl.LoadSound("Sounds/eat.mp3")
	g.wallSound = rl.LoadSound("Sounds/wall.mp3")
}

func (g *Game) UnloadAssests() {
	g.food.UnloadTexture()

	rl.UnloadSound(g.eatSound)
	rl.UnloadSound(g.wallSound)

	if rl.IsAudioDeviceReady() {
		rl.CloseAudioDevice()
	}
}

func (g *Game) Draw() {
	rl.ClearBackground(Green)
	rl.DrawRectangleLinesEx(rl.NewRectangle(OFFSET-5, OFFSET-5, CELL_WIDTH*COLUMNS+10,
		CELL_WIDTH*COLUMNS+10), 5, DarkGreen)
	rl.DrawText("Retro Snake", OFFSET-5, 20, 40, DarkGreen)
	rl.DrawText(fmt.Sprintf("%d", g.points), OFFSET-5, OFFSET+CELL_WIDTH*COLUMNS+10, 40, DarkGreen)
	g.food.Draw()
	g.snake.Draw()
}

func (g *Game) handleInput() {
	if rl.IsKeyPressed(rl.KeyUp) && g.snake.dir.Y != 1 {
		g.snake.dir = rl.NewVector2(0, -1)
		g.running = true
	} else if rl.IsKeyPressed(rl.KeyRight) && g.snake.dir.X != -1 {
		g.snake.dir = rl.NewVector2(1, 0)
		g.running = true
	} else if rl.IsKeyPressed(rl.KeyDown) && g.snake.dir.Y != -1 {
		g.snake.dir = rl.NewVector2(0, 1)
		g.running = true
	} else if rl.IsKeyPressed(rl.KeyLeft) && g.snake.dir.X != 1 {
		g.snake.dir = rl.NewVector2(-1, 0)
		g.running = true
	}
}

func (g *Game) Update() {
	g.handleInput()

	if !g.running {
		return
	}

	// Update every 200ms
	if rl.GetTime()-g.updateTriggeredAt < INTERVAL {
		return
	}
	g.updateTriggeredAt = rl.GetTime()
	g.snake.body.PushFront(rl.Vector2Add(g.snake.body.Front(), g.snake.dir))

	// check collision with food
	if g.checkCollisionWithFood() {
		rl.PlaySound(g.eatSound)
		g.food.position = GenerateRandomPosition(g.snake.body.AllElements())
		g.points++
	} else {
		g.snake.body.PopBack()
	}

	// check collision with wall
	if g.checkCollisionWithWall() || g.checkCollisionWithBody() {
		rl.PlaySound(g.wallSound)
		g.Reset()
	}
}

func (g *Game) checkCollisionWithFood() bool {
	return rl.Vector2Equals(g.snake.body.Front(), g.food.position)
}

func (g *Game) checkCollisionWithWall() bool {
	head := g.snake.body.Front()
	return head.X < 0 || head.X >= COLUMNS || head.Y < 0 || head.Y >= COLUMNS
}

func (g *Game) checkCollisionWithBody() bool {
	for i, pos := range g.snake.body.AllElements() {
		if i != 0 && rl.Vector2Equals(pos, g.snake.body.Front()) {
			return true
		}
	}
	return false
}
