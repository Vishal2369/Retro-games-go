package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 800
	FPS           = 60

	BALL_RADIUS   = 20
	PADDLE_WIDTH  = 25
	PADDLE_HEIGHT = 120
	BALL_SPEED    = 7
	PADDLE_SPEED  = 6
)

var (
	Green      = rl.NewColor(38, 185, 154, 255)
	DarkGreen  = rl.NewColor(20, 160, 133, 255)
	LightGreen = rl.NewColor(129, 204, 184, 255)
	Yellow     = rl.NewColor(234, 213, 154, 255)
)

// GameState represents the current state of the game
type GameState int

const (
	StateHome GameState = iota
	StatePlaying
)

// Ball represents the ball in the game
type Ball struct {
	x, y   float32
	radius float32
	speedX float32
	speedY float32
}

// Paddle represents a paddle
type Paddle struct {
	x, y   float32
	width  float32
	height float32
	speed  float32
	isCPU  bool
	delay  float32 // Delay for AI movement
}

// Game holds the game state
type Game struct {
	state           GameState
	playerScore     int
	cpuScore        int
	ball            *Ball
	playerPaddle    *Paddle
	cpuPaddle       *Paddle
	bounceSound     rl.Sound
	scoreSound      rl.Sound
	backgroundMusic rl.Music
}

// NewGame initializes a new game
func NewGame() *Game {
	return &Game{
		state: StateHome,
		ball: &Ball{
			x: SCREEN_WIDTH / 2, y: SCREEN_HEIGHT / 2,
			radius: BALL_RADIUS, speedX: BALL_SPEED, speedY: BALL_SPEED,
		},
		playerPaddle: &Paddle{
			x: 10, y: SCREEN_HEIGHT/2 - PADDLE_HEIGHT/2,
			width: PADDLE_WIDTH, height: PADDLE_HEIGHT, speed: PADDLE_SPEED,
		},
		cpuPaddle: &Paddle{
			x: SCREEN_WIDTH - 10 - PADDLE_WIDTH, y: SCREEN_HEIGHT/2 - PADDLE_HEIGHT/2,
			width: PADDLE_WIDTH, height: PADDLE_HEIGHT, speed: PADDLE_SPEED, isCPU: true,
			delay: 0.1, // Slight delay to simulate human-like reflexes
		},
	}
}

// LoadAssets loads game assets like sounds and music
func (g *Game) LoadAssets() {
	rl.InitAudioDevice()
	g.bounceSound = rl.LoadSound("bounce.wav")
	g.scoreSound = rl.LoadSound("score.wav")
	g.backgroundMusic = rl.LoadMusicStream("background.mp3")
	rl.PlayMusicStream(g.backgroundMusic)
}

// UnloadAssets frees up resources
func (g *Game) UnloadAssets() {
	rl.UnloadSound(g.bounceSound)
	rl.UnloadSound(g.scoreSound)
	rl.UnloadMusicStream(g.backgroundMusic)
	rl.CloseAudioDevice()
}

// Draw renders the game based on the current state
func (g *Game) Draw() {
	if g.state == StateHome {
		rl.ClearBackground(DarkGreen)
		rl.DrawText("PING PONG", SCREEN_WIDTH/2-200, SCREEN_HEIGHT/2-50, 60, rl.White)
		rl.DrawText("Click to Start", SCREEN_WIDTH/2-150, SCREEN_HEIGHT/2+50, 40, rl.White)
	} else if g.state == StatePlaying {
		rl.ClearBackground(DarkGreen)
		rl.DrawRectangle(0, 0, SCREEN_WIDTH/2, SCREEN_HEIGHT, Green)
		rl.DrawCircle(SCREEN_WIDTH/2, SCREEN_HEIGHT/2, 100, LightGreen)
		rl.DrawLine(SCREEN_WIDTH/2, 0, SCREEN_WIDTH/2, SCREEN_HEIGHT, rl.White)

		g.ball.Draw()
		g.playerPaddle.Draw()
		g.cpuPaddle.Draw()

		rl.DrawText(fmt.Sprintf("%v", g.playerScore), SCREEN_WIDTH/4-20, 20, 80, rl.White)
		rl.DrawText(fmt.Sprintf("%v", g.cpuScore), 3*SCREEN_WIDTH/4-20, 20, 80, rl.White)
	}
}

// Update handles game logic based on the state
func (g *Game) Update() {
	rl.UpdateMusicStream(g.backgroundMusic)

	if g.state == StateHome {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			g.state = StatePlaying
			g.Reset()
		}
	} else if g.state == StatePlaying {
		g.ball.Update(g)
		g.playerPaddle.Update()
		g.cpuPaddle.UpdateAI(g.ball)

		if g.playerScore >= 10 || g.cpuScore >= 10 {
			g.state = StateHome
		}
	}
}

// Reset resets the game state
func (g *Game) Reset() {
	g.playerScore = 0
	g.cpuScore = 0
	g.ball.Reset()
}

// Draw method for Ball
func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.x), int32(b.y), b.radius, Yellow)
}

// Update method for Ball
func (b *Ball) Update(g *Game) {
	b.x += b.speedX
	b.y += b.speedY

	// Bounce on walls
	if b.y+b.radius >= SCREEN_HEIGHT || b.y-b.radius <= 0 {
		b.speedY *= -1
		rl.PlaySound(g.bounceSound)
	}

	// Score
	if b.x+b.radius >= SCREEN_WIDTH {
		g.playerScore++
		b.Reset()
		rl.PlaySound(g.scoreSound)
	}
	if b.x-b.radius <= 0 {
		g.cpuScore++
		b.Reset()
		rl.PlaySound(g.scoreSound)
	}

	// Paddle collision
	if rl.CheckCollisionCircleRec(rl.NewVector2(b.x, b.y), b.radius, g.playerPaddle.Rect()) ||
		rl.CheckCollisionCircleRec(rl.NewVector2(b.x, b.y), b.radius, g.cpuPaddle.Rect()) {
		b.speedX *= -1
		rl.PlaySound(g.bounceSound)
	}
}

// Reset method for Ball
func (b *Ball) Reset() {
	b.x, b.y = SCREEN_WIDTH/2, SCREEN_HEIGHT/2
	b.speedX *= -1
}

// Draw method for Paddle
func (p *Paddle) Draw() {
	rl.DrawRectangleRounded(p.Rect(), 0.8, 0, rl.White)
}

// Rect returns the paddle's rectangle
func (p *Paddle) Rect() rl.Rectangle {
	return rl.NewRectangle(p.x, p.y, p.width, p.height)
}

// Update method for Player Paddle
func (p *Paddle) Update() {
	if rl.IsKeyDown(rl.KeyUp) {
		p.y -= p.speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.y += p.speed
	}
	if p.y < 0 {
		p.y = 0
	}
	if p.y+p.height > SCREEN_HEIGHT {
		p.y = SCREEN_HEIGHT - p.height
	}
}

// UpdateAI method for CPU Paddle
func (p *Paddle) UpdateAI(ball *Ball) {
	targetY := ball.y - p.height/2
	p.y += (targetY - p.y) * p.delay
	if p.y < 0 {
		p.y = 0
	}
	if p.y+p.height > SCREEN_HEIGHT {
		p.y = SCREEN_HEIGHT - p.height
	}
}

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Ping Pong")
	defer rl.CloseWindow()
	rl.SetTargetFPS(FPS)

	game := NewGame()
	game.LoadAssets()
	defer game.UnloadAssets()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		game.Update()
		game.Draw()
		rl.EndDrawing()
	}
}
