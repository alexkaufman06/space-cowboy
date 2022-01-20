package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
	"log"
)

var spaceCowboyImg *ebiten.Image
var spaceCowboyX float64 = 1
var spaceCowboyY float64 = 1
var touchIDs []ebiten.TouchID
var gamepadIDs []ebiten.GamepadID

func init() {
	var err error
	spaceCowboyImg, _, err = ebitenutil.NewImageFromFile("space-cowboy.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	// Gravity
	spaceCowboyY += 1
	if spaceCowboyY > 200 {
		spaceCowboyY = 200
	}

	if g.isKeyJustPressed() {
		spaceCowboyY -= 15
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Welcome to Space Cowboy!", 90, 5)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(1.2, spaceCowboyY )
	screen.DrawImage(spaceCowboyImg, op)
}

func (g *Game) isKeyJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}
	touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs)
	if len(touchIDs) > 0 {
		return true
	}
	gamepadIDs = ebiten.AppendGamepadIDs(gamepadIDs[:0])
	for _, g := range gamepadIDs {
		if ebiten.IsStandardGamepadLayoutAvailable(g) {
			if inpututil.IsStandardGamepadButtonJustPressed(g, ebiten.StandardGamepadButtonRightBottom) {
				return true
			}
			if inpututil.IsStandardGamepadButtonJustPressed(g, ebiten.StandardGamepadButtonRightRight) {
				return true
			}
		} else {
			// The button 0/1 might not be A/B buttons.
			if inpututil.IsGamepadButtonJustPressed(g, ebiten.GamepadButton0) {
				return true
			}
			if inpututil.IsGamepadButtonJustPressed(g, ebiten.GamepadButton1) {
				return true
			}
		}
	}
	return false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Space Cowboy")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
