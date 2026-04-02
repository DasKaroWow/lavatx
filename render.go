package main

import (
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

const DefaultGlyph rune = '█'

func fieldStrengthAt(x, y int, balls []Ball) float64 {
	var strength float64

	for _, ball := range balls {
		dx := float64(x - ball.X)
		dy := float64(y - ball.Y)

		dist2 := dx*dx + dy*dy + 1
		r := float64(ball.Power)

		strength += (r * r) / dist2
	}

	return strength
}

func cellFromStrength(strength float64) (rune, tcell.Style) {
	switch {
	case strength < 0.18:
		return ' ', tcell.StyleDefault.Background(color.Black)

	case strength < 0.35:
		return '░', tcell.StyleDefault.Foreground(color.Maroon)

	case strength < 0.65:
		return '▒', tcell.StyleDefault.Foreground(color.Red)

	case strength < 1.00:
		return '▓', tcell.StyleDefault.Foreground(color.Orange)

	case strength < 1.45:
		return '█', tcell.StyleDefault.Foreground(color.Yellow)

	default:
		return '█', tcell.StyleDefault.Foreground(color.White)
	}
}

func drawField(balls []Ball, screen tcell.Screen) {
	width, height := screen.Size()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			strength := fieldStrengthAt(x, y, balls)
			glyph, style := cellFromStrength(strength)
			screen.SetContent(x, y, glyph, nil, style)
		}
	}
}
