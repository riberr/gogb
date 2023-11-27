package ebiten

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gogb/gameboy"
	"log"
)

type Game struct {
	gb           *gameboy.GameBoy
	activeScreen int
}

const (
	screenWidth    = 160 * Scale
	screenHeight   = 144 * Scale
	Scale          = 4
	CyclesPerFrame = 69905
	DebugWidth     = (16 * 8 * Scale) + (16 * Scale)
	DebugHeight    = (24 * 8 * Scale) + (24 * Scale)
)

func (g *Game) Update() error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyDigit1):
		g.activeScreen = 1
		ebiten.SetWindowSize(screenWidth, screenHeight)
	case inpututil.IsKeyJustPressed(ebiten.KeyDigit2):
		g.activeScreen = 2
		ebiten.SetWindowSize(DebugWidth, DebugHeight)
	}

	g.keyDown()
	g.keyUp()

	cyclesThisUpdate := 0
	// 4194304 (cpuFreq) / 60 (targetFPS) = 69905
	for cyclesThisUpdate < CyclesPerFrame {
		cyclesThisUpdate += g.gb.Step()
	}
	g.gb.GenerateGraphics()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	switch g.activeScreen {
	case 1:
		// this is to catch the panic that is thrown when switching window size
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		screen.WritePixels(g.gb.Ppu.Fb.Pix) // dunno which one is faster
	case 2:
		// this is to catch the panic that is thrown when switching window size
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		screen.WritePixels(g.gb.Ppu.FbVram.Pix) // dunno which one is faster
		//screen.DrawImage(ebiten.NewImageFromImage(g.gb.Ppu.FbVram), nil)

	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New(gb *gameboy.GameBoy) *Game {
	return &Game{
		gb:           gb,
		activeScreen: 2,
	}
}

func (g *Game) Run() {
	ebiten.SetWindowSize(DebugWidth, DebugHeight)
	ebiten.SetWindowTitle("Noise (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
