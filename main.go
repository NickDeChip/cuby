package main

import (
	"fmt"
	"github.com/NickDeChip/Cuby/player"
	"github.com/gen2brain/raylib-go/raylib"
)

type state struct {
	player     player.Player
	fruit      rl.Rectangle
	score      int32
	mute       bool
	difficulty string
	enemySpeed float32
	isGameOver bool
	enemies    []rl.Rectangle
	eating     rl.Sound
}

func main() {
	//screen/sound
	rl.InitWindow(500, 500, "Cuby - V1.4.2")
	rl.InitAudioDevice()

	//Fpscap
	screenFps := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
	rl.SetTargetFPS(screenFps)

	//Variables

	state := &state{
		player: player.Player{
			Cube:  playerCube(),
			Speed: 0,
		},
		fruit:      setUpFruit(),
		score:      1,
		mute:       true,
		difficulty: "",
		enemySpeed: 0,
		isGameOver: false,
		enemies:    make([]rl.Rectangle, 0),
		eating:     rl.LoadSound("resources/eating.mp3"),
	}

	state.enemies = append(state.enemies, setUpEnemy())

	for !rl.WindowShouldClose() {

		update(state)
		//Drawing
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if state.difficulty != "" {
			rl.DrawFPS(10, 478)
			rl.DrawRectangle(int32(state.player.Cube.X), int32(state.player.Cube.Y), 45, 45, rl.Red)
			rl.DrawRectangle(int32(state.fruit.X), int32(state.fruit.Y), 30, 30, rl.Green)
			for _, enemy := range state.enemies {
				rl.DrawRectangle(int32(enemy.X), int32(enemy.Y), int32(enemy.Width), int32(enemy.Height), rl.Magenta)
			}
			rl.DrawText(fmt.Sprintf("Score: %d", state.score), 10, 10, 32, rl.Black)
			rl.DrawText(state.difficulty, 360, 10, 32, rl.Blue)
			if state.mute {
				rl.DrawText("Muted", 360, 35, 28, rl.Black)
			}
		}
		if state.difficulty == "" {
			rl.DrawText("Pick Your difficulty:", 75, 100, 32, rl.Magenta)
			rl.DrawText("EASY = 1", 175, 135, 32, rl.Lime)
			rl.DrawText("MEDIUM = 2", 175, 170, 32, rl.Gold)
			rl.DrawText("HARD = 3", 175, 205, 32, rl.Orange)
			rl.DrawText("HELL = 4", 175, 240, 32, rl.Red)
		}

		if state.isGameOver {
			rl.DrawText("Press R to restart!", 10, 50, 32, rl.Black)
		}
		rl.EndDrawing()
	}

	//unloading
	rl.UnloadSound(state.eating)

	rl.CloseAudioDevice()
	rl.CloseWindow()
}

// Functions
func setUpEnemy() rl.Rectangle {
	size := float32(rl.GetRandomValue(30, 45))
	return rl.Rectangle{
		X:      float32(rl.GetRandomValue(30, 470)),
		Y:      float32(-10),
		Height: size,
		Width:  size,
	}
}

func setUpFruit() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(rl.GetRandomValue(30, 470)),
		Y:      float32(rl.GetRandomValue(30, 470)),
		Width:  30,
		Height: 30,
	}
}

func playerCube() rl.Rectangle {
	return rl.Rectangle{
		X:      250 - 45,
		Y:      250 - 45,
		Width:  45,
		Height: 45,
	}
}

func clampRecToScreen(rec *rl.Rectangle) {
	if rec.X >= 500-45 {
		rec.X = 500 - 45
	}
	if rec.X <= 0 {
		rec.X = 0
	}
	if rec.Y >= 500-45 {
		rec.Y = 500 - 45
	}
	if rec.Y <= 0 {
		rec.Y = 0
	}

}

func restart(state *state) {
	state.player.Cube.X = 250 - 45
	state.player.Cube.Y = 250 - 45
	state.isGameOver = false
	state.fruit = setUpFruit()
	state.score = 1
	state.enemies = make([]rl.Rectangle, 0)
	state.enemies = append(state.enemies, setUpEnemy())
}

func update(state *state) {
	if rl.IsKeyPressed(rl.KeyR) {
		restart(state)
	}
	if state.isGameOver {
		return
	}
	dt := rl.GetFrameTime()
	state.player.Move(state.score, dt)

	for idx := range state.enemies {
		state.enemies[idx].Y += state.enemySpeed * (float32(state.score) / 2) * dt
	}

	//difficultys
	if rl.IsKeyPressed(rl.KeyOne) {
		state.player.Speed = 60
		state.enemySpeed = 30
		state.difficulty = "EASY"
		restart(state)
	}

	if rl.IsKeyPressed(rl.KeyTwo) {
		state.player.Speed = 70
		state.enemySpeed = 50
		state.difficulty = "MEDIUM"
		restart(state)
	}

	if rl.IsKeyPressed(rl.KeyThree) {
		state.player.Speed = 120
		state.enemySpeed = 100
		state.difficulty = "HARD"
		restart(state)
	}
	if rl.IsKeyPressed(rl.KeyFour) {
		state.player.Speed = 100
		state.enemySpeed = 80
		state.difficulty = "HELL"
		restart(state)
	}

	//Collsions
	for idx, enemy := range state.enemies {
		if enemy.Y >= 500 {
			state.enemies[idx] = setUpEnemy()
		}
		if rl.CheckCollisionRecs(state.player.Cube, enemy) {
			state.isGameOver = true
			break
		}
	}

	if rl.CheckCollisionRecs(state.player.Cube, state.fruit) {
		state.fruit = setUpFruit()
		state.score += 1
		if state.score%5 == 0 && state.difficulty == "HELL" {
			state.enemies = append(state.enemies, setUpEnemy())
		}
		if !state.mute {
			rl.PlaySound(state.eating)
		}
	}

	clampRecToScreen(&state.player.Cube)

	//Muting
	if rl.IsKeyPressed(rl.KeyM) {
		state.mute = !state.mute
	}
}
