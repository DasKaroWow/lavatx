package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v3"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("failed to create screen: %v", err)
	}
	err = screen.Init()
	if err != nil {
		log.Fatalf("failed to init screen: %v", err)
	}
	defer screen.Fini()

	ticker := time.NewTicker(time.Second / 15)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			screen.EventQ() <- tcell.NewEventInterrupt(nil)
		}
	}()

	balls := spawnBalls(10, screen)
	for {
		ev := <-screen.EventQ()

		switch ev := ev.(type) {
		case *tcell.EventInterrupt:
			width, height := screen.Size()
			screen.Clear()
			for index := range balls {
				moveBall(&balls[index], width, height)
			}
			drawField(balls, screen)

			screen.Show()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Str() == "q" {
				return
			}
		}

	}
}
