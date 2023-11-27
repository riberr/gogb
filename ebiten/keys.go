package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gogb/gameboy/joypad"
)

func (g *Game) keyDown() {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		g.gb.JoyPad.ButtonPressed(joypad.PadLeft)
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		g.gb.JoyPad.ButtonPressed(joypad.PadRight)
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		g.gb.JoyPad.ButtonPressed(joypad.PadUp)
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		g.gb.JoyPad.ButtonPressed(joypad.PadDown)
	case inpututil.IsKeyJustPressed(ebiten.KeyZ):
		g.gb.JoyPad.ButtonPressed(joypad.PadA)
	case inpututil.IsKeyJustPressed(ebiten.KeyX):
		g.gb.JoyPad.ButtonPressed(joypad.PadB)
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		g.gb.JoyPad.ButtonPressed(joypad.PadSelect)
	case inpututil.IsKeyJustPressed(ebiten.KeyEnter):
		g.gb.JoyPad.ButtonPressed(joypad.PadStart)
	}
}

func (g *Game) keyUp() {
	switch {
	case inpututil.IsKeyJustReleased(ebiten.KeyLeft):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	case inpututil.IsKeyJustReleased(ebiten.KeyRight):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	case inpututil.IsKeyJustReleased(ebiten.KeyUp):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	case inpututil.IsKeyJustReleased(ebiten.KeyDown):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	case inpututil.IsKeyJustReleased(ebiten.KeyZ):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	case inpututil.IsKeyJustReleased(ebiten.KeyX):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	case inpututil.IsKeyJustReleased(ebiten.KeySpace):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	case inpututil.IsKeyJustReleased(ebiten.KeyEnter):
		g.gb.JoyPad.ButtonReleased(joypad.PadStart)
	}
}
