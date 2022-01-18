package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

var spaceCowboyImg *ebiten.Image
var spaceCowboyX float64 = 1
var spaceCowboyY float64 = 1

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
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Welcome to Space Cowboy!", 90, 5)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(1.2, spaceCowboyY )
	screen.DrawImage(spaceCowboyImg, op)
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
