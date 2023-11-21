package ebiten

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gogb/gameboy"
	"image/color"
	"log"
)

type Game struct {
	gb *gameboy.GameBoy
}

const (
	screenWidth    = 1024
	screenHeight   = 768
	scale          = 4
	cyclesPerFrame = 69905
	debugWidth     = (16 * 8 * scale) + (16 * scale)
	debugHeight    = (24 * 8 * scale) + (24 * scale)
)

func (g *Game) Update() error {
	cyclesThisUpdate := 0
	for cyclesThisUpdate < cyclesPerFrame {
		cyclesThisUpdate += g.gb.Step()
	}
	return nil
}

// https://github.com/rockytriton/LLD_gbemu/blob/main/part11/lib/ui.c
func (g *Game) drawTile(screen *ebiten.Image, startLocation uint16, tileNum uint16, x int, y int) {
	colors := []color.RGBA{
		{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xFF},
		{R: 0x55, G: 0x55, B: 0x55, A: 0xFF},
		{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
	}

	for tileY := uint16(0); tileY < 16; tileY += 2 {
		b1 := g.gb.Bus.Read(startLocation + (tileNum * 16) + tileY)
		b2 := g.gb.Bus.Read(startLocation + (tileNum * 16) + tileY + 1)

		for bit := 7; bit >= 0; bit-- {
			//hi := !!(b1 & (1 << bit)) << 1
			//lo := !!(b2 & (1 << bit))

			hi := uint8(b1 & (1 << bit))
			if hi > 0 {
				hi = 1 << 1
			} else {
				hi = 0
			}

			lo := uint8(b2 & (1 << bit))
			if lo > 0 {
				lo = 1
			} else {
				lo = 0
			}

			c := hi | lo // color

			rcx := x + ((7 - bit) * scale)
			rcy := y + (int(tileY) / 2 * scale)
			rcw := scale
			rch := scale

			//fmt.Printf("%v, %v, %v, %v, %v\n", rcx, rcy, rcw, rch, colors[c])

			vector.DrawFilledRect(screen, float32(rcx), float32(rcy), float32(rcw), float32(rch), colors[c], false)
		}
	}
	//return rc
}

func (g *Game) Draw(screen *ebiten.Image) {
	xDraw := 0
	yDraw := 0
	tileNum := uint16(0)
	addr := uint16(0x8000)

	//384 tiles, 24 x 16
	for y := 0; y < 24; y++ {
		for x := 0; x < 16; x++ {
			g.drawTile(screen, addr, tileNum, xDraw+(x*scale), yDraw+(y*scale))
			xDraw += 8 * scale
			tileNum++

			//screen.DrawImage(tile, nil)
		}

		yDraw += 8 * scale
		xDraw = 0
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New(gb *gameboy.GameBoy) *Game {
	return &Game{
		gb: gb,
	}
}

func (g *Game) Run() {
	ebiten.SetWindowSize(debugWidth, debugHeight)
	ebiten.SetWindowTitle("Noise (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
