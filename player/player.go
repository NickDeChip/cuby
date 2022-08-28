package player

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Cube  rl.Rectangle
	Speed float32
}

func (p *Player) Move(score int32, dt float32) {
	movement := rl.Vector2{
		X: 0,
		Y: 0,
	}

	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		movement.X = 1
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		movement.X = -1
	}
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		movement.Y = -1
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		movement.Y = 1
	}
	if movement.X != 0 || movement.Y != 0 {
		movement = rl.Vector2Normalize(movement)
		p.Cube.X += movement.X * p.Speed * (float32(score) / 2) * dt
		p.Cube.Y += movement.Y * p.Speed * (float32(score) / 2) * dt
	}

}
