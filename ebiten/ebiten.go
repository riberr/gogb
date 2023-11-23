package ebiten

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gogb/gameboy"
	"log"
)

type Game struct {
	gb *gameboy.GameBoy
}

const (
	screenWidth    = 1024
	screenHeight   = 768
	Scale          = 4
	CyclesPerFrame = 69905
	DebugWidth     = (16 * 8 * Scale) + (16 * Scale)
	DebugHeight    = (24 * 8 * Scale) + (24 * Scale)
)

func (g *Game) Update() error {
	cyclesThisUpdate := 0
	// 4194304 (cpuFreq) / 60 (targetFPS) = 69905
	for cyclesThisUpdate < CyclesPerFrame {
		cyclesThisUpdate += g.gb.Step()
	}
	g.gb.GenerateGraphics()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.gb.Ppu.FbVram.Pix)
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
	ebiten.SetWindowSize(DebugWidth, DebugHeight)
	ebiten.SetWindowTitle("Noise (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
