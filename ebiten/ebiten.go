package ebiten

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gogb/gameboy"
	"image"
	"image/color"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type rand struct {
	x, y, z, w uint32
}

func (r *rand) next() uint32 {
	// math/rand is too slow to keep 60 FPS on web browsers.
	// Use Xorshift instead: http://en.wikipedia.org/wiki/Xorshift
	t := r.x ^ (r.x << 11)
	r.x, r.y, r.z = r.y, r.z, r.w
	r.w = (r.w ^ (r.w >> 19)) ^ (t ^ (t >> 8))
	return r.w
}

var theRand = &rand{12345678, 4185243, 776511, 45411}

type Game struct {
	noiseImage *image.RGBA
	gb         *gameboy.GameBoy
}

const CYCLES = 69905

func (g *Game) Update() error {
	// Generate the noise with random RGB values.
	const l = screenWidth * screenHeight
	for i := 0; i < l; i++ {
		x := theRand.next()
		g.noiseImage.Pix[4*i] = uint8(x >> 24)
		g.noiseImage.Pix[4*i+1] = uint8(x >> 16)
		g.noiseImage.Pix[4*i+2] = uint8(x >> 8)
		g.noiseImage.Pix[4*i+3] = 0xff
	}

	cyclesThisUpdate := 0
	for cyclesThisUpdate < CYCLES {
		cyclesThisUpdate += g.gb.Step()
	}
	return nil
}

const scale = 4

// https://github.com/rockytriton/LLD_gbemu/blob/main/part11/lib/ui.c
func (g *Game) drawTile(screen *ebiten.Image, startLocation uint16, tileNum uint16, x int, y int) {
	//rc := ebiten.NewImage(scale, scale)
	//var colors = int[]{0xFFFFFFFF, 0xFFAAAAAA, 0xFF555555, 0xFF000000}
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

			fmt.Printf("%v, %v, %v, %v, %v\n", rcx, rcy, rcw, rch, colors[c])

			vector.DrawFilledRect(screen, float32(rcx), float32(rcy), float32(rcw), float32(rch), colors[c], false)
		}
	}
	//return rc
}

func (g *Game) Draw(screen *ebiten.Image) {

	//screen.WritePixels(g.noiseImage.Pix)
	//dst := ebiten.Image{}
	vector.DrawFilledRect(screen, 0, 0, 100, 100, color.White, false)

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
	return screenWidth, screenHeight
}

func New(gb *gameboy.GameBoy) *Game {
	return &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
		gb:         gb,
	}
}

func (g *Game) Run() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Noise (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

/*
func Run() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Noise (Ebitengine Demo)")
	g := &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
*/
