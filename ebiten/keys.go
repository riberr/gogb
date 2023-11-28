package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gogb/gameboy/joypad"
)

func (g *Game) handleKeys() {
	// press
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.gb.JoyPad.KeyEvent(joypad.PadLeft, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.gb.JoyPad.KeyEvent(joypad.PadRight, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.gb.JoyPad.KeyEvent(joypad.PadUp, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.gb.JoyPad.KeyEvent(joypad.PadDown, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.gb.JoyPad.KeyEvent(joypad.PadA, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		g.gb.JoyPad.KeyEvent(joypad.PadB, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.gb.JoyPad.KeyEvent(joypad.PadSelect, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.gb.JoyPad.KeyEvent(joypad.PadStart, true)
	}

	// release
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		g.gb.JoyPad.KeyEvent(joypad.PadLeft, false)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		g.gb.JoyPad.KeyEvent(joypad.PadRight, false)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		g.gb.JoyPad.KeyEvent(joypad.PadUp, false)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		g.gb.JoyPad.KeyEvent(joypad.PadDown, false)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyZ) {
		g.gb.JoyPad.KeyEvent(joypad.PadA, false)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyX) {
		g.gb.JoyPad.KeyEvent(joypad.PadB, false)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		g.gb.JoyPad.KeyEvent(joypad.PadSelect, false)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
		g.gb.JoyPad.KeyEvent(joypad.PadStart, false)
	}
}
