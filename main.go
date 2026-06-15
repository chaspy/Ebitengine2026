package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 360
	playerRadius = 18
	playerSpeed  = 4.0
)

type Game struct {
	playerX float32
	playerY float32
}

func NewGame() *Game {
	return &Game{
		playerX: screenWidth / 2,
		playerY: screenHeight / 2,
	}
}

func (g *Game) Update() error {
	dx, dy := g.keyboardDirection()

	if tx, ty, ok := touchOrMousePosition(); ok {
		vx := tx - g.playerX
		vy := ty - g.playerY
		dist := float32(math.Hypot(float64(vx), float64(vy)))
		if dist > playerSpeed {
			dx = vx / dist
			dy = vy / dist
		}
	}

	if dx != 0 || dy != 0 {
		length := float32(math.Hypot(float64(dx), float64(dy)))
		g.playerX += dx / length * playerSpeed
		g.playerY += dy / length * playerSpeed
	}

	g.playerX = clamp(g.playerX, playerRadius, screenWidth-playerRadius)
	g.playerY = clamp(g.playerY, playerRadius, screenHeight-playerRadius)

	return nil
}

func (g *Game) keyboardDirection() (float32, float32) {
	var dx, dy float32
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		dx--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		dx++
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		dy--
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		dy++
	}
	return dx, dy
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 18, G: 22, B: 28, A: 255})

	vector.DrawFilledCircle(screen, g.playerX, g.playerY, playerRadius, color.RGBA{R: 83, G: 212, B: 169, A: 255}, false)
	vector.StrokeCircle(screen, g.playerX, g.playerY, playerRadius, 3, color.RGBA{R: 232, G: 247, B: 239, A: 255}, false)

	ebitenutil.DebugPrint(screen, "Move: Arrow keys / WASD / touch")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func touchOrMousePosition() (float32, float32, bool) {
	touches := ebiten.AppendTouchIDs(nil)
	if len(touches) > 0 {
		x, y := ebiten.TouchPosition(touches[0])
		return float32(x), float32(y), true
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		return float32(x), float32(y), true
	}

	return 0, 0, false
}

func clamp(v, min, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebitengine2026 Prototype")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
