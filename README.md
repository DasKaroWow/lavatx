# lavatx

A small terminal animation project written in Go.

This project renders a lava lamp / metaballs-like effect directly in the terminal. Several moving points influence a scalar field, and each terminal cell is drawn based on the field strength at that position. Depending on the value, the program chooses a character and a color, which creates a fluid animated look.

## About

This is a small practice project I made while learning Go.

The goal was not to build a perfect or highly optimized program, but to make something visual, simple, and fun while practicing real Go code instead of another basic toy example.

It was mainly built to practice:

- structuring a small Go project
- terminal rendering
- frame-based updates
- simple simulation logic
- separating code by responsibility

So this is basically a random first project, but a useful one for learning.

## Inspiration

This project is based on the idea from AngelJumbo/lavat:

https://github.com/AngelJumbo/lavat

So this is essentially a Go reimplementation of that concept, made with the goal of being OS-independent and easy to run as a normal terminal application.

## Project Structure

- ball.go — defines the Ball structure
- sim.go — handles spawning and movement
- render.go — calculates field strength and draws the frame
- main.go — initializes the screen and runs the main loop

## Tech

- Go
- tcell

## Notes

This is a learning project, so the implementation is intentionally simple. The focus here was more on practice, experimentation, and getting comfortable with Go than on making a perfect clone or a fully polished terminal renderer.

## Run

```bash
git clone https://github.com/DasKaroWow/lavatx.git
cd lavatx
go run .
```
