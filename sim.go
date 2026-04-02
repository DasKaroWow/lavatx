package main

import (
	"math/rand/v2"

	"github.com/gdamore/tcell/v3"
)

func spawnBalls(amount int, screen tcell.Screen) []Ball {
	balls := make([]Ball, amount)
	width, height := screen.Size()
	for index := range balls {
		balls[index] = Ball{
			X:     rand.IntN(width),
			Y:     rand.IntN(height),
			VX:    rand.IntN(2)*2 - 1,
			VY:    rand.IntN(2)*2 - 1,
			Power: rand.IntN(4) + 4,
		}
	}
	return balls
}

func moveBall(ball *Ball, width int, height int) {
	nextX := ball.X + ball.VX
	nextY := ball.Y + ball.VY

	if nextX < 0 || nextX >= width {
		ball.VX = -ball.VX
		nextX = ball.X + ball.VX
	}

	if nextY < 0 || nextY >= height {
		ball.VY = -ball.VY
		nextY = ball.Y + ball.VY
	}

	ball.X = nextX
	ball.Y = nextY
}
